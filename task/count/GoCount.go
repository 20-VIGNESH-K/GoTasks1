/*package main

import "fmt"

func main(){
	str:="go in language called go lang"
	var count,i int
	count=0
    length:=len(str)
	for i=0;i<length-1;i++{
		if str[i]=='g' && str[i+1]=='o'{
			count+=1
		}
	}
	fmt.Println("Count is ",count)
}*/

/*package main

import (
   "fmt"
   "strings"
   "bufio"
   "os"
)

func main(){
    scanner:=bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str:=scanner.Text()
	ans:=strings.Count(str,"go")
	fmt.Println("Count is ",ans)
}*/

package main

import (
   "fmt"
   "strings"
   "bufio"
   "os"
)

func main(){
    scanner:=bufio.NewReader(os.Stdin)
	str,_:=scanner.ReadString('\n')
	ans:=strings.Count(str,"go")
	fmt.Println("Count is ",ans)
	//fmt.Println(i)
}



