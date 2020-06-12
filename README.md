# AVL Tree

I am using this project as an opportunity to learn Golang better, as well as a way to better understand AVL trees before moving on to the WAVL implementation.

* All basic BST functionality is there and working. 
* All AVL functionality is there and working. Though all functions are in their bottom-up recursive versions, instead of the more optimal top-down iterative ones.
* Height/Balance property optimizations have been added such that node height & balance are updated in constant time during Insert and Remove backtracking.
* The data structure allows for generic data types with the condition that the data type implements the **Less** function. This condition is what enables the tree to compare and sort without knowing the underlying type of the data it stores.

Next up:
1. Add as many types as reasonably possible to the list of types which implement Less. Should make the tree more flexible.
2. Fix the test code. While it gets the job done, it's the first test code I'v ever written so it's pretty bad.
