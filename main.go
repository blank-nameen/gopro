package main 
import (
	"fmt"
)

func putNum(intchan chan int) {
	
	for i := 1; i < 8000; i++ {
		intchan <- i
	}
	close(intchan)

}

func premiNum(intchan chan int , premichan chan int , exitchan chan bool) {
	//var num int
	var flag bool
	for {
		num, ok := <- intchan
		if  !ok {
			break 
		}
		flag = true 

		for i := 2; i < num; i++ {
			if num % i == 0 {
				flag = false 
				break 
			}
		}

		if flag {
			premichan<- num
		}
	}

	fmt.Println("primnum  goroutine cant recevie num  exit ")
	exitchan<- true
}

func main() {
	
	intchan := make(chan int , 1000)
	premichan := make(chan int , 1000)
	exitchan := make(chan bool , 4)
	
	go putNum(intchan)

	for i := 0; i < 4; i++{
		go premiNum(intchan, premichan, exitchan)
	}

	go func(){
		for i := 0; i < 4; i++ {
			<-exitchan
		}
		close(premichan)
	}()
	
	for {
		res, ok := <-premichan
		if !ok{
			break
		}

		fmt.Printf("sushu is %d\n", res)
	}

		fmt.Println("main  thread  exit ")

}