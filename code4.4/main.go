package main

import "fmt"

func main() {
	fmt.Println(GCD(200, 6)) // 200と６の最大公約数をユークリッドの互助法で求める
}

func GCD(m int, n int) int {
	if n == 0 {
		return m
	}
	return GCD(n, m%n)
}
