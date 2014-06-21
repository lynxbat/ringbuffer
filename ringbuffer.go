package ringbuffer

import "fmt"


/*
A read moves the Start index.

The Start existing on the same index as the End means the buffer is empty with nothing to be read. 
* An empty buffer means nothing is read (return null) and no indexes move.

************
S = 1
E = 1
[ ] [ ] [ ] [ ] [ ]
************

A write moves the End index.
* If the index is full a write moves the End and Start index

The Start index equaling the End index + 1 means the buffer is full.

************
S = 2
E = 1
[4] [ ] [1] [2] [3]
************

*/

/*
	Implementation of Circular Buffer using always keep one slot open pattern
	to differentiate between full and empty.
*/

/*

TODO

1. Interface for elements, allowing of X type of element
2. tests tests tests
3. Higher level lib file
4. Interface for RingBuffer, type this one to be of the one slot open (spare slot) type
5. Choice of Memory(non-persistent/fast) or File(persistent/slower) backed ringbuffer


*/

// BufferElement
// Takes a BufferValue interface


type bufferElement interface {
	GetValue() elementValueType
	WriteValue(elementValueType)
}

type elementValueType interface {}



//

type nilElement struct {
}

func NewNilElement() bufferElement{
	be := new(nilElement)
	return be
}

func (ne *nilElement) GetValue() elementValueType{
	return ne
}

func (ie *nilElement) WriteValue(value elementValueType) {
	// Do nothing
}

// integerElement

func NewIntegerElement(value int) bufferElement{
	be := new(integerElement)
	be.Value = value
	return be
}

type integerElement struct {
	Value int
}

func (ie *integerElement) GetValue() elementValueType{
	return ie.Value
}

func (ie *integerElement) WriteValue(value elementValueType) {
	i := value.(int) 
	ie.Value = i
}




// ringBuffer

type ringBuffer struct {
	Size int
	Start int
	End int
	Elements []bufferElement
	ClearFlag bool
}

func NewRingBuffer(size int) *ringBuffer{
	rb := new(ringBuffer)
	rb.Size = size
	rb.Start = 0
	rb.End = 0
	// Enables writing nil values back to an index
	// Can be very expensive with high frequency writes and reads (13-20% overhead)
	// You can enable this if you want easier debugging of ringbuffer structure as
	// it will replace read or blank slot values with a nilElement type.
	rb.ClearFlag = false
	rb.Free()
	return rb
}

func (rb *ringBuffer) Free() {
	// The size of the element array is the ring buffer size + 1
	// which adds one slot which is always empty.
	rb.Elements = make([]bufferElement, rb.Size + 1)
}




func (rb *ringBuffer) IsFull() bool{
	return rb.Start == rb.next(rb.End)
}


func (rb *ringBuffer) IsEmpty() bool{
	return rb.Start == rb.End
}

func (rb *ringBuffer) DebugPrint() {
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

func (rb *ringBuffer) Write(e bufferElement) {	
	// TODO MAKE THREADSAFE
	
	// Write the value to the current End(write index)
	rb.Elements[rb.End] = e
	rb.End = rb.next(rb.End)
	
	// If index is full we move the End(write) and Start(read) indexes
	// else we just move the End(write) index
	if rb.IsFull() {
		if rb.ClearFlag {
			rb.Elements[rb.Start] = NewNilElement()
		}
		rb.Start = rb.next(rb.Start)
	}
	
}

func (rb *ringBuffer) Read() bufferElement{
	if rb.IsEmpty() {
		e := NewNilElement()
		return e
	} else {
		e := rb.Elements[rb.Start]
		if rb.ClearFlag {
			rb.Elements[rb.Start] = NewNilElement()
		}
		rb.Start = rb.next(rb.Start)		
		return e
	}
}

// Private

// Next returns the next int index for the int c_index provided
// panics if index provided is out of bounds
func (rb *ringBuffer) next(c_index int) int{
	// Ensure index is within the element bounds. Just a bit of go ahead and panic since we would panic.
	if c_index + 1 > rb.Size + 1 {
		panic(fmt.Sprintf("A index value was passed that is out of bounds [%d]. Max is [%d]", c_index, rb.Size - 1))
	}
	// Check the existing index and return new index
	if c_index + 1 == rb.Size + 1 {
		// If at the end of the index we start over at 0 and return
		return 0
	} else {
		// Otherwise we increment the index int and return
		return c_index + 1
	}
}
