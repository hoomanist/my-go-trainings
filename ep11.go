package main

import "fmt"

func main() {
  x := make([]int 0)
  for i:=0; i < 101; i++ {
    if i % 7 == 0 {
      x = append(x, i)
    }
  }
  fmt.Println(x)
}
