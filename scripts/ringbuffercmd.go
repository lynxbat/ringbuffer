package main

import (
	"github.com/lynxbat/ringbuffer"
	"fmt"	
)

func main() {
	size := 8
	to_write := 1000
	
	rb := ringbuffer.NewRingBuffer(size)
	// rb := ringbuffer.NewRingBuffer(size, ringbuffer.Default)
	fmt.Printf("**Initial ring buffer [size: %d]**\n", size)
	rb.DebugPrint()
	
	// Loop 16 times writing a new integer element	
	for x := 1; x <= to_write; x ++ {
		// println(x)
		ie := ringbuffer.NewIntegerElement(x)
		rb.Write(ie)
	}
	
	fmt.Printf("**%d sequence writes**\n", to_write)
	rb.DebugPrint()

	fmt.Print("**Read out all values**\n")
	for !rb.IsEmpty() {
		y := rb.Read().GetValue()
		fmt.Printf("%d ", y)		
	}
	fmt.Print("\n\n")
	
	fmt.Print("**After read until empty**\n")
	rb.DebugPrint()
}