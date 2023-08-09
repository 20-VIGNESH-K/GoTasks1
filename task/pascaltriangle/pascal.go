package main

import "fmt"

func main() {
	var row, i, j int
	fmt.Scan(&row)

	var arr = make([][]int, 0)

	for i = 0; i < row; i++ {
		var ans = make([]int, 0)

		for j = 0; j <= i; j++ {
			val := 1
			if i > 1 && i > j && j > 0 {
				val = arr[i-1][j-1] + arr[i-1][j]
			}
			ans = append(ans, val)
		}
		arr = append(arr, ans)
	}
	for i := 0; i < row; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%d ", arr[i][j])
		}
		fmt.Println()
	}

}

/*package main
import "fmt"
func main(){
var n int
fmt.Scan(&n)
printPascal(n)
}
func printPascal(n int){
for line:= 1; line <= n; line++{
	 C := 1
	for i := 1; i <= line; i++{
	fmt.Printf("%d ", C)
	C = C * (line - i) / i
	}
	fmt.Println()
    }
}*/
