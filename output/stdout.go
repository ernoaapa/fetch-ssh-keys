package output

import "fmt"

// StdoutWriter just writes the output to stdout
type StdoutWriter struct{}

func (w *StdoutWriter) write(output string) error {
	fmt.Println(output)
	return nil
}
