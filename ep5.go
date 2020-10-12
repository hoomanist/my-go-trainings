package main
import "fmt"

func main() {
  fmt.Print("Enter feet: ")
  var input float64
  fmt.Scanf("%f", &input)
  fmt.Printf("%f\n", input * 0.3048)
}
