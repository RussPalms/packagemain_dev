package fuzztestingingo

import "testing"

func TestEqual(t *testing.T) {
	if !Equal([]byte{'f', 'u', 'z', 'z'}, []byte{'f', 'u', 'z', 'z'}) {
		t.Error("expected true, got false")
	}
}

func FuzzEqual(f *testing.F) {
	f.Add([]byte{'f', 'u', 'z', 'z'}, []byte{'t', 'e', 's', 't'})

	f.Fuzz(func(t *testing.T, a []byte, b []byte) {
		Equal(a, b)
	})
}
