/*
 format strings etc...
*/

package main

import (
	"bytes"
	"math/rand"
)

type randomPotStruct struct {
	smallInt  int // 0-10
	mediumInt int // 0-255
	bigInt    int // 0-65535
	//	hugeInt      int	// -2**32-2**32
	magicInt     int    // magicMultiplier
	singleChar   byte   // single byte char
	multipleChar string // single char multipled N times
	formatString string // format string
}

var randomPot randomPotStruct // global, modified every N ms

func rubik() {
	// typical and frequently used buffer sizes
	magicMultipliers := [...]int{32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 32767, 65535}

	// typical sizes that can trigger off-by-ones etc
	/*
		magicNumbers := [...]int{
			0, -1, 1,
			-127, -128, -129, 127, 128, 129,
			-32766, -32767, -32768, 32766, 32767, 32768,
			-65534, -65535, -65536, 65534, 65535, 65536,
		}
	*/

	// parser bugs and format strings
	magicStrings := []string{"%s%n%x", "%", ",", "", " "}

	// fill in randomPot
	randomPot.smallInt = rand.Intn(10)
	randomPot.mediumInt = rand.Intn(256)
	randomPot.bigInt = rand.Intn(65535)
	randomPot.magicInt = magicMultipliers[rand.Intn(len(magicMultipliers))]
	randomPot.singleChar = byte(rand.Intn(95) + 32) // ASCII printables >32<128
	var tempBuffer bytes.Buffer
	for i := 0; i <= randomPot.bigInt; i++ {
		tempBuffer.WriteByte(randomPot.singleChar)
	}
	randomPot.multipleChar = tempBuffer.String()
	tempBuffer.Reset()

	// Format string generator
	for i := 0; i <= randomPot.magicInt+randomPot.mediumInt; i++ {
		tempBuffer.WriteString(magicStrings[rand.Intn(len(magicStrings))])
	}
	randomPot.formatString = tempBuffer.String()
}
