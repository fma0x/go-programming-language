package comma

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	return s[:n-3] + "," + s[n-3:]
}

func buffer_comma(values []int) string {
	var b bytes.Buffer

	n := len(values)
	if n <= 3 {
		fmt.Fprintf(&b, "%d", values)
		return b.String()
	}
	return "1234566"
}
