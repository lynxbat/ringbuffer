// Package ringbuffer provides a ring(circular) buffer with options on size, element type, and circular buffer type.
package ringbuffer

import "fmt"

// RingBufferType(s):
// 		Default / OneSlot = Uses a ring buffer that maintains one slot open. Good for larger ring buffers with small elements
//		FillCount = TODO
// 		MirrorBit = TODO
// 		RWCount = TODO
// 		Absolute = TODO
// 		LastOp = TODO
const (
	Default RingBufferType = iota // 0
	OneSlot RingBufferType = Default // == Default
	FillCount RingBufferType = iota // 1
	MirrorBit RingBufferType = iota // 2
	RWCount RingBufferType = iota // 3
	Absolute RingBufferType = iota // 4
	LastOp RingBufferType = iota // 5
)

// RingBufferType is a type of ring buffer.
type RingBufferType int

// An interface for a buffer element. Used by built-in element types.
type bufferElement interface {
	GetValue() elementValueType
	WriteValue(elementValueType)
}

// TODO
type elementValueType interface {}

// TODO
type nilElement struct {
}

// Returns a nilElement
func NewNilElement() bufferElement{
	be := new(nilElement)
	return be
}

func (ne *nilElement) GetValue() elementValueType{
	return ne
}

// No-op Write method to satisfy BufferElement interface for nilElement.
func (ie *nilElement) WriteValue(value elementValueType) {
	// Do nothing
}

// Returns a integerElement containing an int provided in value
func NewIntegerElement(value int) bufferElement{
	be := new(integerElement)
	be.Value = value
	return be
}

// Integer value based buffer element
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




// ringBufferOneSlot

type ringBufferOneSlot struct {
	Size int
	Start int
	End int
	Elements []bufferElement
	ClearFlag bool
}

// Returns a ring buffer of the int size and RingBufferType provided
func NewRingBuffer(size int, rb_type RingBufferType) *ringBufferOneSlot{
	rb := new(ringBufferOneSlot)
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

func (rb *ringBufferOneSlot) Free() {
	// The size of the element array is the ring buffer size + 1
	// which adds one slot which is always empty.
	rb.Elements = make([]bufferElement, rb.Size + 1)
}




func (rb *ringBufferOneSlot) IsFull() bool{
	return rb.Start == rb.next(rb.End)
}


func (rb *ringBufferOneSlot) IsEmpty() bool{
	return rb.Start == rb.End
}

func (rb *ringBufferOneSlot) Write(e bufferElement) {	
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

func (rb *ringBufferOneSlot) Read() bufferElement{
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
func (rb *ringBufferOneSlot) next(c_index int) int{
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
