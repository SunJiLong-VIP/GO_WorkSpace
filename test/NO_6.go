package main

import (
	"fmt"
)

func test2(a int, b int) {
	sum := a + b
	fmt.Println(sum)
}
func test1(a int, b int) {
	test2(a, b)
}
func main0601() {
	test1(10, 20)
}

func test4(b ...int) {
	sum := 0
	//for i := 0; i < len(b); i++ {
	//	sum+=b[i]
	//	fmt.Println("test4下标:", i, "test4数值:", b[i],"testsum:",sum)
	for _, data := range b {
		sum += data
		fmt.Println("testsum:", sum)
	}
}
func test3(a ...int) {
	//test4(a[:]...)   //下标位 0 开始 到结尾
	//test4(a[2:]...)  //下标位 2 开始 到结尾
	//test4(a[:3]...)  //下标位 0 开始，到下标位 3 结束 不含第 3 位
	test4(a[1:3]...) //下标位 1 开始，到下标位 3 结束 不含第 3 位
}
func main0602() {
	test3(0, 1, 2, 3, 4)
	test3(5, 6, 7, 8, 9)

}

func sub(a int, b int) int {
	sum := a - b
	return sum
}
func main0603() {
	//var value int
	value := sub(30, 13)
	fmt.Println("结果：", value)
}

func sut(a int, b int) int {
	return a - b
}
func main0604() {
	va := sut(55, 32)
	fmt.Println("结果：", va)
}

func sti(a int, b int) (sur int) {
	sur = a - b
	return
}
func main0605() {
	tr := sti(67, 42)
	fmt.Println("结果：", tr)
}

func test5() (a int, b int, c int) {
	a, b, c = 11, 22, 33
	return
}
func main0606() {
	a, b, c := test5()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func test6() {
	fmt.Println("你瞅啥")
}

type Ftest func()

func main0607() {
	var t Ftest
	t = test6
	t()
}

func test7() (a int, b int, c int) {
	a, b, c = 10, 20, 30
	return
}
func main0608() {
	a, b, c := test7()
	fmt.Println(test7)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(a + b + c)
}

func test8(a int, b int, c int) int {
	return a * b * c
}

// Fte type 定义函数类型
// Fte type 给已存在的函数定义别名
//type Fte func(int, int, int) int

func main0609() {
	// 定义函数类型变量
	//var Ft Fte
	// 通过函数变量调用函数
	//Ft = test8
	V := test8(10, 20, 30)
	fmt.Println(V)
	fmt.Printf("%T", V)
}

func main0610() {
	i := 0
	for ; i < 10; i++ {

	}
	fmt.Println(i)

	//程序中如果出现相同的变量名，如果函数有自己的变量，就使用自己的，如果没有就上外层寻找变量；
	//如果变量名字相同会采用就近原则，使用自己的变量；
	t := 0
	for t := 0; t < 10; t++ {
		fmt.Println(t)
	}
	fmt.Println(t)

}

//全局变量
var a int = 20

func main0611() {
	a++
	fmt.Println(a)
	a++
	fmt.Println(a)
}

func main() {
	//                *
	//              * * *
	//            * * * * *
	//          * * * * * * *
	//        * * * * * * * * *
	//      * * * * * * * * * * *
	//	  * * * * * * * * * * * * *
	//  * * * * * * * * * * * * * * *
	for i := 0; i <= 7; i++ { //  控制整体循环行数
		for j := 0; j < 7-i; j++ { //  控制每行空格数量
			fmt.Print("  ")
		}
		for k := 0; k < i*2+1; k++ { //  控制每行 * 数量
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
