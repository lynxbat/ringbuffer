package main

import (
	"github.com/lynxbat/ringbuffer"
	"fmt"	
)

func main() {
	size := 7
	to_write := 12
	
	rb := ringbuffer.NewRingBuffer(size, ringbuffer.Default, true, true, 0)
	fmt.Printf("**Initial ring buffer [size: %d]**\n", size)
	
	fmt.Printf("**%d sequence writes**\n", to_write)
	// Loop to_write times writing a new integer element	
	for x := 1; x <= to_write; x ++ {	
		ie := ringbuffer.NewIntegerElement(x)
		rb.Write(ie)
		fmt.Printf("%d ", x)
	}
	fmt.Print("\n\n")
	
	fmt.Print("**Read out all values**\n")
	for !rb.IsEmpty() {
		y := rb.Read().GetValue()
		fmt.Printf("%d ", y)
	}
	fmt.Print("\n\n")
}