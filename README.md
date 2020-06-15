# AVL Tree

I am using this project as an opportunity to learn Golang better, as well as a way to better understand AVL trees and WebAssembly.

* The library can be imported as a WASM module by Javascript programs.
* All basic BST functionality is there and working. 
* All AVL functionality is there and working. Though all functions are in their bottom-up recursive versions, instead of the more optimal top-down iterative ones.
* Height/Balance property optimizations have been added such that node height & balance are updated in constant time during Insert and Remove backtracking.
* The data structure allows for generic data types with the condition that the data type implements the **Less** function. This condition is what enables the tree to compare and sort without knowing the underlying type of the data it stores.

Next up:
1. Fix the test code. While it gets the job done, it's the first test code I've ever written so it's pretty bad.
2. Make syscall/js code conditionally compiled. Some people may want to use it, others might not since the package is experimental or something. Would be better if the sections of this library utilizing the syscall/js code were optional.
3. Add a Postorder traversal method.
