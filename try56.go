package main

import(
  "fmt"
)

func swap(x,y int)(x1,y1 int){
  y1,x1 = x,y
  return
}

func main(){
  n, m := swap(10, 20)
	fmt.Println(n, m)
}
