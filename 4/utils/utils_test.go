package utils

import "testing"
import "reflect"
import "slices"
import "regexp"

func TestCountOccurances(t *testing.T) {
  input := "XMASXMAS"
  expected := 2
  result := CountOccurances(input)

  if(result != expected) {
    t.Errorf("Expected %v but got %v for input %v", expected, result, input)
  }
}

func TestTranspose(t *testing.T) {
  input := []string{
    "XMASXMAS",
    "FOOBARAA",
  }
  expected := []string {
    "XF",
    "MO",
    "AO",
    "SB",
    "XA",
    "MR",
    "AA",
    "SA",
  }
  result := Transpose(input)

  if(!reflect.DeepEqual(expected, result)) {
    t.Errorf("Expected %v but got %v for input %v", expected, result, input)
  }
}

func TestDiagonals(t *testing.T) {
  input := []string{
    "XMAS",
    "XMAS",
  }
  expected := []string{
    "S",
    "AS",
    "MA",
    "XM",
    "X",
		"X",
		"MX",
		"AM",
		"SA",
		"S",
  }
  result := Diagonals(input)

	for _, diagonal := range expected {
		if !slices.Contains(result, diagonal) {
			t.Errorf("Missing %v", diagonal)
		}
	}

	if len(result) != len(expected) {
		t.Errorf("%v != %v", len(result), len(expected))
	}
}

func TestCeresSearch(t *testing.T) {
  re := regexp.MustCompile("(?=(XMAS|SAMX))")

  input := []string{"XMASXMASXMAS"}
  expected := 3
  result := CeresSearch(input, re)

  if(result != expected) {
    t.Errorf("Expected %v but got %v for input %v", expected, result, input)
  }

  input = []string{"XMASXMASXMASAMX"}
  expected = 4
  result = CeresSearch(input, re)

  if(result != expected) {
    t.Errorf("Expected %v but got %v for input %v", expected, result, input)
  }
}










