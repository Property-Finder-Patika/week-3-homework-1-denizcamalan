package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main(){

	arr :=[]int{}

	if os.Args[1] != "" { // You must enter a number.
		
		n,err := strconv.Atoi(os.Args[1])
		if err != nil{

			log.Println("Given argument is not a number")
		}
		p:=2
	
		fmt.Print(n,"'s prime factorization is => ")
		
		if n > 0 {
			for n >= p*p {
				if n%p == 0 {
					fmt.Print(p," * ")
					arr = append(arr, p)
					n/=p
				}else{
					p+=1
				}
			}
			arr = append(arr, n)
			if len(arr) == 1{
				fmt.Println(n," It is a prime number")
			}else{
				fmt.Print(n)
			}
			
		}
	}
}