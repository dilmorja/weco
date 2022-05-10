package main

import(
	"flag"
	"os"
	"os/exec"
	"runtime"
	"log"
	"path/filepath"
	"html/template"
	"net/http"
)

// tool subcommands
var(
	testCmd = flag.NewFlagSet("test", flag.ExitOnError)
)

// testCmd subcommand flags
var(
	testInput = testCmd.String("i", "", "input file path")
	testOutput = testCmd.String("o", "", "output file path")
)

func init() {

	switch os.Args[1] {
	case "test":
		testCmd.Parse(os.Args[2:])
	default:
		flag.Usage()
	}

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

func test() {
	if *testInput != "" && *testOutput != "" {
		// save the GOOS and the GOARCH
		actualOs, actualArch := runtime.GOOS, runtime.GOARCH

		if actualOs == "js" && actualArch == "wasm" {
			println("Your OS (js) and architecture (wasm) do not match the specifications of your computer.")
			os.Exit(2)
		}

		log.Println("[VERIFYING] environment")
		if actualOs != "js" && actualArch != "wasm" {
			log.Println("[SETTING] environment")
			envCMD := exec.Command("go", "env", "-w", "GOOS=js", "GOARCH=wasm")
			envCMD.Stdout = os.Stdout
			err := envCMD.Run()
			if err != nil {
				log.Fatal("[FAILED] setting environment:\n\t", err)
			}
		}

		log.Println("[BUILDING] wasm")
		buildCMD := exec.Command("go", "build", "-o", filepath.Join(*testOutput), filepath.Join(*testInput))
		buildCMD.Stdout = os.Stdout
		err := buildCMD.Run()
		if err != nil {
			log.Fatal("[FAILED] building wasm:\n\t", err)
		}

		if actualOs != "js" && actualArch != "wasm" {
			log.Println("[RESTORING] environment")
			restoreEnvCMD := exec.Command("go", "env", "-w", "GOOS="+actualOs, "GOARCH="+actualArch)
			restoreEnvCMD.Stdout = os.Stdout
			err = restoreEnvCMD.Run()
			if err != nil {
				log.Fatal("[FAILED] restoring environment:\n\t", err)
			}
		}

		log.Println("[SERVER] Starting...")
		testServer(filepath.Join(runtime.GOROOT(), "misc", "wasm", "wasm_exec.js"), *testOutput)

	} else {
		testCmd.Usage()
	}
}

func testServer(wasmexec string, gowasm string) {

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		t, err := template.New("home").Parse(testHTML)

		if err != nil {
			panic(err)
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
	http.ListenAndServe(":9090", nil)
}

func main() {}