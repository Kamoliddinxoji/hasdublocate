package main

import "fmt"

func main() {
	arr := []int{10, 5, 8, 6, 9, 4, 3, 5, 9}
	for i := 0; i < len(arr); i++ {
		num := arr[i]
		for j := 0; j < len(arr); j++ {
			if i != j {
				if num == arr[j] {
					fmt.Printf("Remove duplicate number %d \n ", arr[j])
					arr[i] = arr[len(arr)-1]
					arr = arr[:len(arr)-1]
				}
			}
		}
	}
	fmt.Println(arr)
}
