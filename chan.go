package main 


import (

   "fmt"
   "sync"
)

func main() {
	done := make(chan int)

	go func() {
			fmt.Println("Hello,World")
			done <- 1
		}()

		<-done



	var mu sync.Mutex

	mu.Lock()

	go func() {
				fmt.Println("Hello,World Againest")
				mu.Unlock()
		}()

	mu.Lock()
}