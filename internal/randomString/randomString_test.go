package randomString

import "testing"

/* RandomString(20) */

func TestRandomString(t *testing.T) {
	// test RandomString
	s := RandomString(20)
	if len(s) != 20 {
		t.Errorf("RandomString(20) failed, got %s", s)
	}
}

func BenchmarkRandStringBytesMaskImprSrcUnsafe(b *testing.B) {
	// run the RandomString function b.N times
	for n := 0; n < b.N; n++ {
		_ = RandomString(20)
	}
}
