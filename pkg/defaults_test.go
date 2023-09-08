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
