package main

import (
	"fmt"
)

func main0301() {
	var year int
	fmt.Println("请输入年份：")
	fmt.Scan(&year)
	fmt.Println(year%400 == 0 || (year%4 == 0 && year%100 != 0))
}

func main0302() {
	var so int
	fmt.Printf("请输入成绩：")
	fmt.Scan(&so)
	if so > 700 {
		fmt.Println("我上清华")
		if so > 750 {
			fmt.Println("学习GO、GO、GO")
		} else if so > 740 {
			fmt.Println("学习计算机")
		} else if so > 730 {
			fmt.Println("学习挖掘机")
		} else if so > 720 {
			fmt.Println("学习炒菜")
		} else {
			fmt.Println("学习计算机控制挖掘机炒菜")
		}
	} else if so > 650 {
		fmt.Println("我上北大")
		if so >= 680 {
			fmt.Println("我要学习 Python")
		} else if so >= 660 {
			fmt.Println("我要学习计算机控制挖掘机")
		}
	} else {
		fmt.Println("我上蓝翔学习计算机控制挖掘机炒菜")
	}
}

func main0303() {
	var a, b, c int
	fmt.Println("输入")
	fmt.Scan(&a, &b, &c)
	if a > b {
		if a > c {
			fmt.Println("a 重")
		}
	} else if b > c {
		fmt.Println("b 重")
	} else {
		fmt.Println("c 重")
	}
}

func main0304() {
	var a int
	fmt.Println("输入")
	fmt.Scan(&a)
	switch a {
	case 1:
		fmt.Println("星期一")
	case 2:
		fmt.Println("星期二")
	case 3:
		fmt.Println("星期三")
	case 4:
		fmt.Println("星期四")
	case 5:
		fmt.Println("星期五")
	case 6:
		fmt.Println("星期六")
	case 7:
		fmt.Println("星期日")
	default:
		fmt.Println("错误")

	}
}
