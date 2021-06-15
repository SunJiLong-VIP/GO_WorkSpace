package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main0801() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now().Nanosecond()
	for i := 0; i < 1000; i++ {
		fmt.Println(rand.Intn(100))
	}
	end := time.Now().Nanosecond()
	fmt.Println(end - start)
}

func main() {
	//使用随机数种子
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		//用数组切片定义红色球
		red := make([]int, 6)
		//随机取6个不重复的红色球
		for i := 0; i < len(red); i++ {
			//取1-33号的随机数
			temp := rand.Intn(33) + 1
			//去掉重复的数字
			for j := 0; j < i; j++ {
				if red[j] == temp {
					temp = rand.Intn(16) + 1
					j = -1
				}
			}
			//将取出的随机数赋值给red
			red[i] = temp
		}
		//输出6个红色球号（1-33）+1个蓝色球号（1-16）
		fmt.Println(" 红色球：", red, "\t\t\t", "蓝色球：", rand.Intn(16)+1)
	}

}
