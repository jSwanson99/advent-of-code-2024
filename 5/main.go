package main

import (
  "strconv"
  "fmt"
  "os"
  "strings"
  "slices"
)

func ParseRules(input string) map[string][]string {
  rules := strings.Split(input, "\n")
  rulesMap := map[string][]string{}
  for _, rule := range rules {
    parts := strings.Split(rule, "|")
    if(len(parts) != 2) {
      fmt.Printf("Invalid input %v\n", rule)
      continue
    }
    before, after := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
    rulesMap[before] = append(rulesMap[before], after)
  }
  return rulesMap
}

func IsOrdered(input []string, rules map[string][]string) []int {
  for i, value := range input {
    shouldBeAfter := rules[value]
    before := input[:i]

    for j, x := range before {
      if slices.Contains(shouldBeAfter, x) {
        fmt.Printf("record fails because %v should be after %v\n", x, value)
        return []int{j, i}
      }
    }
  }
  fmt.Printf("record is OK\n")
  return nil
}

func main() {
  originalBytes, err := os.ReadFile("input.txt")
  if err!= nil {
    fmt.Printf("Error parsing file %v", err)
    return
  }

  arr := strings.Split(string(originalBytes), "\n\n")
  rules := ParseRules(arr[0])
  for key, rule := range rules {
    fmt.Printf("%v: %v\n", key, rule)
  }

  records := strings.Split(arr[1], "\n")

  partOne := 0
  partTwo := 0
  for _, r := range records {
    didSwap := false

    if(r == "") {
      continue
    } 

    record := strings.Split(r, ",")
    fmt.Printf("\n=== record %v\n", record)
    toSwap := IsOrdered(record, rules)
    if(toSwap != nil) {
      i, j := toSwap[0], toSwap[1]
      fmt.Printf("swapping %v & %v\n", record[i], record[j])
      temp := record[i]
      record = slices.Delete(record, i, i+1)
      record = slices.Insert(record, j, temp)
      didSwap = true 
      fmt.Printf("-- result %v\n", record)
    }

    length := len(record)
    if(length == 0 || length % 2 == 0) {
      fmt.Printf("Skipped invalid %v", record)
      continue
    }

    middle, err := strconv.Atoi(record[(length - 1) / 2])
    if err != nil {
      continue 
    }

    if(didSwap) {
      partTwo += middle
    } else {
      partOne += middle
    }
  }
  fmt.Printf("Part One: %v\n", partOne)
  fmt.Printf("Part Two: %v\n", partTwo)
}
