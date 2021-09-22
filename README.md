# rot8000

Package rot8 provides utilities for ROT-8000 cyphering of Unicode text.

## Usage

```go
import "github.com/241m/rot8000/rot8"

func main() {
	s := "Hello, World!"
	r := rot8.Rotate(s)

	fmt.Printf("r = %q\n", r)

	// Output:
	// r = "籑籮籵籵籸簵 籠籸类籵籭簪"
}
```

## Prior art

[rottytooth/rot8000](https://github.com/rottytooth/rot8000) - original
C#/Javascript implementation.

[bits/UTF-8-Unicode-Test-Documents](https://github.com/bits/UTF-8-Unicode-Test-Documents) - Unicode BMP test file.
