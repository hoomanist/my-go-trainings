package main
import "fmt"

func main(){
  x := [5]float64{20, 20, 20, 20, 20}
  var total float64 = 0
  for _, value :=range x {
    total += value
  }
  fmt.Println(total/5)
}
