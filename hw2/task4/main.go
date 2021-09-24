package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 *rot13Reader) Read(b []byte) (int, error) {
	n, err := rot13.r.Read(b)
	if err != nil {
		return 0, err
	}
	for i := 0; i < n; i++ {
		if b[i] >= 'A' && b[i] <= 'Z' {
			if b[i] >= 'A'+13 {
				b[i] -= 13
				continue
			}
			b[i] += 13
		}
		if b[i] >= 'a' && b[i] <= 'z' {
			if b[i] >= 'a'+13 {
				b[i] -= 13
				continue
			}
			b[i] += 13
		}
	}
	return n, nil
}
func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
