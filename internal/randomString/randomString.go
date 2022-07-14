package randomString

import (
	"math/rand"
	"time"
	"unsafe"
)

var src = rand.NewSource(time.Now().UnixNano())

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// strings.Builder builds the string in an internal []byte, the same as we did ourselves.
// So basically doing it via a strings.Builder has some overhead,
// the only thing we switched to strings.Builder for is to avoid the final copying of the slice.
// strings.Builder avoids the final copy by using package unsafe:
// // String returns the accumulated string.
// func (b *Builder) String() string {
//     return *(*string)(unsafe.Pointer(&b.buf))
// }
// The thing is, we can also do this ourselves, too.
// So the idea here is to switch back to building the random string in a []byte,
// but when we're done, don't convert it to string to return, but do an unsafe conversion:
// obtain a string which points to our byte slice as the string data.
// https://stackoverflow.com/a/31832326/16298783
func RandomString(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&b))
}
