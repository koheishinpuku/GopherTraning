package main
import(
  "fmt"
)

type MyInt int

func (self *MyInt) Inc(){
  *self++
  fmt.Println(self)
}

func main() {
	var n MyInt
  n = 1
  fmt.Sprintf("%x ", n)
	n.Inc()
	fmt.Println(n)

}
