package main

import "fmt"

func main0401() {
	var a int
	for a = 1; a <= 10; a++ {
		fmt.Println(a)
	}
}

func main0402() {
	sum := 0
	for a := 2; a <= 100; a += 2 {
		sum += a
		fmt.Println(a, sum)
	}
	fmt.Println(sum)
}

func main0403() {
	for i := 100; i < 1000; i++ {
		a := i / 100
		b := i / 10 % 10
		c := i % 10
		if a*a*a+b*b*b+c*c*c == i {
			fmt.Println(i)
		}
	}
}

func main0404() {
	for a := 100; a < 1000; a++ {
		q := a / 100
		w := a / 10 % 10
		e := a % 10
		if q*q*q+w*w*w+e*e*e == a {
			fmt.Println(a)
		}
	}
}

func main0405() {
	for i := 1; i <= 100; i++ {
		if i%7 == 0 || i/10 == 7 || i%10 == 7 {
			fmt.Println("敲桌子", i)
		} //else {
		//	fmt.Println(i)
		//}
	}
}

func main0406() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, j*i)
			//if j==i{
			//	break
			//}
		}
		fmt.Println()
	}
}

func main() {
	for i := 0; i <= 100; i++ {
		if i%2 == 1 {
			continue
		}
		fmt.Println(i)
	}
}
