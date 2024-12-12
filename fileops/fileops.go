package fileops

import ("fmt"
        "os"
        "strconv"
        "errors")

func ReadFromFile(fileName string,defaultValue float64) (float64, error) {
  data, err := os.ReadFile(fileName)
  if err != nil {
    return defaultValue, errors.New("Failed to read file")
  }
  valueText := string(data)
  value, err := strconv.ParseFloat(valueText, 64)
  if err != nil {
    return defaultValue, errors.New("Failed to parse file")
  }
  return value, nil
}

func WriteBalanceToFile(value float64,fileName string) {
  valueText := fmt.Sprintf("%f", value)
  os.WriteFile(fileName, []byte(valueText), 0644)
}