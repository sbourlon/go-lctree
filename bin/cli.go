package main

import (
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/sbourlon/go-lctree"
)

func main() {
	if len(os.Args) != 2 {
		usage(os.Args)
		os.Exit(1)
	}
	encoded := os.Args[1]

	tree := lctree.Deserialize(encoded)
	dot := tree.DOT()
	fmt.Println(dot)
}

func usage(args []string) {
	bin := args[0]
	usage := `Usage: {{.}} <leetcode tree>
Convert a leetcode tree into DOT

Example:
- Print the tree in DOT
{{.}} "[1,null,2,3]"

- Open the tree directly with an image viewer (e.g feh)
{{.}} "[1,null,2,3]" | dot -Tpng | feh -
`
	t := template.Must(template.New("usage").Parse(usage))
	err := t.Execute(os.Stdout, bin)
	if err != nil {
		log.Println("executing template:", err)
	}
}
