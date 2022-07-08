package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

// go build -o sieve main.go
// ./sieve 29 (give some number)

var prime []int
var conv float64

func main(){

	if os.Args[1] != "" { // You must enter a max n to determine prime silce lenght and values.
		
		n,err := strconv.Atoi(os.Args[1])
		if err != nil{
			log.Println("Given argument is not a number")
		}
		prime = make([]int, n-1)

		for index := range prime {
			prime[index] = index+2
		}
		fmt.Println(FindPrime(n))
	}
}

// this function finds and creates prime number of slice
func FindPrime(n int) []int{
	index :=0
	for _,value := range prime {
		for index=0; index < n-1 ; index++{
			conv = float64(prime[index]) 
			// siece of erosthenes rules
			if prime[index] > value && (prime[index]%value == 0 && float64(value) <= math.Sqrt(conv)) { 
				prime = RemoveIndex(prime,index)
				n--
			}
		}
	}
	return prime
}

// this func removes not prime number values in the slice
func RemoveIndex(arr []int, index int) []int {
    return append(arr[:index], arr[index+1:]...)
}

