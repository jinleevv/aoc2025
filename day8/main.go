package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
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
	var points []Point
	var edges []Connection
	var sizes []int

	scanner := bufio.NewScanner(file)
	// create an array with its position
	for scanner.Scan() {
		line := scanner.Text()
		position := strings.Split(line, ",")

		x, _ := strconv.Atoi(strings.TrimSpace(position[0]))
		y, _ := strconv.Atoi(strings.TrimSpace(position[1]))
		z, _ := strconv.Atoi(strings.TrimSpace(position[2]))

		points = append(points, Point{x, y, z})
	}

	// create an array with its distance calculated
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			dist := dist(points[i], points[j])
			edges = append(edges, Connection{a: i, b: j, dist: dist})
		}
	}

	// sort the edges
	sort.Slice(edges, func(i int, j int) bool {
		return edges[i].dist < edges[j].dist
	})

	// start Union Find logic
	uf := newUnionFind(len(points))

	limit := 1000

	// 1000 closest points are merged
	for i := 0; i < limit; i++ {
		edge := edges[i]
		uf.Union(edge.a, edge.b)
	}

	// update the size
	for i := 0; i < len(points); i++ {
		if uf.parent[i] == i {
			sizes = append(sizes, uf.size[i])
		}
	}

	// sort the sizes
	sort.Slice(sizes, func(i int, j int) bool {
		return sizes[i] > sizes[j]
	})

	res := 1
	count := 0
	for _, s := range sizes {
		res *= s
		count++
		// count largest 3
		if count == 3 {
			break
		}
	}


	return res
}

func part2(file *os.File) int {
	var points []Point
	var edges []Connection

	scanner := bufio.NewScanner(file)
	// create an array with its position
	for scanner.Scan() {
		line := scanner.Text()
		position := strings.Split(line, ",")

		x, _ := strconv.Atoi(strings.TrimSpace(position[0]))
		y, _ := strconv.Atoi(strings.TrimSpace(position[1]))
		z, _ := strconv.Atoi(strings.TrimSpace(position[2]))

		points = append(points, Point{x, y, z})
	}

	// create an array with its distance calculated
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			dist := dist(points[i], points[j])
			edges = append(edges, Connection{a: i, b: j, dist: dist})
		}
	}

	// sort the edges
	sort.Slice(edges, func(i int, j int) bool {
		return edges[i].dist < edges[j].dist
	})

	// start Union Find logic
	uf := newUnionFind(len(points))

	limit := len(points)

	for _, edge := range edges {
		rootA := uf.Find(edge.a)
		rootB := uf.Find(edge.b)

		if rootA != rootB {
			uf.Union(edge.a, edge.b)
			limit--
			if limit == 1 {
				x1 := points[edge.a].x
				x2 := points[edge.b].x

				return x1 * x2
			}
		}
	}


	return 0
}

func main() {
	file, err := os.Open("data.txt")
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