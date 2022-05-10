package main

import(
	"flag"
	"runtime"
	"log"
	"path/filepath"
	"html/template"
	"net/http"
)

var wasmPath string

func init() {
	flag.StringVar(&wasmPath, "i", "", "Path to WASM")
	flag.Parse()
}

const testHTML string = `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <script src="/wasm_exec.js"></script>
  <script>
    if (!WebAssembly.instantiateStreaming) { // polyfill
			WebAssembly.instantiateStreaming = async (resp, importObject) => {
				const source = await (await resp).arrayBuffer();
				return await WebAssembly.instantiate(source, importObject);
			};
		}

		const go = new Go();

		WebAssembly.instantiateStreaming(fetch("/test.wasm"), go.importObject).then((result) => {
			run(result.instance, result.module)
		}).catch((err) => {
			console.error(err);
		});

    async function run(instance, module) {
      await go.run(instance);
			await WebAssembly.instantiate(module, go.importObject);
    }
  </script>
  <title>WebAssembly Test</title>
</head>
<body>
</body>
</html>`

func testServer(wasmexec string, gowasm string) {

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		t, err := template.New("home").Parse(testHTML)

		if err != nil {
			log.Fatal(err)
		}

		t.Execute(res, nil)
	})

	http.HandleFunc("/wasm_exec.js", func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, wasmexec)
	})

	http.HandleFunc("/test.wasm", func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, gowasm)
	})

	log.Println("[SERVER] Started on http://localhost:9090")
	log.Fatal(http.ListenAndServe(":9090", nil))
}

func main() {

	if wasmPath == "" {
		log.Fatal("Invalid path to WASM")
	}

	testServer(filepath.Join(runtime.GOROOT(),
		"misc", "wasm", "wasm_exec.js"), wasmPath)
}