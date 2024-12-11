package main

import (
  "fmt"
  "os"
  "regexp"
  "strings"
  "strconv"
)

func numsStringToList(input string) []int {
    numbers := []int{}
    for _, str := range strings.Fields(input) {
      num, err := strconv.Atoi(str)
      if err != nil {
        fmt.Printf("ERROR")
        continue
      }
      numbers = append(numbers, num)
    }
    return numbers
}

func testEquation(target int, numbers []int) bool {

  var explore func(idx int, cur int) bool
  explore = func(idx int, cur int) bool{
    if(idx == len(numbers)) {
      return cur == target
    }

    // Pt1
    addRes := explore(idx + 1, cur + numbers[idx])
    mulRes := explore(idx + 1, cur * numbers[idx])
    // Pt2
    concatStr := strconv.Itoa(cur) + strconv.Itoa(numbers[idx])
    concatNum, err := strconv.Atoi(concatStr)
    if err != nil {
      fmt.Printf("Failed to concat")
      return false
    }
    concatRes := explore(idx + 1, concatNum)

    return addRes || mulRes || concatRes
  }
  return explore(1, numbers[0])
}

func main() {
  buff, err := os.ReadFile("input.txt")
  if err != nil {
    fmt.Printf("Cannot read file")
    return
  }

  sum := 0
  re := regexp.MustCompile(`([0-9]+)\: (.*)`) 
  matches := re.FindAllStringSubmatch(string(buff), -1)

  for _, match := range matches {
    totalStr, numbersStr := match[1], match[2]
    total, err := strconv.Atoi(totalStr)
    if err != nil {
      fmt.Printf("Cannot parse total")
      return
    }
    numbers := numsStringToList(numbersStr)

    isValid := testEquation(total, numbers)
    if(isValid) {
      sum += total
    }
  }
  fmt.Printf("Total: %v", sum)
}
