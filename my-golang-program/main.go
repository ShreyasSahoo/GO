package main

import "fmt"

// single entry point like c/cpp
func main() {
	fmt.Println("Hello from Shreyas")
	fmt.Println(countPrimes(10))

}
func countPrimes(n int) int {
	var count int = 0
	for i := 0; i < n; i++ {
		if isPrime(n) {
			count = count + 1

		}
	}
	return count
}
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
