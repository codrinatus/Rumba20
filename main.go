package main

import (
	"encoding/binary"
	"fmt"
)

// Constants for Rumba20
const (
	// You might adjust constants as per the actual Rumba20 spec
	NumRounds = 20
	BlockSize = 64
)

// RotateLeft performs a left bitwise rotation
func RotateLeft(x, n uint32) uint32 {
	return (x << n) | (x >> (32 - n))
}

// RumbaQuarterRound applies a quarter round of the Rumba20 algorithm
func RumbaQuarterRound(state *[16]uint32, a, b, c, d int) {
	state[a] += state[b]
	state[d] ^= state[a]
	state[d] = RotateLeft(state[d], 16)

	state[c] += state[d]
	state[b] ^= state[c]
	state[b] = RotateLeft(state[b], 12)

	state[a] += state[b]
	state[d] ^= state[a]
	state[d] = RotateLeft(state[d], 8)

	state[c] += state[d]
	state[b] ^= state[c]
	state[b] = RotateLeft(state[b], 7)
}

// Rumba20Block generates a block of Rumba20 output
func Rumba20Block(key [8]uint32, nonce [3]uint32, blockCounter uint32) [BlockSize]byte {
	var state [16]uint32
	// Initialize state: key, constants, nonce
	state[0], state[1], state[2], state[3] = 0x61707865, 0x3320646e, 0x79622d32, 0x6b206574 // Constants
	state[4], state[5], state[6], state[7] = key[0], key[1], key[2], key[3]
	state[8], state[9], state[10], state[11] = key[4], key[5], key[6], key[7]
	state[12] = blockCounter
	state[13], state[14], state[15] = nonce[0], nonce[1], nonce[2]

	// Make a copy of the state for scrambling
	workingState := state

	// Perform Rumba20 rounds
	for i := 0; i < NumRounds; i += 2 {
		// Column rounds
		RumbaQuarterRound(&workingState, 0, 4, 8, 12)
		RumbaQuarterRound(&workingState, 1, 5, 9, 13)
		RumbaQuarterRound(&workingState, 2, 6, 10, 14)
		RumbaQuarterRound(&workingState, 3, 7, 11, 15)

		// Diagonal rounds
		RumbaQuarterRound(&workingState, 0, 5, 10, 15)
		RumbaQuarterRound(&workingState, 1, 6, 11, 12)
		RumbaQuarterRound(&workingState, 2, 7, 8, 13)
		RumbaQuarterRound(&workingState, 3, 4, 9, 14)
	}

	// Add original state to scrambled state
	for i := range state {
		workingState[i] += state[i]
	}

	// Serialize the state to bytes
	var block [BlockSize]byte
	for i := 0; i < 16; i++ {
		binary.LittleEndian.PutUint32(block[i*4:], workingState[i])
	}

	return block
}

func main() {
	// Example key, nonce, and block counter
	key := [8]uint32{0x11111111, 0x22222222, 0x33333333, 0x44444444, 0x55555555, 0x66666666, 0x77777777, 0x88888888}
	nonce := [3]uint32{0x12345678, 0x9abcdef0, 0xdeadbeef}
	blockCounter := uint32(1)

	block := Rumba20Block(key, nonce, blockCounter)

	fmt.Printf("Rumba20 Block: %x\n", block)
}
