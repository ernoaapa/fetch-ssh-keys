package output

import (
	"github.com/ernoaapa/fetch-ssh-keys/format"
	"os"
)

// Writer is interface for all output writers
type Writer interface {
	write(output string) error
}

// Write writes keys to given outputName in given formatName
func Write(formatName, target string, perm os.FileMode, keysByUsername map[string][]string) error {
	writer := getWriter(target, perm)
	return writer.write(format.Build(formatName, keysByUsername))
}

func getWriter(target string, perm os.FileMode) Writer {
	switch target {
	case "":
		return &StdoutWriter{}
	default:
		return NewFileWriter(target, perm)
	}
}
