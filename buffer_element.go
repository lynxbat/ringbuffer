package ringbuffer

// An interface for a buffer element. Used by built-in element types.
type bufferElement interface {
	GetValue() elementValueType
	WriteValue(elementValueType)
}

// ElementType allows for different value types for GetValue() and WriteValue()
type elementValueType interface {}

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