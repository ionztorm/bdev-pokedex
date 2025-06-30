package utils

import (
	"fmt"
	"io"
	"os"
)

func SafeClose(c io.Closer) {
	if err := c.Close(); err != nil {
		fmt.Fprintf(os.Stderr, "error closing resource: %v\n", err)
	}
}
