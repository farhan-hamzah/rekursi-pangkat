// Menghitung Pangkat
package main
import "fmt"
func pangkat(a, b int)int{
	if b <= 0{
		return 1
	}
	return a*  pangkat(a, b-1)
}

func main(){
	var x, y int
	fmt.Scan(&x, &y)
	fmt.Print(pangkat(x, y))
}
