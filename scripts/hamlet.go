package main

import (
	"io/ioutil"
	"github.com/lynxbat/ringbuffer"
	"fmt"
)

func main() {
	data, _ := ioutil.ReadFile("hamlet.txt")
	rb := ringbuffer.NewRingBuffer(935, ringbuffer.Default, true, true, 0)
	
	// sv := string(data)
	for _, element := range data {
		// println(string(element))
		be := ringbuffer.NewByteElement(element)
		rb.Write(be)
	}
	
	for !rb.IsEmpty() {
		y := rb.Read().GetValue().(byte)
		fmt.Printf("%v", string(y))
	}
	println(len(data))
}