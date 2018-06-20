package output

import (
	"os"

	"github.com/ernoaapa/fetch-ssh-keys/format"
)

// Writer is interface for all output writers
type Writer interface {
	write(output string) error
}

// Write writes keys to given outputName in given formatName
func Write(formatName, target string, perm os.FileMode, keysByUsername map[string][]string, comment string) error {
	writer := getWriter(target, perm)
	return writer.write(format.Build(formatName, keysByUsername, comment))
}

func getWriter(target string, perm os.FileMode) Writer {
	switch target {
	case "":
		return &StdoutWriter{}
	default:
		return NewFileWriter(target, perm)
	}
}
