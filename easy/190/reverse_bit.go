package easy

func reverseBits(num uint32) uint32 {
	var bit uint32 = 0
	for i := 0; i < 32; i++ {
		bit = bit<<1 | num&1
		num >>= 1
	}
	return bit
}
