package main

import (
  "fmt"
  "os"
  "strings"
	"regexp"
	"day_4/utils"
)

func partOne(lines []string) {
  sum := 0
  sum += utils.CeresSearch(lines, regexp.MustCompile(`XMAS`))
  sum += utils.CeresSearch(lines, regexp.MustCompile(`SAMX`))

  fmt.Printf("Part 1: Found %v occurances\n", sum)
}

func partTwo(lines []string) {
  sum := 0

  for i:=1; i < len(lines) - 1; i ++ {
    for j:=1; j < len(lines[i]) - 1; j ++ {
      if(lines[i][j] == 'A') {
        diagOne := string(lines[i-1][j-1]) + string(lines[i][j]) + string(lines[i+1][j+1])
        diagTwo := string(lines[i+1][j-1]) + string(lines[i][j]) + string(lines[i-1][j+1])
      
        if(diagOne == "MAS" || diagOne == "SAM") {
          if(diagTwo == "MAS" || diagTwo == "SAM") {
            sum ++
          }
        }
      }
    }
  }

  fmt.Printf("Part 2: Found %v occurances\n", sum)
}

func main() {
  originalBytes, err := os.ReadFile("input.txt")
  if err != nil {
    fmt.Printf("%v", err)
    return
  }

  lines := strings.FieldsFunc(string(originalBytes), func(c rune) bool {
    return c == '\n'
  })

  partOne(lines)
  partTwo(lines)
}
