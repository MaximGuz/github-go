//Программа получает на вход трехзначное число и выводит сумму его чисел
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := n/100 + n/10%10 + n%10
	fmt.Println(a)
}
