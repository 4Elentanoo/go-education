package main

//? Импорты
import (
	"fmt"  //? "fmt" -> импортируем пакет fmt для форматированного ввода-вывода
	"math" //? -> импортируем пакет math для математических функций
)

func main() {
	fmt.Println("Hello, World! This is a simple Go program.")
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

}
