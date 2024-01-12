package main

import "fmt"

func main() {
	Array := []int{9, 3, 2, 5, 6, 3, 4, 72}
	res:=quickSort(Array)
	fmt.Print("Sorted array: ")
	for _, v := range res {
		fmt.Print(v, " ")
	}
	fmt.Println()
}




// func quickSort(arr []int) []int{

// 	if len(arr) < 2{
// 		return arr
// 	}

// 	pivot:=arr[len(arr)-1]

// 	var left,right []int

// 	for i:=0;i<len(arr)-1;i++{
		
// 		if arr[i] < pivot{
// 			left = append(left,arr[i])
// 		} else{
// 			right = append(right,arr[i])
// 		}
// 	}

// 	return append(append(quickSort(left),pivot), quickSort(right)...)
//}

func quickSort(arr []int) []int{
	if len(arr) < 2{
		return arr
	}

	pivot:=arr[len(arr)-1]
	j:=0
	for i:=0;i<len(arr)-1;i++{
		if arr[i] < pivot{
		arr[j],arr[i] = arr[i],arr[j]
		j++
	}
}

	quickSort(arr[:j])
	quickSort(arr[j+1:])

	return arr

}
