package main 

import (

   "fmt"
)


var limit = make(chan int ,3)

func main() {
	for _,2 := range work {
		go func() {
				limit <- 1
				w()
				<- limit
			}()
	}

	select{}
}