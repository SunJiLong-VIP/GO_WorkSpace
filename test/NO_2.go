package main

import "fmt"

//func main() {
////	F1 := 5.67
//	var  F1  float64
//	const  PI = 3.14
//	fmt.Printf("F1  Type is %T\n" ,F1)
//	fmt.Scan(&F1)
//	fmt.Printf("面积：%.2f",PI*F1*F1)
//}

func main() {
	a := 'b'
	fmt.Printf("%pa\n", &a)
	fmt.Printf("*a", a)
}
