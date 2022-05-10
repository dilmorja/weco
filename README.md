# Weco

Weco is an acronym for "Web Core". It provides a collection of APIs to easily interact with the browser natively, as well as tools that help with browser-side application development.

## Fast examples

Here are some commonly used examples. For more, you can review the [tests](./test/).

### Console.Log ([window](./test/window/))

**In JavaScript**:
```js
console.log("This is a example")
```

**In Go (with `weco/window`)**:
```go
package main

import(
  . "github.com/dilmorja/weco/window"
)

func main() {
  Console.Log("This is a example")
}
```

**In Go (with `syscall/js`)**:

```go
package main

import "syscall/js"

func main() {
  js.Global().Get("window").Get("console").Call("log", "This is a example")
}
```

## License

This project and its submodules are licensed under the [BSD 3-Clause](LICENSE) license.