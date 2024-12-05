package utils 

import (
  "regexp"
)

func CountOccurances(input string) int {
  re := regexp.MustCompile(`XMAS`)
  return len(re.FindAll([]byte(input), -1))
}

func Transpose(input []string) []string {
  result := []string{}
  height := len(input)
  if(height < 1) {
    return input
  }
  width := len(input[0])

  for i:=0; i < width; i ++ {
    cur := ""
    for j:=0; j < height; j ++ {
      cur += string(input[j][i])
    }
    result = append(result, cur)
  }

  return result
}

func Diagonals(input []string) []string {
  result := []string{}
  height := len(input)
  width  := len(input[0])

    for row := 0; row < height; row++ {
        diagonal := ""
        i, j := row, 0
        for i < height && j < width {
            diagonal += string(input[i][j])
            i++
            j++
        }
        result = append(result, diagonal)
    }

    // Diagonals going down-right, starting from top row (except first column)
    for col := 1; col < width; col++ {
        diagonal := ""
        i, j := 0, col
        for i < height && j < width {
            diagonal += string(input[i][j])
            i++
            j++
        }
        result = append(result, diagonal)
    }

    // Diagonals going down-left, starting from rightmost column
    for row := 0; row < height; row++ {
        diagonal := ""
        i, j := row, width-1
        for i < height && j >= 0 {
            diagonal += string(input[i][j])
            i++
            j--
        }
        result = append(result, diagonal)
    }

    // Diagonals going down-left, starting from top row (except last column)
    for col := width-2; col >= 0; col-- {
        diagonal := ""
        i, j := 0, col
        for i < height && j >= 0 {
            diagonal += string(input[i][j])
            i++
            j--
        }
        result = append(result, diagonal)
    }

  return result
}

func CeresSearch(lines []string, re *regexp.Regexp) int {
  sum := 0
	linesToTest := append([]string{}, lines...)
	linesToTest  = append(linesToTest, Transpose(lines)...)
	linesToTest  = append(linesToTest, Diagonals(lines)...)

	for _, line := range linesToTest {
		sum += len(re.FindAll([]byte(line), -1))
	}

  return sum
}

