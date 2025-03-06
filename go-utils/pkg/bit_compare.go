package main

func CompareBits(A int, B int) int {
	binaryRepresentation := A * B
	bitsCount := 0

	for binaryRepresentation > 0 {
		// check the bit equations
		if binaryRepresentation&1 == 1 {
			bitsCount++
		}

		// move on to the next bit
		binaryRepresentation >>= 1
	}

	return bitsCount
}
