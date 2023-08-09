package main

import "fmt"

func main(){
	var s string
	fmt.Scan(&s)
	slice:=[]byte(s)
	length:=len(s)
	for i:=2;i<length;i+=4{
		temp:=slice[i]
		slice[i]=slice[i+1]
		slice[i+1]=temp
	}
	fmt.Println("Reversed string is ",string(slice))
}