package main

import (
  "fmt"
  "os"
  "time"
  "strings"
)

var moves int = 1 // Add 1 for exit
var debug_wait = 0 * time.Millisecond
var loop_wait = 0 * time.Millisecond

const (
  clearScreen = "\033[2J"
  moveToTop  = "\033[H"
  clearLine  = "\033[K"
)

const (
  up = 0
  right = 1
  down = 2
  left = 3
  done = 4
)

func getStart(board [][]byte) (int, int, int) {
  for y, line := range board {
    for x, char := range line {
      if char == '^' {
        return x, y, up
      } else if char == '<' {
        return x, y, left 
      } else if char == '>' {
        return x, y, right 
      } else if char == 'v' {
        return x, y, down
      }
    }
  }
  return 0, 0, down
}

func move(board [][]byte, x int, y int, dir int) (int, int, int) {
  if(dir == down) {
    if(len(board) > y + 1) {
      if(board[y + 1][x] == '#') {
        return move(board, x, y, (dir + 1) % 4)
      }
      nextChar := board[y + 1][x]
      if(nextChar == '.') {
        moves ++
      } else {
        time.Sleep(debug_wait)
      }
      board[y][x] = 'x'
      board[y + 1][x] = 'v'
      return x, y + 1, dir
    } else {
      return x, y, done
    }
  } else if (dir == up) {
    if(y - 1 >= 0) {
      if(board[y - 1][x] == '#') {
        return move(board, x, y, (dir + 1) % 4)
      }
      nextChar := board[y - 1][x]
      if(nextChar == '.') {
        moves ++
      } else {
        time.Sleep(debug_wait)
      }
      board[y][x] = 'x'
      board[y - 1][x] = '^'
      return x, y - 1, dir
    } else {
      return x, y, done
    }
  } else if (dir == left) {
    if(x - 1 >= 0) {
      if(board[y][x - 1] == '#') {
        return move(board, x, y, (dir + 1) % 4)
      }
      nextChar := board[y][x - 1]
      if(nextChar == '.') {
        moves ++
      } else {
        time.Sleep(debug_wait)
      }
      board[y][x] = 'x'
      board[y][x - 1] = '<'
      return x - 1, y, dir
    } else {
      return x, y, done
    }
  } else if (dir == right) {
    if(x + 1 < len(board[y])) {
      if(board[y][x + 1] == '#') {
        return move(board, x, y, (dir + 1) % 4)
      }
      nextChar := board[y][x + 1]
      if(nextChar == '.') {
        moves ++
      } else {
        time.Sleep(debug_wait)
      }
      board[y][x] = 'x'
      board[y][x + 1] = '>'
      return x + 1, y, dir
    } else {
      return x, y, done
    }
  }
  return x, y, done
}

func printBoard(board [][]byte, x int, y int) {
  yPrintFrom := max(0, y - 20)
  yPrintTo := min(len(board) - 1, y + 20)

  partialBoard := board[yPrintFrom:yPrintTo]
  partialLines := []string{}
  for _, line := range partialBoard {
    partialLines = append(partialLines, string(line))
  }

  fmt.Println(strings.Join(partialLines, "\n"))
  fmt.Println("Unique moves: %v", moves)
}

func main() {
  input, err := os.ReadFile("input.txt")
  if err != nil {
    return
  }
  boardStr := strings.Split(string(input), "\n")
  board := make([][]byte, 0, len(boardStr))
  for _, s := range boardStr {
    if s != "" {
      board = append(board, []byte(s))
    }
  }

  x, y, dir := getStart(board)
  fmt.Print(clearScreen)
  for {
    fmt.Print(moveToTop)
    x, y, dir = move(board, x, y, dir)
    if(dir == done) {
      fmt.Print(clearScreen)
      fmt.Printf("%v moves completed before exiting!", moves)
      return 
    }
    printBoard(board, x, y)
    time.Sleep(loop_wait)
  }
}
