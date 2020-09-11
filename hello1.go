package main

import(
  "fmt"
)

func number_judge(x int)string{
  var str string
  if x%2 == 0 {
    str="-偶数uuu"
  } else {
    str="-奇数"
  }
  return str
}
func main() {
	fmt.Println(number_judge(11))
}
