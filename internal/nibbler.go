package internal

type Nibbler struct {
	byte  *uint8
	bytes []uint8
	i     int
}

func NewNibbler(bytes []uint8) *Nibbler {
	nibbler := Nibbler{
		byte:  nil,
		bytes: bytes,
		i:     0,
	}
	return &nibbler
}

func (nibbler *Nibbler) Next() *uint8 {
	if nibbler.byte != nil {
		value := nibbler.byte
		nibbler.byte = nil
		return value
	}
	if nibbler.i < len(nibbler.bytes) {
		value := nibbler.bytes[nibbler.i]
		nibbler.i++
		hi := value & 0xf0
		lo := value & 0x0f
		nibbler.byte = &lo
		hi >>= 4
		return &hi
	}
	return nil
}
