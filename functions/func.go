package functions

import (
	"fmt"
	// "math"
)

func add(a, b int) int {
	return a + b
}

func reverse(a, b int) (int, int) {
	return b, a
}

func swap(a, b string) (string, string) {
	return b, a
}

// ? Named return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return //? Naked return 
}

func main() {
	fmt.Println("Welcome functions module")
	fmt.Println(add(5, 3))
	fmt.Println(reverse(1, 49))

	var a, b string = "test 1", "test 2"
	
	fmt.Println(a, b)
	fmt.Println(swap("1_first word", "2_second word"))
	fmt.Println(split(200))

}
