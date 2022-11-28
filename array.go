package main
import ("fmt")

func main() {

//   var arr3 = []string{"A"}
//   arr2 := [5]int{4,5,6,7,8}

//   fmt.Println(arr3[0])
//   fmt.Println(arr2)


  arr1 := [4]string{"Volvo", "BMW", "Ford", "Mazda"}
  arr2 := [...]int{1,2,3,4,5,6}

  fmt.Println(len(arr1),arr1[0])
  fmt.Println(len(arr2))
}