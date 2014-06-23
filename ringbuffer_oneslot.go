package ringbuffer

import (
	"fmt"
)

func (rb *ringBufferOneSlot) Free() {
	// The size of the element array is the ring buffer size + 1
	// which adds one slot which is always empty.
	rb.Elements = make([]bufferElement, rb.Size + 1)
	
	log.Debug("Ring buffer cleared - size [%d]", len(rb.Elements))
}


func (rb *ringBufferOneSlot) IsFull() bool{
	return rb.Start == rb.next(rb.End)
}


func (rb *ringBufferOneSlot) IsEmpty() bool{
	return rb.Start == rb.End
}

// Next returns the next int index for the int c_index provided
// panics if index provided is out of bounds
func (rb *ringBufferOneSlot) next(c_index int) int{
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
	
	// If index is full we move the End(write) and Start(read) indexes
	// else we just move the End(write) index
	if rb.IsFull() {
		if rb.ClearFlag {
			rb.clearAtIndex(&rb.Start)
		}
		rb.Start = rb.next(rb.Start)
		if rb.debug { log.Debug("Ring buffer is full incremented Read index to [%d]", rb.Start) }
	}
	
	// Write the value to the current End(write index)
	rb.Elements[rb.End] = e
	if rb.debug { log.Debug("Writing [%v] to the index [%d]", e.GetValue(), rb.End) }
	
	rb.End = rb.next(rb.End)
}

func (rb *ringBufferOneSlot) Read() bufferElement{
	if rb.IsEmpty() {
		e := NewNilElement()
		
		if rb.debug { log.Debug("Read nil value ring buffer is empty") }
		return e
	} else {
		e := rb.Elements[rb.Start]
		if rb.debug { log.Debug("Read [%v] from index [%d]", e.GetValue(), rb.Start) }
		if rb.ClearFlag {
			rb.clearAtIndex(&rb.Start)
		}
		rb.Start = rb.next(rb.Start)		
		if rb.debug { log.Debug("Incremented Read index to [%d]", rb.Start) }
		return e
	}
}

func (rb *ringBufferOneSlot) clearAtIndex(index *int) {
	rb.Elements[*index] = NewNilElement()
	// log.Debug("Cleared Read value from index [%d]", index)
}
