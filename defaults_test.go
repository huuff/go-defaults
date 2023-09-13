package defaults

import (
  "testing"
)

func TestDefaultString(t *testing.T) {
  result := DefaultString("", "test") 

  if result != "test" {
    t.Fatalf("Expected: test, Got: %s", result)
  }
}


func TestInputString(t *testing.T) {
  result := DefaultString("test", "default")

  if result != "test" {
    t.Fatalf("Expected: test, Got: %s", result)
  }
}

func TestDefaultPtr(t *testing.T) {
  value := new(string)
  *value = "test"
  result := DefaultPtr(nil, value)

  if *result != *value {
    t.Fatalf("Expected %s, Got: %s", *value, *result)
  }
}

func TestInputPtr(t *testing.T) {
  def := new(string)
  *def = "default"
  input := new(string)
  *input = "input"
  result := DefaultPtr(input, def)

  if *result != *input {
    t.Fatalf("Expected %s, Got: %s", *input, *result)
  }
}
