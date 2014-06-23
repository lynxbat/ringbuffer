// Package ringbuffer provides a ring(circular) buffer with options on size, element type, and circular buffer type.
package ringbuffer

import (
	stdlog "log"
	"github.com/op/go-logging"
	"os"
)

var log = logging.MustGetLogger("ringbuffer")
var log_levels = []logging.Level{logging.CRITICAL, logging.ERROR, logging.WARNING, logging.NOTICE, logging.INFO, logging.DEBUG}

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
		
	// private
	next(int) int
	clearAtIndex(int)
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
// 	args:
//		int 			- The size of the ring buffer, default: 16
//		RingBufferType 	- The type of ring buffer to create, default: Default
//		bool 			- True/False on whether to clear empty or reserved slots (used for debug), default: false
//      bool			- True/False on enabling log messages to Stderr, default: true
//      int             - The log level 0 - 5, default: 0
func NewRingBuffer(things...interface{}) ringBuffer{	
	//	
	var size int = 16
	var ringbuffer_type RingBufferType = Default
	var clear_flag bool = true
	var log_enabled bool = true
	var log_level int = 0	
	
	// size arg
	if len(things) > 0 {
		size = things[0].(int)
	}
	
	// ringbuffer_type arg
	if len(things) > 1 {
		ringbuffer_type = things[1].(RingBufferType)
	}
	
	// clear_flag arg
	if len(things) > 2 {
		clear_flag = things[2].(bool)
	}
	
	// log_enabled arg
	if len(things) > 3 {
		log_enabled = things[3].(bool)
	}
	
	// log_level arg
	if len(things) > 4 {
		log_level = things[4].(int)
	}
	
	// log_enable init
	if log_enabled {
		initLogging(log_level)
	}
	
	// call correct type constructor
	switch ringbuffer_type {
	case Default:
		log.Debug("Default ring buffer type selected")
		return newOneSlotRingBuffer(size, clear_flag)
	default:
		panic("How did we get here?")
	}
}

func newOneSlotRingBuffer(size int, clear_flag bool) *ringBufferOneSlot{
	rb := new(ringBufferOneSlot)
	log.Debug("Created new ring buffer - type [One Slot Open] size [%d]", size)
	rb.Size = size
	rb.Start = 0
	rb.End = 0
	// Enables writing nil values back to an index
	// Can be very expensive with high frequency writes and reads (13-20% overhead)
	// You can enable this if you want easier debugging of ringbuffer structure as
	// it will replace read or blank slot values with a nilElement type.
	rb.ClearFlag = clear_flag
	log.Debug("Clear flag is [%v]", clear_flag)	
	rb.Free()
	return rb
}

func initLogging(log_level int) {
	logging.SetFormatter(logging.MustStringFormatter("▶ %{level} - %{message}"))
	logBackend := logging.NewLogBackend(os.Stderr, "", stdlog.LstdFlags|stdlog.Lshortfile)
	logBackend.Color = true
	logging.SetBackend(logBackend)
	logging.SetLevel(log_levels[log_level], "ringbuffer")
}
