/*
 Random inputs, with usual suspects like MAXINT +/- 1, powers of two,
 format strings etc... 
*/

package main

import (
	"math/rand"
	"time"
	"bytes"
)

func http() string {
	// typical and frequently used buffer sizes
	magicMultipliers := [...]int{32, 64, 128, 256, 512, 1024, 2048, 4096, 8192, 32767, 65535}

/*
	// typical sizes that can trigger off-by-ones etc
	magicNumbers := [...]int{
		0, -1, 1,
		-127, -128, -129, 127, 128, 129,
		-32766, -32767, -32768, 32766, 32767, 32768,
		-65534, -65535, -65536, 65534, 65535, 65536,
	}
*/
	// parser bugs and format strings
	magicStrings := []string{"%s%n%x", "%", ",", "", " "}

	/* DEBUG TEST */
	httpHeaders := []string{
		"Accept",
		"Accept-Charset",
		"Accept-Encoding",
		"Accept-Language",
		"Alternates",
		"Authorization",
		"Connection",
		"Content-Length",
		"Content-Type",
		"Cookie",
		"Cookie2",
		"DNT",
		"Forwarded",
		"Host",
		"If-Match",
		"If-Modified-Since",
		"If-None-Match",
		"If-Range",
		"If-Unmodified-Since",
		"Keep-Alive",
		"Max-Forwards",
		"Negotiate",
		"None",
		"Pragma",
		"Proxy-Authorization",
		"Proxy-Connection",
		"Range",
		"Referer",
		"Request-Range",
		"TCN",
		"TE",
		"Transfer-Encoding",
		"Upgrade",
		"User-Agent",
		"Via",
		"X-Forwarded-For",
		"X-Forwarded-Host",
	}


	var buffer bytes.Buffer
	rand.Seed(time.Now().UnixNano())

	// this part generates up to 7 headers filled with random string of magicString length
	for header := 0; header<= 7; header++ {
		buffer.WriteString(httpHeaders[rand.Intn(len(httpHeaders))])
		buffer.WriteString(": ")
		for x := 0; x <= magicMultipliers[rand.Intn(len(magicMultipliers))] + 3; x++{
			buffer.WriteString(magicStrings[rand.Intn(len(magicStrings))])
		}
		buffer.WriteString("\n")
	}
	return buffer.String()
	/* DEBUG TEST */


}
