package main

import (
	"fmt"
)

func main0701() {
	var arr [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//arr[0] = 10
	//arr[1] = 123
	//fmt.Println(arr[0] * arr[1])
	//fmt.Println(arr)
	//for i := 0; i < len(arr); i++ {
	//	fmt.Println(arr[i])
	//}
	for i, v := range arr {
		fmt.Println("下标", i, "数值", v)

	}
	fmt.Printf("%T\n\n", arr)
	b := [5]int{9, 8, 7, 6, 5}
	for i, v := range b {
		fmt.Println("下标", i, "数值", v)
	}
	fmt.Printf("%T\n\n", b)
	c := [...]int{3, 4, 5, 6, 7, 8, 9}
	for i, v := range c {
		fmt.Println("下标", i, "数值", v)
	}
	fmt.Printf("%T\n", c)

}

//数组最大值
func main0702() {
	max := 0
	arr := [10]int{3, 5, 7, 9, 2, 10, 6, 8, 4}
	for i := 0; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	fmt.Println("最大值", max)
}

//数组置换
func main0703() {
	zu := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//定义起始和结束下标
	start := 0
	end := len(zu) - 1
	for {
		if start > end {
			break
		}
		zu[start], zu[end] = zu[end], zu[start]
		start++
		end--
	}
	fmt.Println(zu)
}

//冒泡排序
func main0704() {
	ttt := [10]int{9, 6, 8, 3, 5, 2, 1, 4, 0, 7}
	//外层控制行
	for t := 0; t < len(ttt)-1; t++ {
		//内行控制列
		for y := 0; y < len(ttt)-1-t; y++ {
			if ttt[y] > ttt[y+1] {
				ttt[y], ttt[y+1] = ttt[y+1], ttt[y]
			}
		}
	}
	fmt.Println(ttt)
}

func Boll(w [10]int) [10]int {
	for o := 0; o < len(w)-1; o++ {
		for k := 0; k < len(w)-1-o; k++ {
			if w[k] > w[k+1] {
				w[k], w[k+1] = w[k+1], w[k]
			}
		}
	}
	//fmt.Println(w)
	return w
}

func main() {
	w := [10]int{4, 9, 5, 2, 1, 6, 8, 7, 0, 3}
	Boll(w)
	fmt.Printf("%T\t", &w)
	w = Boll(w)
	fmt.Println(w)
}
