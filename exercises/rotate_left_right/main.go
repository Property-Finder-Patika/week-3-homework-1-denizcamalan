package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// go build -o rotate main.go
// ./rotate right 3 (give some number)
// ./rotate left 6 (give some number)

func main(){
	arr := []int{0,1,2,3,4,5,6,7,8,9} // given slice
	new_arr := []int{}
	arr_len := len(arr)
	
	number,err := strconv.Atoi(os.Args[2]) // arg 1 is left or right & arg 2 is number of steps
	if err != nil {
		log.Fatalln(err)
	}
	if number <= 10 {	// slice lenght is 10
		if os.Args[1] == "right" {
			new_arr = append(new_arr,arr[arr_len-number:]...) 
			new_arr = append(new_arr,arr[:arr_len-number]... )
			fmt.Println(new_arr)
		}else if os.Args[1] == "left"{
			reverse(arr[:number-1])
			reverse(arr[number-1:])
			reverse(arr)
			fmt.Println(arr)
		}else{
			log.Println("wrong argument !")
		}
	}else{
		log.Println("slice bounds out of range !")
	}
}

func reverse(arr []int){	// this func reverses given slice

	for i,value := 0, len(arr)-1; i < value; i, value = i+1,value-1{
		arr[i], arr[value] = arr[value], arr[i]
	}
}