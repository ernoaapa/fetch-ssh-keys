package output

import (
	"github.com/youtube/vitess/go/ioutil2"
	"os"
)

// FileWriter writes the output to file
type FileWriter struct {
	targetFile string
	fileMode   os.FileMode
}

// NewFileWriter creates new FileWriter what writes to targetFile
func NewFileWriter(targetFile string, perm os.FileMode) *FileWriter {
	return &FileWriter{
		targetFile: targetFile,
		fileMode:   perm,
	}
}

func (w *FileWriter) write(output string) error {
	return ioutil2.WriteFileAtomic(w.targetFile, []byte(output), w.fileMode)
}
