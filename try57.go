package main
import(
  "fmt"
)

func swap2(x,y *int){
  sub := *x
  *x = *y
  *y = sub
}

func main() {
	n, m := 10, 20
	swap2(&n, &m)
	fmt.Println(n, m)
}
