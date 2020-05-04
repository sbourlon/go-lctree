// Package lctree provides primitives to serialize and deserialize leetcode trees
package lctree

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// TreeNode is a node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// WalkFn is the signature of the function called on each node
// visited by the function WalkDepthFirst
type WalkFn func(n *TreeNode) error

// WalkDepthFirst traverses the tree depth-first.
// fn is called on each visited node
func (t *TreeNode) WalkDepthFirst(fn WalkFn) error {
	stack := make([]*TreeNode, 0)
	stack = append(stack, t)

	for len(stack) > 0 {
		last := len(stack) - 1
		node := stack[last]
		stack = stack[:last]
		if err := fn(node); err != nil {
			return err
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	return nil
}

// WalkBFFn is the signature of the function called on each node
// by the function WalkBreadthFirst
type WalkBFFn func(n *TreeNode, depth int) error

type walkBFTreeNode struct {
	Node  *TreeNode
	Depth int
}

// WalkBreadthFirst traverses the tree breadth-first.
// fn is called on each visited node
func (t *TreeNode) WalkBreadthFirst(fn WalkBFFn) error {
	queue := make([]walkBFTreeNode, 0)
	queue = append(queue, walkBFTreeNode{t, 0})

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if err := fn(node.Node, node.Depth); err != nil {
			return err
		}
		if node.Node != nil {
			queue = append(queue, walkBFTreeNode{node.Node.Left, node.Depth + 1})
			queue = append(queue, walkBFTreeNode{node.Node.Right, node.Depth + 1})
		}
		// if node.Node.Left != nil {
		// 	queue = append(queue, walkBFTreeNode{node.Node.Left, node.Depth + 1})
		// }
		// if node.Node.Right != nil {
		// 	queue = append(queue, walkBFTreeNode{node.Node.Right, node.Depth + 1})
		// }
	}
	return nil
}

// DOT converts a TreeNode tree into Graphviz DOT
// See: http://www.graphviz.org/documentation/
func (t *TreeNode) DOT() string {
	tree := strings.Builder{}
	edges := strings.Builder{}

	tree.WriteString("digraph {")

	edgeStyleInvisible := "[style=invis]"
	nodeStyleInvisible := "[label=\"\", width=.1, style=invis]"

	fn := func(n *TreeNode, depth int) error {
		if depth == 0 {
			if n != nil {
				tree.WriteString("\ngraph [ordering=\"out\"];")
			} else {
				return nil
			}
		}
		if n != nil {
			v := strconv.Itoa(n.Val)
			tree.WriteString("\n")
			tree.WriteString(v)
			tree.WriteString(";")

			if n.Left != nil || n.Right != nil {
				for i, leaf := range []*TreeNode{n.Left, n.Right} {
					edges.WriteString("\n")
					edges.WriteString(v)
					edges.WriteString(" -> ")
					if leaf != nil {
						edges.WriteString(strconv.Itoa(leaf.Val))
					} else {
						// Add invisible node to keep the tree balanced
						// See https://graphviz.gitlab.io/faq/#FaqBalanceTree
						node := fmt.Sprintf("%s.%d", v, i)
						tree.WriteString("\n")
						tree.WriteString(fmt.Sprintf("%s %s;", node, nodeStyleInvisible))
						edges.WriteString(fmt.Sprintf("%s %s", node, edgeStyleInvisible))
					}
					edges.WriteString(";")
				}
			}
		}
		return nil
	}

	t.WalkBreadthFirst(fn)

	if edges.Len() > 0 {
		tree.WriteString(edges.String())
		tree.WriteString("\n")
	}
	tree.WriteString("}\n")

	return tree.String()
}

// Deserialize converts a leetcode serialized tree into a TreeNode tree
func Deserialize(tree string) *TreeNode {
	var root, parent *TreeNode
	var right bool

	if tree == "[]" {
		return nil
	}

	buf := bytes.Buffer{}
	queue := make([]*TreeNode, 0)

	for _, r := range tree {
		switch r {
		case ',', ']':
			var node *TreeNode

			// Get node
			s := buf.String()
			buf.Reset()
			if s != "null" {
				v, _ := strconv.Atoi(s)
				node = &TreeNode{Val: v}
			}

			// Set root
			if root == nil {
				root, parent = node, node
				continue
			}

			// Append none null nodes to the queue
			if node != nil {
				queue = append(queue, node)
			}

			// Set Right or left node
			if right {
				parent.Right = node
				parent = queue[0]
				queue = queue[1:]
			} else {
				parent.Left = node
			}

			// Switch on and off right
			right = !right
		case '[':
			continue
		default:
			buf.WriteRune(r)
		}
	}

	return root
}

// Serialize converts a TreeNode tree into a leetcode serialized tree
func Serialize(t *TreeNode) string {
	if t == nil {
		return "[]"
	}

	tree := make([]string, 0)

	fn := func(n *TreeNode, depth int) error {
		var v string
		if n != nil {
			v = strconv.Itoa(n.Val)
		} else {
			v = "null"
		}
		tree = append(tree, v)
		return nil
	}

	t.WalkBreadthFirst(fn)

	// Remove all trailing null
	cut := len(tree)
	for i := len(tree) - 1; i > 0; i-- {
		if tree[i] == "null" {
			cut--
		} else {
			break
		}
	}

	return fmt.Sprintf("[%s]", strings.Join(tree[:cut], ","))
}
