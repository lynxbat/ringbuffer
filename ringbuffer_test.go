package ringbuffer

import "testing"

func TestWrite(t *testing.T) {
	rb := NewRingBuffer(8)
	for x := 1; x <= 16; x ++ {
		ie := NewIntegerElement(x)
		rb.Write(ie)
	}	
}

func BenchmarkWrite(b *testing.B) {
	size := 128
	rb := NewRingBuffer(size)
	
	// Wrap several times
	for x := 1; x <= 768; x ++ {
		ie := NewIntegerElement(x)
		rb.Write(ie)
	}	
	
	// Benchmark the writing of one element
	for i := 0; i < b.N; i++ {
		ie := NewIntegerElement(i)
		rb.Write(ie)
	}
}

func BenchmarkRead(b *testing.B) {
	rb := NewRingBuffer(128)	
	for x := 1; x <= 1024; x ++ {
		ie := NewIntegerElement(x)
		rb.Write(ie)
	}	
	
	// Benchmark the reading of one element
	for i := 0; i < b.N; i++ {	
		rb.Read()
	}
}