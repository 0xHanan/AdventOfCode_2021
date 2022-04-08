package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type VertexTable struct {
	v   [][]Vertex
	row int
	col int
}

type Vertex struct {
	weight   int
	distance int
	visited  bool
	prev     Index
}

type Index struct {
	x int
	y int
}

func (vt *VertexTable) CalculateDistance(from Index) {
	// left, right, down, up
	row := []int{0, 0, 1, -1}
	col := []int{-1, 1, 0, 0}

	// fmt.Println(from)
	// fmt.Println(vt.row, vt.col)
	for i := range row {
		if (from.x+row[i] <= vt.row) && (from.x+row[i] >= 0) {
			if (from.y+col[i] <= vt.col) && (from.y+col[i] >= 0) {
				if !vt.v[from.x+row[i]][from.y+col[i]].visited {
					distance := vt.v[from.x+row[i]][from.y+col[i]].weight + vt.v[from.x][from.y].distance
					cellDistance := vt.v[from.x+row[i]][from.y+col[i]].distance
					if distance < cellDistance || cellDistance == 0 {
						vt.v[from.x+row[i]][from.y+col[i]].distance = distance
						vt.v[from.x+row[i]][from.y+col[i]].prev = from
					}
				}

			}
		}
	}
	vt.v[from.x][from.y].visited = true
}

func (vt *VertexTable) AllVisited() bool {
	for l := range vt.v {
		for c := range vt.v[l] {
			if !vt.v[l][c].visited {
				return false
			}
		}
	}

	return true
}

func (vt *VertexTable) GetSmallIdx() Index {
	smallest := ^uint(0)
	smallestIdx := Index{0, 0}
	for l := range vt.v {
		for i, c := range vt.v[l] {
			if !c.visited && c.distance != 0 {
				if uint(c.distance) < smallest {
					smallest = uint(c.distance)
					smallestIdx = Index{l, i}
				}
			}
		}
	}
	return smallestIdx
}

func (vt *VertexTable) PrintTable() {
	for l := range vt.v {
		fmt.Println(vt.v[l])
	}
}

func part1(data *VertexTable) {

	// data.v[0][0].visited = true
	data.v[0][0].weight = 0
	visitIdx := Index{0, 0}
	for {
		data.CalculateDistance(visitIdx)
		if data.AllVisited() {
			break
		}
		visitIdx = data.GetSmallIdx()
	}
	// data.PrintTable()
	fmt.Println("The distance from {0,0} to the last point is: ", data.v[data.row][data.col].distance)

}

func parseData(day int) *VertexTable {
	file, _ := ioutil.ReadFile(fmt.Sprintf("day%d/input.txt", day))
	// file, _ := ioutil.ReadFile(fmt.Sprintf("day%d/testinput.txt", day))
	// file, _ := ioutil.ReadFile(fmt.Sprintf("day%d/test1.txt", day))
	data := strings.Split(string(file), "\n")

	vt := &VertexTable{
		make([][]Vertex, len(data)-1, len(data)-1),
		len(data) - 2,
		len(data[0]) - 1,
	}

	for lineIdx, d := range data {
		if len(d) > 0 {
			tmp := make([]Vertex, len(d))
			for i, v := range d {
				tmp[i].weight, _ = strconv.Atoi(string(v))
			}
			vt.v[lineIdx] = tmp
		}
	}

	return vt
}

/* I've implemented the solution for part 1 using Dijkstra's Shortest Path Algorithem. I don't know much about algorithems but I watched this (https://www.youtube.com/watch?v=pVfj6mxhdMw) video and implemented the algorithem by myself.*/
func main() {
	data := parseData(15)
	part1(data)
}
