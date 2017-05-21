package main

import (
	"fmt"
	"log"
)

var size int

func fill() []int{
	array := make([]int, size)
	
	fmt.Println("\nType in the numbers and seperate them by enter ...")
	
	for i:=0;i<size;i++ {
		fmt.Printf("Number %d: ",i)
		fmt.Scanf("%d",&array[i])
	}
	
	return array
}

func main(){
	fmt.Println("********** INSERTION SORT **********")
	fmt.Printf("\nhow many numbers do you want to sort? ")

	if    _, err := fmt.Scan(&size);    err != nil {
	    log.Print("  Scan for size failed, due to ", err)
	    return
	}


	array:=fill()
	
	fmt.Println()
	
	/*START sorting the array*/
	
	//runs through array
	for i:=1;i<size;i++ {
		//remembers the number to sort
		sort:=array[i]
		
		//from the index of the number to sort
		j:=i
		//the loop counts down
		//overwrites the number with the higher index by the number with the lower index
		//until the index to overwrite is 0 or the number with the smaller index is higher than the number to sort
		for j>0 && array[j-1]>sort {
			array[j]=array[j-1]
			j--
		}
		
		//insertion of the number to sort at the right place
		//=> at the lowest index or at the index +1 of the first lower number
		array[j]=sort
		fmt.Println(array)
	}
	
	/*END sorting the array*/
	
	fmt.Println()
}
