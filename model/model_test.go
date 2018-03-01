package model

import "testing"

func TestRandomString(t *testing.T) {

	want := 100
	got := len(RandomString(want))
	if got != want {
		t.Errorf("model.RandomString err, got %d,want %d", got, want)
	}
}

func BenchmarkRandomString(b *testing.B) {
	b.ReportAllocs()
	want := 1000
	for i := 0; i < b.N; i++ {
		got := len(RandomString(want))
		if got != want {
			b.Errorf("model.RandomString err, got %d,want %d", got, want)
		}
	}
}
