package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func reverseLeftWords(s string, n int) string {
	return s[n:] + s[:n]
}
