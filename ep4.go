package main
import "fmt"

func main(){
  fmt.Print("Enter Farhnhight: ")
  var input float64
  fmt.Scanf("%f", &input)
  output := (input - 32.0) * 5 / 9
  fmt.Printf("%f\n", output)
}
