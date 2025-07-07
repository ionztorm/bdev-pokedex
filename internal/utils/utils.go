// Package utils provides utility functions for common tasks
// such as safely closing resources and input string processing.
package utils

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func SafeClose(c io.Closer) {
	if err := c.Close(); err != nil {
		fmt.Fprintf(os.Stderr, "error closing resource: %v\n", err)
	}
}

func CleanInput(text string) []string {
	return strings.Fields(strings.TrimSpace(strings.ToLower(text)))
}
