package rot8

import (
	"fmt"
	"os"
	"testing"
)

func TestRotate(t *testing.T) {
	s := "Hello, World!"
	r := Rotate(s)
	b := Rotate(r)

	if b != s {
		t.Errorf("%q != %q", b, s)
	}
}

func TestRotateFile(t *testing.T) {
	f, err := os.ReadFile("rot_test.txt")

	if err != nil {
		t.Error(err)
	}

	s := string(f)
	r := Rotate(s)
	b := Rotate(r)

	if b != s {
		t.Error("rot_test.txt")
	}
}

func BenchmarkRotate(b *testing.B) {
	f, err := os.ReadFile("rot_test.txt")

	if err != nil {
		b.Error(err)
	}

	s := string(f)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Rotate(s)
	}
}

func ExampleRotate() {
	s := "Hello, World!"
	r := Rotate(s)

	fmt.Printf("r = %q\n", r)

	// Output:
	// r = "籑籮籵籵籸簵 籠籸类籵籭簪"
}
