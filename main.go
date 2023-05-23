// Package main demonstrates Huffman coding algorithm.
package main

import (
	"fmt"
)

// Node represents a generic node in the code tree.
type Node interface{}

// Leaf represents a leaf node in the code tree.
type Leaf struct {
	Symbol string // Symbol associated with the leaf node.
	Weight int    // Weight (frequency) of the symbol.
}

// CodeTree represents the Huffman code tree.
type CodeTree struct {
	Left    Node     // Left branch of the tree.
	Right   Node     // Right branch of the tree.
	Symbols []string // Symbols associated with the tree.
	Weight  int      // Weight (sum of weights of symbols) of the tree.
}

// makeLeaf creates a new leaf node with the given symbol and weight.
func makeLeaf(symbol string, weight int) Node {
	return &Leaf{Symbol: symbol, Weight: weight}
}

// isLeaf checks if the given node is a leaf node.
func isLeaf(object Node) bool {
	_, ok := object.(*Leaf)
	return ok
}

// symbolLeaf returns the symbol associated with the leaf node.
func symbolLeaf(x Node) string {
	leaf := x.(*Leaf)
	return leaf.Symbol
}

// weightLeaf returns the weight associated with the leaf node.
func weightLeaf(x Node) int {
	leaf := x.(*Leaf)
	return leaf.Weight
}

// makeCodeTree creates a new code tree by combining two nodes (trees).
func makeCodeTree(left, right Node) Node {
	leftSymbols := symbols(left)
	rightSymbols := symbols(right)
	combinedSymbols := append(leftSymbols, rightSymbols...)

	return &CodeTree{
		Left:    left,
		Right:   right,
		Symbols: combinedSymbols,
		Weight:  weight(left) + weight(right),
	}
}

// leftBranch returns the left branch of the code tree.
func leftBranch(tree Node) Node {
	codeTree := tree.(*CodeTree)
	return codeTree.Left
}

// rightBranch returns the right branch of the code tree.
func rightBranch(tree Node) Node {
	codeTree := tree.(*CodeTree)
	return codeTree.Right
}

// symbols returns the symbols associated with the node/tree.
func symbols(tree Node) []string {
	if isLeaf(tree) {
		return []string{symbolLeaf(tree)}
	} else {
		codeTree := tree.(*CodeTree)
		return codeTree.Symbols
	}
}

// weight returns the weight associated with the node/tree.
func weight(tree Node) int {
	if isLeaf(tree) {
		return weightLeaf(tree)
	} else {
		codeTree := tree.(*CodeTree)
		return codeTree.Weight
	}
}

// decodeBits decodes the given bits using the provided code tree.
func decodeBits(bits []int, tree Node) []string {
	var decode func(bits []int, currentBranch Node) []string
	decode = func(bits []int, currentBranch Node) []string {
		if len(bits) == 0 {
			return nil
		} else {
			nextBranch := chooseBranch(bits[0], currentBranch)
			if isLeaf(nextBranch) {
				return append([]string{symbolLeaf(nextBranch)}, decode(bits[1:], tree)...)
			} else {
				return decode(bits[1:], nextBranch)
			}
		}
	}
	return decode(bits, tree)
}

// chooseBranch selects the appropriate branch based on the given bit value.
func chooseBranch(bit int, branch Node) Node {
	if bit == 0 {
		return leftBranch(branch)
	} else if bit == 1 {
		return rightBranch(branch)
	} else {
		panic(bit)
	}
}

func main() {
	// Example 1: Huffman coding
	leafA := makeLeaf("A", 5)
	leafB := makeLeaf("B", 2)
	leafC := makeLeaf("C", 1)
	tree := makeCodeTree(makeCodeTree(leafA, leafB), leafC)

	bits := []int{0, 1, 1, 0, 1, 0, 0, 1, 0}
	decoded := decodeBits(bits, tree)

	fmt.Println("Example 1:")
	fmt.Println("Encoded bits:", bits)
	fmt.Println("Decoded symbols:", decoded)
	fmt.Println()

	// Example 2: Huffman coding
	leafX := makeLeaf("X", 3)
	leafY := makeLeaf("Y", 2)
	leafZ := makeLeaf("Z", 1)
	tree2 := makeCodeTree(makeCodeTree(leafX, leafY), leafZ)

	bits2 := []int{1, 0, 0, 1, 0, 1, 0}
	decoded2 := decodeBits(bits2, tree2)

	fmt.Println("Example 2:")
	fmt.Println("Encoded bits:", bits2)
	fmt.Println("Decoded symbols:", decoded2)
	fmt.Println()
}
