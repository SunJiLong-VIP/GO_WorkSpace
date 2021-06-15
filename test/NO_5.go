package main

import (
	"fmt"
)

func main0501() {
	count := 0
	for i := 0; i <= 20; i++ {
		for j := 0; j <= 33; j++ {
			for k := 0; k <= 100; k += 3 {
				count++
				if i+j+k == 100 && i*5+j*3+k/3 == 100 {
					fmt.Printf("i %d j %d k %d\n", i, j, k)
				}
			}
		}
	}
	fmt.Println(count)
}

func main() {
	count := 0
	c := 0
	for a := 0; a <= 20; a++ {
		for b := 0; b <= 33; b++ {
			count++
			c = 100 - a - b
			if a*5+b*3+c/3 == 100 && c%3 == 0 {
				fmt.Printf("a %d  b%d  c %d\n", a, b, c)
			}
		}
	}
	fmt.Println(count)
}
