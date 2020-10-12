package main
import "fmt"

func swap(a *int, b *int){
  c := *a
  *a = *b
  *b = c
}

func main(){
  x := 10
  y := 12
  swap(&x, &y)
  fmt.Printf("X is %d , Y is %d\n", x, y)
}
