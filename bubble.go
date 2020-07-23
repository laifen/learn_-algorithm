package main

import "fmt"

func main(){
}
//如果该数组完全乱序，BubbleSort1新增的代码就显得很多余
//因此我们可以记录每次排序最后交换的位置，下一次交换到此处就可以了
func BubbleSortEnd(arr []int64) []int64 {
	if len(arr) <= 1 {
		return arr
	}

	for i := len(arr); i > 1; i-- {
		sortIndex := 0
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
				sortIndex = j
			}
		}
		i = sortIndex
		fmt.Println("i:", i)
		fmt.Println("arr:", arr)
	}
	return arr
}

//如果数组已经有序,添加一个标志位来判断
func BubbleSort1(arr []int64) []int64 {
	if len(arr) <= 1 {
		return arr
	}

	sorted := true
	for j := len(arr); j > 1; j-- {
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				temp := arr[i+1]
				arr[i+1] = arr[i]
				arr[i] = temp
				sorted = false
			}
		}
		if sorted {
			break
		}
		fmt.Println("j:", j)
		fmt.Println("arr:", arr)
	}
	return arr
}

//相邻元素比较，较小的放在前面，从小到大
func BubbleSort(arr []int64) []int64 {
	if len(arr) <= 1 {
		return arr
	}

	count := 0
	/*
		for i := 0; i < len(arr)-1; i++ {
			count += 1
			for j := i; j < len(arr)-1; j++ {
				if arr[j] > arr[j+1] {
					temp := arr[j]
					arr[j] = arr[j+1]
					arr[j+1] = temp
				}
			}
		}
	*/

	for i := 0; i < len(arr)-1; i++ {
		count += 1
		for j := len(arr) - 1; j > 0; j-- {
			if arr[j] < arr[j-1] {
				temp := arr[j-1]
				arr[j-1] = arr[j]
				arr[j] = temp
			}
		}
	}

	fmt.Println(count)
	return arr
}

//从大到小 ,应该从最后面开始比较
func BubbleSort2(arr []int64) []int64 {
	if len(arr) <= 1 {
		return arr
	}

	for i := 0; i < len(arr)-1; i++ {
		for j := len(arr) - 1; j > 0; j-- {
			if arr[j] > arr[j-1] {
				temp := arr[j-1]
				arr[j-1] = arr[j]
				arr[j] = temp
			}
		}
	}
	return arr
}
