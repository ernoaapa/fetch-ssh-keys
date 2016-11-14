package output

import "github.com/ernoaapa/fetch-ssh-keys/format"

// Writer is interface for all output writers
type Writer interface {
	write(output string) error
}

// Write writes keys to given outputName in given formatName
func Write(formatName, target string, keysByUsername map[string][]string) error {
	writer := getWriter(target)
	return writer.write(format.Build(formatName, keysByUsername))
}

func getWriter(target string) Writer {
	switch target {
	case "":
		return &StdoutWriter{}
	default:
		return NewFileWriter(target)
	}
}
