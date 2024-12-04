package main

import (
  "fmt" 
  "os"
  "regexp"
  "strconv"
)

func main() {
  num_re := regexp.MustCompile(`mul\(([0-9]{1,3}), ?([0-9]{1,3})\)`)
  token_re := regexp.MustCompile(`(don?'?t?\(\)|mul\([0-9]{1,3}, ?[0-9]{1,3}\))`)

  originalBytes, err := os.ReadFile("input.txt")
  if err != nil {
    fmt.Printf("Error reading input: %v", err)
  }

  matches := token_re.FindAllStringSubmatch(string(originalBytes), -1)
  sum := 0
  skipped := 0
  lastInstr := "do()"

  for _, match := range matches {
    captured := match[1]
    fmt.Printf("Token: %v | ", captured)

    if(captured == "don't()" || captured == "do()") {
      lastInstr = captured
      continue
    } 
    num_matches := num_re.FindStringSubmatch(string(captured))
    strOne, strTwo := num_matches[1], num_matches[2]

    numOne, err := strconv.Atoi(strOne)
    if err != nil {
      fmt.Printf("Error while parsing numOne\nError: %v", err)
      return
    }
    numTwo, err:= strconv.Atoi(strTwo)
    if err != nil {
      fmt.Printf("Error while parsing numTwo\nError: %v", err)
      return
    }

    if (lastInstr != "don't()") {
      sum += numOne * numTwo
      fmt.Printf("Added %v\n", numOne * numTwo)
    } else {
      skipped += numOne * numTwo
      fmt.Printf("Skipped\n")
    }
  }
  fmt.Printf("Part One: %v\nPart Two: %v", sum + skipped, sum)
}
