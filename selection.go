package main

import (
	"fmt"
	"math/rand"
)

var size int

/*generates and fills an array with userinput in a slice and returns it*/
func fill() []int{
	array := make([]int, size)
	
	fmt.Println("\nType in the numbers and seperate them by enter ...")
	
	for i:=0;i<size;i++ {
		fmt.Printf("Number %d: ",i)
		fmt.Scanf("%d",&array[i])
	}
	
	return array
}

func fillrand() []int{
	array := make([]int, size)
	
	for i:=0;i<size;i++ {
		array[i]=rand.Intn(100)
	}
	
	return array
}

func main(){
	var array []int

	fmt.Println("********** INSERTION SORT **********")
	fmt.Printf("\nhow many numbers do you want to sort? ")

	fmt.Scanln(&size)

	fillmethod:=""
	fmt.Printf("\nfill the array with random numbers [Y/n] ")
	if num, _ := fmt.Scanln(&fillmethod);fillmethod=="y" || fillmethod =="Y" || num==0 {
		array=fillrand()
	} else {
		array=fill()
	}
	
	fmt.Println(array)
	
	/*START sorting*/
	/*ranging the whole array*/	
	for i:=0;i<size-1;i++ {
		/*searching the lowest number in the array*/
		mini:=i+1
		for j:=i+1;j<size;j++ {
			if array[j]<array[mini] {
				mini=j
			}
		}
		/*swap*/
		tmp:=array[i]
		array[i]=array[mini]
		array[mini]=tmp
	}
	/*END sorting*/
	
	fmt.Println(array)
	fmt.Println()
}
