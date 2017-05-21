package main

import (
	"fmt"
	"log"
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

/*the real sort algorithm*/
func sort(arr *[]int){
	array:=*arr

	/*reverses an array, needed for min1*/
	reverse := func(array []int) []int {
		retrun:=make([]int,len(array))

		index:=0
		for i:=len(array)-1;i>=0;i-- {
			retrun[i]=array[index]
			index++
		}
		return retrun
	}
		
	/*minimizes an array to the first half and returns it*/
	min1 := func(array []int) []int {
		retrun:= reverse(array)
		if len(array)%2==0 {
			return reverse(retrun[len(array)/2:])
		} else {
			return reverse(retrun[len(array)/2+1:])
		}
	}

	/*minimizes an array to the second half and returns it*/
	min2 := func(array []int) []int {
		return array[len(array)/2:]
	}
	
	lfchild, rtchild := min1(array), min2(array)
	
	/*if the childs are no single object arrays they need to be divided again*/
	if len(lfchild) != 1 {
		sort(&lfchild)
	}
	if len(rtchild) != 1 {
		sort(&rtchild)
	}

	/*CONQUER*/
	/*after dividing the small arrays get reconnected to bigger ones*/
	
	lfci, rtci, arri := 0, 0, 0
			
	for lfci<len(lfchild) && rtci<len(rtchild) {
		if lfchild[lfci]<=rtchild[rtci] {
			array[arri]=lfchild[lfci]
			lfci++
		} else {
			array[arri]=rtchild[rtci]
			rtci++
		}
		arri++
	}
	if lfci>len(lfchild) {
		for i,v := range rtchild {
			if i>=rtci {
				array[arri]=v
				arri++
			}
		}
	} else {
		for i,v := range lfchild {
			if i>=lfci {
				array[arri]=v
				arri++
			}
		}
	}
	arr=&array
}

/* ************************ MAIN ************************ */

func main(){
	var array []int

	fmt.Println("********** INSERTION SORT **********")
	fmt.Printf("\nhow many numbers do you want to sort? ")

	if    _, err := fmt.Scan(&size);    err != nil {
	    log.Print("  Scan for size failed, due to ", err)
	    return
	}

	fillmethod:=""
	fmt.Printf("\nfill the array with random numbers [Y/n] ")
	if num, _ := fmt.Scanln(&fillmethod);fillmethod=="y" || fillmethod =="Y" || num==0 {
		array=fillrand()
	} else {
		array=fill()
	}
	
	fmt.Println(array)
	sort(&array)
	fmt.Println(array)
	
	fmt.Println()
}
