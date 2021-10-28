#   Binary Search Tree
This is a project that deals with implementations of a binary tree and the following functions.

### ```Print```
Prints the entire tree. <br>
Arguments: <br>
&emsp; 1. w -> any writer of type io.Writer <br>
&emsp; 2. space -> For pretty printing the tree. <br> 
&emsp; 3. node -> Root node of the tree <br>
&emsp; 4. ch -> a character for to represent the side of the node(s) being printed. <br>

### ```Tree.Insert```
Inserts a node into the tree
Arguments: <br>
&emsp; 1. data -> an int64 <br>

### ```Node.New```
Creates a new node and inserts the node left or right depending on the size of the data relative to the root node.<br>
Arguments: <br>
&emsp; 1. data -> an int64 <br>

### ```IsSameTree```
Asserts that the provided trees are identical trees are identical<br>
Arguments: <br>
&emsp; 1. root1 -> root node for the first tree <br>
&emsp; 1. root2 -> root node for the second tree <br>
