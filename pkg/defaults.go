package defaults 

func defaultString(input, defaultValue string) string {
  if input == "" {
    return defaultValue
  }  

  return input
}
