package main

import "fmt"

func main() {
	fmt.Println("Hello, Go!")

	var quantity int
	var length, width float64
	var customerName string

	quantity = 4
	length, width = 1.2, 2.4
	customerName = "Jayce"

	fmt.Println(customerName)
	fmt.Println("has ordered", quantity, "sheet")
	fmt.Println("each with an area of")
	fmt.Println(length*width, "square meters")

	// 변수 선언과 동시에 값을 할당
	var quantity2 int = 5
	var length2, width2 float64 = 1.5, 2.5
	var customerName2 string = "Jayce2"

	fmt.Println(customerName2)
	fmt.Println("has ordered", quantity2, "sheet")
	fmt.Println("each with an area of")
	fmt.Println(length2*width2, "square meters")

	// 변수 선언과 동시에 값을 할당 and 타입 생략 가능 because 할당한 값으로 자동지정
	var quantity3 = 7
	var length3, width3 = 1.7, 2.7
	var customerName3 = "Jayce3"

	fmt.Println(customerName3)
	fmt.Println("has ordered", quantity3, "sheet")
	fmt.Println("each with an area of")
	fmt.Println(length3*width3, "square meters")

	// 변수 선언과 동시에 값을 할당 and 타입 생략 가능 because 할당한 값으로 자동지정
	// 단축 변수 선언언
	quantity4 := 7
	length4, width4 := 1.7, 2.7
	customerName4 := "Jayce4"

	fmt.Println(customerName4)
	fmt.Println("has ordered", quantity4, "sheet")
	fmt.Println("each with an area of")
	fmt.Println(length4*width4, "square meters")

}
