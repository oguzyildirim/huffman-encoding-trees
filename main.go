package main

import (
	"fmt"
)

func makeLeaf(symbol string, weight int) []interface{} {
	return []interface{}{"leaf", symbol, weight}
}

func isLeaf(object []interface{}) bool {
	return object[0] == "leaf"
}

func symbolLeaf(x []interface{}) string {
	return x[1].(string)
}

func weightLeaf(x []interface{}) int {
	return x[2].(int)
}

func makeCodeTree(left []interface{}, right []interface{}) []interface{} {
	return []interface{}{"code_tree", left, right, append(symbols(left), symbols(right)...), weight(left) + weight(right)}
}

func leftBranch(tree []interface{}) []interface{} {
	return tree[1].([]interface{})
}

func rightBranch(tree []interface{}) []interface{} {
	return tree[2].([]interface{})
}

func symbols(tree []interface{}) []interface{} {
	if isLeaf(tree) {
		return []interface{}{symbolLeaf(tree)}
	} else {
		return tree[3].([]interface{})
	}
}

func weight(tree []interface{}) int {
	if isLeaf(tree) {
		return weightLeaf(tree)
	} else {
		return tree[4].(int)
	}
}

func decode(bits []int, tree []interface{}) []interface{} {
	var decode1 func(bits []int, currentBranch []interface{}) []interface{}
	decode1 = func(bits []int, currentBranch []interface{}) []interface{} {
		if len(bits) == 0 {
			return nil
		} else {
			nextBranch := chooseBranch(bits[0], currentBranch)
			if isLeaf(nextBranch) {
				return append([]interface{}{symbolLeaf(nextBranch)}, decode1(bits[1:], tree)...)
			} else {
				return decode1(bits[1:], nextBranch)
			}
		}
	}
	return decode1(bits, tree)
}

func chooseBranch(bit int, branch []interface{}) []interface{} {
	if bit == 0 {
		return leftBranch(branch)
	} else if bit == 1 {
		return rightBranch(branch)
	} else {
		panic(bit)
	}
}

func main() {
	// Test Huffman coding
	leafA := makeLeaf("A", 5)
	leafB := makeLeaf("B", 2)
	leafC := makeLeaf("C", 1)
	tree := makeCodeTree(makeCodeTree(leafA, leafB), leafC)

	bits := []int{0, 1, 1, 0, 1, 0, 0, 1, 0}
	decoded := decode(bits, tree)

	fmt.Println("Decoded symbols:", decoded)
}
