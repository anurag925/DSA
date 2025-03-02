package dsa

func reverseBits(num uint32) uint32 {
	powers := make(map[int]uint32)

	for i := 0; i <= 31; i++ {
		powers[i] = 1 << i
	}
	res := uint32(0)
	j := 31
	for i := 0; i < 32; i++ {
		if num&1<<i == 1 {
			res += powers[j-i]
		}
	}
	return res
}
