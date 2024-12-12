package main

import (
  "fmt"
  "os"
  "slices"
)

type Point struct {
  x, y int
}

func clamp(value int) int {
  if value > 0 {
    return 1
  } else if (value < 0) {
    return -1
  }
  return 0
}
func findAntinodesPartOne(antennas map[string][]Point, maxX int, maxY int) []Point {
  antinodes := []Point{}

  for _, points := range antennas {
    for i, pointA := range points {
      for j := i + 1; j < len(points); j ++ {
        pointB := points[j]
        xDist := pointB.x - pointA.x
        yDist := pointB.y - pointA.y

        antinode := Point{
          x: pointB.x + int(xDist), 
          y: pointB.y + int(yDist),
        }
        if antinode.x >= 0 && antinode.y >=0 && antinode.x < maxX && antinode.y < maxY {
          antinodes = append(antinodes, antinode)
        }
        antinode = Point{
          x: pointA.x - int(xDist), 
          y: pointA.y - int(yDist),
        }
        if antinode.x >= 0 && antinode.y >=0 && antinode.x < maxX && antinode.y < maxY {
          antinodes = append(antinodes, antinode)
        }
      }
    }
  }
  slices.SortFunc(antinodes, func(nodeA Point, nodeB Point) int {
    if nodeA.y != nodeB.y {
      return nodeA.y - nodeB.y
    }
    return nodeA.x - nodeB.x
  })
  return slices.CompactFunc(antinodes, func(nodeA Point, nodeB Point) bool {
    return nodeA.x == nodeB.x && nodeA.y == nodeB.y
  })
}

func findAntinodes(antennas map[string][]Point, maxX int, maxY int) []Point {
  antinodes := []Point{}

  for _, points := range antennas {
    for i, pointA := range points {
      for j := i + 1; j < len(points); j ++ {
        pointB := points[j]
        xDist := pointB.x - pointA.x
        yDist := pointB.y - pointA.y

        i := 0
        for {
          antinode := Point{
            x: pointB.x + int(xDist) * i, 
            y: pointB.y + int(yDist) * i,
          }
          i ++
          if antinode.x >= 0 && antinode.y >=0 && antinode.x < maxX && antinode.y < maxY {
            antinodes = append(antinodes, antinode)
          } else {
            break
          }
        }
        i = 0
        for {
          antinode := Point{
            x: pointA.x - int(xDist) * i, 
            y: pointA.y - int(yDist) * i,
          }
          i ++
          if antinode.x >= 0 && antinode.y >=0 && antinode.x < maxX && antinode.y < maxY {
            antinodes = append(antinodes, antinode)
          } else {
            break
          }
        }
      }
    }
  }
  slices.SortFunc(antinodes, func(nodeA Point, nodeB Point) int {
    if nodeA.y != nodeB.y {
      return nodeA.y - nodeB.y
    }
    return nodeA.x - nodeB.x
  })
  return slices.CompactFunc(antinodes, func(nodeA Point, nodeB Point) bool {
    return nodeA.x == nodeB.x && nodeA.y == nodeB.y
  })
}

func findAntennas(grid [][]byte) map[string][]Point {
  antennas := make(map[string][]Point)
  for y, row := range grid {
    for x, byte := range row {
      // number || upper || lower
      if (byte >= 47 && byte <= 57) || (byte >= 65 && byte <= 90) || (byte >= 97 && byte <= 122) {
        char := string(byte)
        antennas[char] = append(antennas[char], Point{x, y})
      }
    }
  }
  return antennas
}


func parseGrid(buff []byte) [][]byte {
  start := 0
  grid := make([][]byte, 0)
  for i, byte := range buff {
    if byte == '\n' {
      row := buff[start:i]
      if i > start {
        grid = append(grid, row)
      }
      start = i + 1
    }
  }
  return grid
}

func main() {
  buff, err := os.ReadFile("input.txt")
  if err != nil {
    fmt.Printf("Error reading file")
    return
  }
  grid := parseGrid(buff)
  maxY, maxX := len(grid), len(grid[0])
  antennas := findAntennas(grid)
  antinodes := findAntinodes(antennas, maxX, maxY)

  fmt.Printf("Debug: maxX %v, maxY %v\n", maxX, maxY)
  fmt.Printf("Debug: %v\n", grid[0])
  fmt.Printf("RESULT: %v unique antinodes", len(antinodes))
}

