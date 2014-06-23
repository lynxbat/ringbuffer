// Package ringbuffer provides a ring(circular) buffer with options on size, element type, and circular buffer type.
package ringbuffer



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

type ringBuffer interface {
	Free()
	IsFull() bool
	IsEmpty() bool
	Write(bufferElement)	
	Read() bufferElement
	Next(int) int
	DebugPrint()
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
func NewRingBuffer(things...interface{}) ringBuffer{	
	var size int
	var ringbuffer_type RingBufferType
	var clear_flag bool = true
	
	
	if len(things) > 0 {
		size = things[0].(int)
	} else {
		size = 16 // default to 16 slots
	}
	
	if len(things) > 1 {
		ringbuffer_type = things[1].(RingBufferType)
	} else {
		ringbuffer_type = Default
	}
	
	if len(things) > 2 {
		ringbuffer_type = things[1].(RingBufferType)
	} else {
		ringbuffer_type = Default
	}
	
	switch ringbuffer_type {
	case Default:
		println("Default")
		return newOneSlotRingBuffer(size, clear_flag)
	default:
		panic("How did we get here?")
	}
}

func newOneSlotRingBuffer(size int, clear_flag bool) *ringBufferOneSlot{
	rb := new(ringBufferOneSlot)
	rb.Size = size
	rb.Start = 0
	rb.End = 0
	// Enables writing nil values back to an index
	// Can be very expensive with high frequency writes and reads (13-20% overhead)
	// You can enable this if you want easier debugging of ringbuffer structure as
	// it will replace read or blank slot values with a nilElement type.
	rb.ClearFlag = clear_flag
	rb.Free()
	return rb
}
