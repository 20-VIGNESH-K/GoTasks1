package main
import "fmt"

type Queue struct{
	item []int
}
func (q *Queue)Enqueue(n int){
	q.item=append(q.item,n)
}
func (q *Queue)Dequeue()int{
	i:= q.item[0]
   q.item= q.item[1:]  //used to remove items
   return i
}
func (q *Queue) IsEmpty() bool {
	return len(q.item) == 0
 }
func main(){
	fmt.Println("Enqueue and Dequeue the elements(FIFO):")
	q := &Queue{}  // create a queue instance which will be used to enqueue and dequeue elements
	for i:=1;i<=10;i++{
		q.Enqueue(i)
	}
	for !q.IsEmpty() {
	   fmt.Println(q.Dequeue())  //check whether the queue is empty or not
	}
}
