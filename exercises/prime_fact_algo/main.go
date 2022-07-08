package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// go build -o prime main.go
// ./prime 100 (give some number)

func main(){

	arr :=[]int{}

	if os.Args[1] != "" { // must enter a number.
		
		n,err := strconv.Atoi(os.Args[1])
		if err != nil{

			log.Fatal(err)
		}
		p:=2
	
		fmt.Print(n,"'s prime factorization is => ") // print selected number

		if n > 0 {
			for n >= p*p {
				if n%p == 0 {
					fmt.Print(p," * ") // print factorization values
					arr = append(arr, p)
					n/=p
				}else{
					p+=1
				}
			}
			arr = append(arr, n)
			if len(arr) == 1{
				fmt.Println(n," It is a prime number") // print if selected number is prime
			}else{
				fmt.Print(n)	// print all factorization values
			}
			
		}
	}
}