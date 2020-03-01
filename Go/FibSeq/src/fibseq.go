package main
import "fmt"

func main(){
	var n int = 30
	var x, y, z int = 0, 0, 0

	fmt.Printf("Xn = Xn-1 + Xn-2\n");

	var i int
	for i=0; i<n+1; i++{
		if(i==0){
			x = 0
		}else if(i==1){
			y = 1
		}else{
			z = y + x
			x = y
			y = z
		}
	}

	fmt.Printf("n = %d, Xn = %d\n", n, z)
}
