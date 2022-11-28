package main
import ("fmt")

func main() {

// 	myslice := []int {1,2,3}
// 	fmt.Println(cap(myslice))

// 	arr1 := [6]int{10, 11, 12, 13, 14,15}
//   myslice2 := arr1[2:4]

//   fmt.Printf("myslice = %v\n", myslice)
//   fmt.Printf("length = %d\n", len(myslice2))
//   fmt.Printf("capacity = %d\n", cap(myslice))

myslice1 := []int{1, 2, 3, 4, 5, 6}
  fmt.Printf("myslice1 = %v\n", myslice1)
  fmt.Printf("length = %d\n", len(myslice1))
  fmt.Printf("capacity = %d\n", cap(myslice1))

  myslice1 = append(myslice1, 20, 21)
  fmt.Printf("myslice1 = %v\n", myslice1)
  fmt.Printf("length = %d\n", len(myslice1))
  fmt.Printf("capacity = %d\n", len(myslice1))

}