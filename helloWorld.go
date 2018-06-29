package main

import(
    "fmt"
    "math/rand"
    "time"
)

type T struct {
	name  string // name
	value int    // b

}

func Sum(a *[3]float64) (sum float64) {
    for _, v := range *a {
        sum += v
    }
    return
}




func random(min, max int) int {
    rand.Seed(time.Now().UTC().UnixNano())
    return rand.Intn(max - min) + min
}

func add( x int, y int) (z int) {
	z = x+y;
	return
}

func main() {
    myrand := random(1, 6)
	fmt.Println(myrand)
	myrand = random(1, 6)
	fmt.Println(myrand)
	fmt.Println(add(2,3))
	sum := 0
	for i:=0; i< 15 ; i++ {
		sum += i
	}
	fmt.Println(sum)	
}
