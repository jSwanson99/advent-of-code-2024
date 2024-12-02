package main

import (
  "math"
  "sort"
  "errors"
  "net/http"
  "fmt"
  "io"
  "strings"
  "strconv"
)

func load_data() ([]byte, error) {
  url := "https://adventofcode.com/2024/day/1/input"
  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    return nil, err
  }
  req.AddCookie(&http.Cookie{
    Name: "session",
    Value: "53616c7465645f5f075ace94ed892f2f44c9517309b7bc708683aa1393e73d29fc717eeb56be40cb11eccfd787fba8560368d704995ea7092276338b7e30f170",
  })
  res, err := http.DefaultClient.Do(req)
  if (err != nil) {
    return nil, err
  }
  return io.ReadAll(res.Body)
}

func parse_lists(body []byte) ([]int, []int, error) {
  var a, b []int
  lines := strings.Split(string(body), "\n")
  for _, line := range lines {
    if line == "" {
      continue
    }
    parts := strings.Fields(line)
    // fmt.Printf("PartA: %s, PartB: %s", parts[0], parts[1])
    numA, err := strconv.Atoi(parts[0])
    if err!=nil {
      return nil, nil, err
    }
    numB, err := strconv.Atoi(parts[1])
    if err!=nil {
      return nil, nil, err
    }
    a = append(a, numA)
    b = append(b, numB)
  }
  return a, b, nil 
}

func partOne(listA []int, listB []int) (int, error) {
  var sum int
  var err error
  lenA, lenB := len(listA), len(listB)
  if(lenA != lenB) {
    return 0, errors.New("Unable to compute")
  }
  sort.Slice(listA, func(i, j int) bool { return listA[i] < listA[j]})
  sort.Slice(listB, func(i, j int) bool { return listB[i] < listB[j]})
  for i:=0; i < lenA; i ++ {
    sum += int(math.Abs(float64(listA[i] - listB[i])))
  }
  return sum, err
}

func partTwo(listA []int, listB []int) (int, error) {
  var simScore int
  var err error
  freqB := make(map[int]int)
  for _, num := range listB {
    freqB[num]++
  }
  for _, num := range listA {
    simScore += freqB[num] * num
  }
  return simScore, err
}

func main() {
  body, err := load_data()
  if(err != nil) {
    fmt.Printf("Error: %v", err)
  }

  listA, listB, err := parse_lists(body)
  if(err != nil) {
    fmt.Printf("Error: %v", err)
  }

  sum, err := partOne(listA, listB)
  if err!=nil {
    fmt.Printf("Error: %v", err)
  }

  simScore, err := partTwo(listA, listB)
  if err!=nil {
    fmt.Printf("Error: %v", err)
  }

  fmt.Printf("Answer One: %v\n", sum)
  fmt.Printf("Answer Two: %v\n", simScore)
}
