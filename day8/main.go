package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strings"
)

type Point struct {
	x, y, z int
}

type Connection struct {
	a, b int
	dist float64
}

type UnionFind struct {
	parent []int
	size []int
}

func newUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	size := make([]int, n)

	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}

	return &UnionFind{parent, size}
}

func (uf *UnionFind) Find(i int) int {
	if uf.parent[i] != i {
		uf.parent[i] = uf.Find(uf.parent[i])
	}
	return uf.parent[i]
}

func (uf *UnionFind) Union(i int, j int) {
	rootI := uf.Find(i)
	rootJ := uf.Find(j)

	if rootI != rootJ {
		uf.parent[rootJ] = rootI
		uf.size[rootI] += uf.size[rootJ]
	}
}

func dist(p1, p2 Point) float64 {
	dx := float64(p1.x - p2.x)
	dy := float64(p1.y - p2.y)
	dz := float64(p1.z - p2.z)
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func part1(file *os.File) int {
	return 1
}

func part2(file *os.File) int {
	return 1
}

func main() {
	file, err := os.Open("testCase.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	answer1 := part1(file)

	_, err = file.Seek(0, 0)
	if err != nil {
		log.Fatal(err)
	}
	answer2 := part2(file)

	fmt.Println("Answer to part 1:", answer1)
	fmt.Println("Answer to part 2:", answer2)
}