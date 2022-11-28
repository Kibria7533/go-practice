package main
import ("fmt")

func main() {

	myslice := []int {1,2,3}
	fmt.Println(cap(myslice))

	arr1 := [6]int{10, 11, 12, 13, 14,15}
  myslice2 := arr1[2:4]

  fmt.Printf("myslice = %v\n", myslice)
  fmt.Printf("length = %d\n", len(myslice2))
  fmt.Printf("capacity = %d\n", cap(myslice))

}