package main

import "fmt"

// 值拷贝
func printFood(food [3]string){
	food[2]="大饼"
	fmt.Println(food)
}

func printFood2(food *[3]string){
	food[2]="大饼"
	fmt.Println(*food)
}


func main(){

	var array1 [6]string
	array2 := [3]string{"火锅","烧烤","家常菜"}
	array3 := [...]string{"火锅","烧烤","家常菜"}
	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array3)
	var matrix [3][4]string
	var matrix2 [3][4]int
	fmt.Println(matrix)
	fmt.Println(matrix2)

	for i:=0; i<len(array2); i++{
		fmt.Println(array2[i])
	}

	for idx, item := range(array2){
		fmt.Println(idx, item)
	}

	for _, item := range(array2){
		fmt.Println(item)
	}

	// golang 数组是值类型， 值拷贝， 会分家

	printFood(array2)
	//printFood2(&array2)
	fmt.Println(array2)



}
