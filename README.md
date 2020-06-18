# AVL Tree

I am using this project as an opportunity to learn Golang better, as well as a way to better understand AVL trees and WebAssembly.

* The library can be imported as a WASM module by Javascript programs.
* All basic BST functionality is there and working. 
* All AVL functionality is there and working. Though all functions are in their bottom-up recursive versions, instead of the more optimal top-down iterative ones.
* Height/Balance property optimizations have been added such that node height & balance are updated in constant time during Insert and Remove backtracking.
* The data structure allows for generic datatypes with the condition that the datatype implements the **Less** function. This condition is what enables the tree to compare and sort without knowing the underlying type of data it stores.
* Generic datatypes includes all Javascript types too. The library can handle Javascript built-in types, as well as user-defined types assuming they implement the **Less** function.

Next up:
1. Fix the test code. While it gets the job done, it's the first test code I've ever written so it's pretty bad.
2. Example directories for how to use this library in a Javascript program.

# Instructions
Grab the WASM module and wasm_exec.js files under the Releases tab, then have a look at steps 6 & 7.
Otherwise, if you'd like to generate the WASM module manually, then the following steps will help.

To import this library as a WASM module in a Javascript program:
1. Ensure you have the most recent version of Go installed.
2. Clone or download this repo locally.
3. Create a main.go file containing the following:
```
package main
import "fmt"
func main() {
  persist := make(chan struct{}, 0)
  fmt.Println("Initializing WASM module...")
  tree := NewAVL()
  fmt.Sprintf("%v", tree)
  <-persist
}
```
4. Place the wasm_exec.js file that comes with the Go installation inside the directory
5. Enter the following command within the same directory to generate the WASM module:
```
GOOS=js GOARCH=wasm go build -o main.wasm
```
6. If you have an index.html file (or equivalent), ensure that you've include the following script tags:
```
<script src="wasm_exec.js"></script>
<script>
  const go = new Go()
  WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
    // Interact with WASM module here so that it may persist
  })
</script>
// The second script tag may be relegated to its own Javascript file, rather than being crammed into an html file.
```
7. Run your server. Your job is done!
