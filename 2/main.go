package main

import (
  "math"
  "net/http"
  "fmt"
  "io"
  "strings"
  "strconv"
)

func load_data() ([]byte, error) {
  url := "https://adventofcode.com/2024/day/2/input"
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

func parse_levels(body []byte) ([][]int, error) {
  lines := strings.Split(string(body), "\n")
  var linesCleaned []string
  for _, line := range lines {
    if(line != "") {
      linesCleaned = append(linesCleaned, line)
    }
  }

  levels := make([][]int, len(linesCleaned))
  for i, line := range linesCleaned {
    if(line == "") {
      continue
    }
    for _, _num := range strings.Fields(line) {
      num, err := strconv.Atoi(_num)
      if err != nil {
        fmt.Printf("Error parsing num %v", err)
      }
      levels[i] = append(levels[i], num)
    }
  }
  return levels, nil 
}

/**
 * returns index of problematic report, otherwise 0
 **/
func isSafe(level []int) int{
  state := "" 
  for i:=1; i < len(level); i ++ {
    isIncrease := level[i-1] < level[i]
    isDecrease := level[i-1] > level[i]
    isEqual    := level[i-1] == level[i]
    diff := int(math.Abs(float64(level[i-1] - level[i])))

    if(diff > 3) {
      return i
    } else if(state == "" && isIncrease) {
      state = "increasing"
    } else if (state == "" && isDecrease) {
      state = "decreasing"
    } else if(isIncrease && state == "decreasing") {
      return i
    } else if(isDecrease && state == "increasing") {
      return i
    } else if (isEqual) {
      return i
    }
  }
  return 0
}

func partOne(levels [][]int) (int, error) {
  var sum int
  var err error

  for _, level := range levels {
    problem := isSafe(level)
    if(problem == 0) {
      sum ++
    }
  }

  return sum, err
}

func partTwo(levels [][]int) (int, error) {
  var sum int
  var err error

  for _, level := range levels {
    problem := isSafe(level)
    if(problem == 0) {
      sum ++
      continue
    }
    for i := range level {
      retry := append([]int{}, level[:i]...)
      retry = append(retry, level[i + 1:]...)
      if(isSafe(retry) == 0) {
        sum ++
        break
      }
    }
  }

  return sum, err
}

func main() {
  data, err := load_data()
  if err != nil {
    fmt.Printf("Error: %v", err)
    return
  }

  levels, err := parse_levels(data)
  if err != nil {
    fmt.Printf("Error: %v", err)
    return
  }

  sum, err := partOne(levels)
  if err!=nil {
    fmt.Printf("Error: %v", err)
  }
  fmt.Printf("Answer One: %v\n", sum)

  sum, err = partTwo(levels)
  if err!=nil {
    fmt.Printf("Error: %v", err)
  }
  fmt.Printf("Answer Two: %v\n", sum)
}
