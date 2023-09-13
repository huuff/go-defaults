package defaults 

func DefaultString(input, defaultValue string) string {
  if input == "" {
    return defaultValue
  }  

  return input
}


func DefaultPtr[T any](input, defaultValue *T) *T {
  if input == nil {
    return defaultValue
  }

  return input
}
