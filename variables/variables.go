package variables

import "fmt"

//? Модуль переменных

func main() {
	//? Строки

	var hello string
	hello = "Hello world"
	string_one := "string one"
	var string_two = "string two"
	var (
		name string
		age  string = "27"
	)
	name = "name Tom"
	fmt.Println(hello, string_one, string_two, name, age)

	//? Конец блока

	//? Целочисленные

	var a int8 = -1
	var b uint8 = 2
	var c byte = 3 // byte - синоним типа uint8
	var d int16 = -4
	var f uint16 = 5
	var g int32 = -6
	var h rune = -7 // rune - синоним типа int32
	var j uint32 = 8
	var k int64 = -9
	var l uint64 = 10
	var m int = 102
	var n uint = 105

	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("c: ", c)
	fmt.Println("d: ", d)
	fmt.Println("f: ", f)
	fmt.Println("g: ", g)
	fmt.Println("h: ", h)
	fmt.Println("j: ", j)
	fmt.Println("k: ", k)
	fmt.Println("l: ", l)
	fmt.Println("m: ", m)
	fmt.Println("n: ", n)

	//? Конец блока

	//? Дробные числа

	var ff float32 = 18
	var gg float32 = 4.5
	var dd float64 = 0.23
	var pi float64 = 3.14
	var e float64 = 2.7

	fmt.Println("f: ", ff)
	fmt.Println("g: ", gg)
	fmt.Println("d: ", dd)
	fmt.Println("pi: ", pi)
	fmt.Println("e: ", e)

	//? Конец блока

	//? Комплексные числа

	var fc complex64 = 1 + 2i
	var gc complex128 = 4 + 3i

	fmt.Println(fc, gc)
	//?

	
}
