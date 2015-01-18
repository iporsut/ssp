// example add program by my wife Kanokon Chansom https://github.com/ihippoi
package main

import "fmt"

func addNumber(number1, number2 int) (result int) {
	return number1 + number2
}
func main() {
	var number1 int
	var number2 int
	fmt.Scanf("%d %d", &number1, &number2)
	var result = addNumber(number1, number2)
	fmt.Println(result)
}
