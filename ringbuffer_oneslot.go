package ringbuffer

import "fmt"

func (rb *ringBufferOneSlot) Free() {
	// The size of the element array is the ring buffer size + 1
	// which adds one slot which is always empty.
	rb.Elements = make([]bufferElement, rb.Size + 1)
}


func (rb *ringBufferOneSlot) IsFull() bool{
	return rb.Start == rb.Next(rb.End)
}


func (rb *ringBufferOneSlot) IsEmpty() bool{
	return rb.Start == rb.End
}

// Next returns the next int index for the int c_index provided
// panics if index provided is out of bounds
func (rb *ringBufferOneSlot) Next(c_index int) int{
	// Ensure index is within the element bounds. Just a bit of go ahead and panic since we would panic.
	if c_index + 1 > rb.Size + 1 {
		panic(fmt.Sprintf("A index value was passed that is out of bounds [%d]. Max is [%d]", c_index, rb.Size - 1))
	}
	// Check the existing index and return new index
	if c_index == rb.Size {
		// If at the end of the index we start over at 0 and return
		return 0
	} else {
		// Otherwise we increment the index int and return
		return c_index + 1
	}
}

func (rb *ringBufferOneSlot) Write(e bufferElement) {	
	// TODO MAKE THREADSAFE
	
	// Write the value to the current End(write index)
	rb.Elements[rb.End] = e
	rb.End = rb.Next(rb.End)
	
	// If index is full we move the End(write) and Start(read) indexes
	// else we just move the End(write) index
	if rb.IsFull() {
		if rb.ClearFlag {
			rb.Elements[rb.Start] = NewNilElement()
		}
		rb.Start = rb.Next(rb.Start)
	}
}

func (rb *ringBufferOneSlot) Read() bufferElement{
	if rb.IsEmpty() {
		e := NewNilElement()
		return e
	} else {
		e := rb.Elements[rb.Start]
		if rb.ClearFlag {
			rb.Elements[rb.Start] = NewNilElement()
		}
		rb.Start = rb.Next(rb.Start)		
		return e
	}
}



//TODO REMOVE
func (rb *ringBufferOneSlot) DebugPrint() {
	fmt.Printf(" IsFull? [%t]\n", rb.IsFull())
	fmt.Printf(" IsEmpty? [%t]\n", rb.IsEmpty())
	fmt.Printf(" StartIndex? [%d]\n", rb.Start)
	fmt.Printf(" EndIndex? [%d]\n ", rb.End)
	for i, s := range rb.Elements {
		switch e := s.(type) {
		default:
			fmt.Printf("[%#v]nil ", i)
		case *nilElement:
			fmt.Printf("[%#v]nil ", i)
		case *integerElement:
			// e := s.(*integerElement)
			fmt.Printf("[%#v]%#v ", i, e.GetValue())
		}		
	}	
	fmt.Print("\n\n")
}