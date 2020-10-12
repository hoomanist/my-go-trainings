package main

import (
    "fmt"
    "math"
)

type Shape interface {
  area() float64
}

type Circle struct {
  r float64
}

type Rect struct {
  width, height float64
}

func (c Circle) area() float64 {
  return c.r * c.r * math.Pi
}

func (r Rect) area() float64 {
  return r.height * r.width
}

func size(s Shape){
  fmt.Println(s.area())
}

func main(){
  c := Circle{15}
  r := Rect{10, 15}
  size(c)
  size(r)
}
