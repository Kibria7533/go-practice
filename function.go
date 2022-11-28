// package main
// import ("fmt")
// func hh(){
// 	fmt.Println("hello ");
// }
// func main() {
// 	hh();
//   fmt.Println("hello world");
// }

package main
import ("fmt")

func myFunction(x int, y int) (result int) {
  result = x + y
  return
}

func main() {
  fmt.Println(myFunction(1, 2))
}