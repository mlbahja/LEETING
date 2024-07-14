package main

import "fmt"

func main() {
    var arr = []int{1, 2, 3, 4}
	slice := []int{}
	for i := 0; i < 2; i++ {
		slice = append(slice,arr...)
	}
	fmt.Println(slice)
}
func  main(){
	var arr = []int{1,2,3}
	n := len(arr)
	slice := make([]int,n*2)
	for i := 0; i < n;i++ {
		slice[i],slice[i+n] = arr[i], arr[i]
	} 
	fmt.Println(slice)

}