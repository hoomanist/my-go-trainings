package main

import "fmt"


func average(args ...float64) float64{
  total := 0.00
  for _, v := range args {
    total += v
  }
  return total / float64(len(args))
}

func main(){
  fmt.Println(average(20,19, 19))
}
