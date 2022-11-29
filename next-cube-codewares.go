package kata

import "math"

func FindNextPower(val, pow int) int {
  r := 0
  
  for i := 1; ; i++ {
    r = int(math.Pow(float64(i), float64(pow)))
    
    if r > val {
      break
    }
  }
  
  return r
}

func main () {
	var val int
    var pow int

    fmt.Println(FindNextPower(1, 2))

}