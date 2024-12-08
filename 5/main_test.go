package main

import (
  "testing"
  "reflect"
  "fmt"
)

func TestParseRules(t *testing.T) {
  input := `17|18
    17|19
    19|8
    8|200
    1|2`
  result := ParseRules(input)
  expected := map[string][]string{
    "17": {"18", "19"},
    "19": {"8"},
    "8": {"200"},
    "1": {"2"},
  }

  if(!reflect.DeepEqual(expected, result)) {
    t.Errorf("Expected %v but got %v for input %v", expected, result, input)
  } else {
    fmt.Printf("Expected %v got %v for input %v", expected, result, input)
  }
}

func TestIsOrdered(t *testing.T) {
  input := []string{"1", "2", "3", "4", "5"}
  result := IsOrdered(input, map[string][]string{"1": {"2"}})
  expected := true 
  
  if(expected != result) {
    t.Errorf("Expected %v but got %v for input %v", expected, result, input)
  } else {
    fmt.Printf("Expected %v got %v for input %v", expected, result, input)
  }
}

func TestIsOrdered2(t *testing.T) {
  input := []string{"1", "2", "3", "4", "5"}
  result := IsOrdered(input, map[string][]string{"2": {"1"}})
  expected := false 
  
  if(expected != result) {
    t.Errorf("Expected %v but got %v for input %v", expected, result, input)
  } else {
    fmt.Printf("Expected %v got %v for input %v", expected, result, input)
  }
}

func TestIsOrdered3(t *testing.T) {
  input := []string{"1", "2", "3", "4", "5"}
  result := IsOrdered(input, map[string][]string{"2": {"7"}})
  expected := true 
  
  if(expected != result) {
    t.Errorf("Expected %v but got %v for input %v", expected, result, input)
  } else {
    fmt.Printf("Expected %v got %v for input %v", expected, result, input)
  }
}

func TestIsOrdered4(t *testing.T) {
  input := []string{"1", "2", "3", "4", "5"}
  result := IsOrdered(input, map[string][]string{"2": {"7", "1"}})
  expected := false 
  
  if(expected != result) {
    t.Errorf("Expected %v but got %v for input %v", expected, result, input)
  } else {
    fmt.Printf("Expected %v got %v for input %v", expected, result, input)
  }
}
