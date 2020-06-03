## Build:
```
go build .
```

## Usage:
```
Usage: ./lctree <leetcode tree>
Convert a leetcode tree into DOT

Example:
- Print the tree in DOT
./lctree "[1,null,2,3]"

- Open the tree directly with an image viewer (e.g feh)
./lctree "[1,null,2,3]" | dot -Tpng | feh -
```
