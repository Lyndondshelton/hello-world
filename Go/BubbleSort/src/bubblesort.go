package main
import "fmt"

func main(){
	var arr = [10]int{1,45,6,12,3, 20, 16, 11, 4,2}
	var arr_len int = len(arr)

	var i, j int
	for i = 0; i<arr_len-1; i++{
		for j = 0; j<arr_len-i-1; j++{
			if (arr[j] > arr[j+1]){
				var temp int = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}


	fmt.Printf("Array (n=%d): ", arr_len);
	var n int
	for n = 0; n<arr_len; n++{
		if (n !=9){
			fmt.Printf("%d, ", arr[n])
		}else{
			fmt.Printf("%d.", arr[n])
		}
	}
}
