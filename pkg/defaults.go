package defaults 

func DefaultString(input, defaultValue string) string {
  if input == "" {
    return defaultValue
  }  

  return input
}
