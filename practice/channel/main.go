package main

import (
  "fmt"
  "time"
)

func ccounter( in string, channel chan string ) {
	for {
		fmt.Println("this functions step is " + in);
		sum := 0
		step := 1
  		for i := 0; i < 5; i++ {
    		sum = sum + step
    		fmt.Println(step, i, sum );
    		time.Sleep(10 * time.Millisecond)
  		}
  		channel <- in
	}
}

func main() {

  fmt.Println( "invoke Channel A" )

  channelList := make([]chan string, 2)
　for i, _ := range channelList { channelList[i] = make(chan string) }

  for _, c := range channelList { go ccounter("a", channel1) }
  

  for {
  	select {
        case c1 := <- channel1:
        	fmt.Println(c1)
        	time.Sleep(1000 * time.Millisecond)
        case c2 := <- channel2:
        	fmt.Println(c2)
        	time.Sleep(1000 * time.Millisecond)
    }
    
  }
}
