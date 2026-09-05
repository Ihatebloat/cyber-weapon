package main
import (
"fmt"
"os"
"strconv"
)
func main() {
if len(os.Args) < 3 {
fmt.Println("Usage: ./cyber-weapon <number1> <number2>")
return
}
num1, err1 := strconv.Atoi(os.Args[1])
num2, err2 := strconv.Atoi(os.Args[2])
if err1 != nil || err2 != nil {
fmt.Println("cyber-weapon: error: nil biliberda")
return
}
result := num1 + num2
fmt.Printf("Result: %d\n", result)
}
