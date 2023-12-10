package main

import (
	"container/heap"
	"fmt"
)

type HuffmanNode struct {
	Value     byte
	Frequency int
	Left      *HuffmanNode
	Right     *HuffmanNode
}

type HuffmanPriorityQueue []*HuffmanNode

func (pq HuffmanPriorityQueue) Len() int {
	return len(pq)
}
func (pq HuffmanPriorityQueue) Less(i, j int) bool {
	return pq[i].Frequency < pq[j].Frequency
}
func (pq HuffmanPriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *HuffmanPriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*HuffmanNode))
}

func (pq *HuffmanPriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	*pq = old[0 : n-1]
	return node
}

func BuildHuffmanTree(frequencies map[byte]int) *HuffmanNode {
	priorityQueue := make(HuffmanPriorityQueue, len(frequencies))
	i := 0
	for char, frequency := range frequencies {
		priorityQueue[i] = &HuffmanNode{Value: char, Frequency: frequency}
		i++
	}

	heap.Init(&priorityQueue)

	for len(priorityQueue) > 1 {
		left := heap.Pop(&priorityQueue).(*HuffmanNode)
		right := heap.Pop(&priorityQueue).(*HuffmanNode)

		internalNode := &HuffmanNode{
			Frequency: left.Frequency + right.Frequency,
			Left:      left,
			Right:     right,
		}

		heap.Push(&priorityQueue, internalNode)
	}

	return priorityQueue[0]
}

func GenerateHuffmanCodes(root *HuffmanNode, code string, codes map[byte]string) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		codes[root.Value] = code
		return
	}

	GenerateHuffmanCodes(root.Left, code+"0", codes)
	GenerateHuffmanCodes(root.Right, code+"1", codes)
}

func Compress(input string) (string, map[byte]string) {
	frequencies := make(map[byte]int)
	for _, char := range input {
		frequencies[byte(char)]++
	}

	root := BuildHuffmanTree(frequencies)

	codes := make(map[byte]string)
	GenerateHuffmanCodes(root, "", codes)

	compressed := ""
	for _, char := range input {
		compressed += codes[byte(char)]
	}

	return compressed, codes
}

func main() {
	input := "BCCABBDDAECCBBAEDDCC"
	compressed, codes := Compress(input)

	fmt.Println("Original:", input)
	fmt.Println("Compressed:", compressed)
	fmt.Println("Huffman Codes:")
	for char, code := range codes {
		fmt.Printf("%c: %s\n", char, code)
	}

	/*
		THE OUTPUT
			Original: BCCABBDDAECCBBAEDDCC
			Compressed: 101111011101000000110101111101001101000001111
			Huffman Codes:
			C: 11
			D: 00
			E: 010
			A: 011
			B: 10
	*/
}
