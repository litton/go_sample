package main 

import (

  "fmt"
)



func main() {
	done := make(chan int,10)

	for i:= 0; i< cap(done);i++ {
		i := i
		go func() {
			fmt.Printf("nihao,world=%d\n",i)
			done <- 1
		}()
	}


	for i := 0;i < cap(done);i++ {
		<- done
	}
}


