package main

import "fmt"


type rect struct {
  width, hight int
}
func (r *rect) area() int {
  return r.width * r.hight
}

func main(){
  r := rect{width: 30, hight:20}
  fmt.Println(r.area())
}
