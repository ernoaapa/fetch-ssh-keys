package output

import (
	"io/ioutil"
	"os"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestFile(t *testing.T) {
	log.SetLevel(log.DebugLevel)

	file, createErr := ioutil.TempFile("", "example")
	assert.NoError(t, createErr, "Unable to create temp file")

	writer := NewFileWriter(file.Name(), 0640)

	writer.write("foobar")
	writer.write("foobar-second-time")

	fileBytes, readErr := ioutil.ReadFile(file.Name())
	assert.NoError(t, readErr, "Unable to read temp file")

	info, _ := os.Stat(file.Name())
	assert.Equal(t, 0640, int(info.Mode().Perm()), "File permission wasn't set as expected")
	assert.Equal(t, "foobar-second-time", string(fileBytes), "FileWriter didnt wrote expected output to file")
}

func TestFailIfNoDirectory(t *testing.T) {
	log.SetLevel(log.DebugLevel)

	writer := NewFileWriter("do/not/exist/my-file.txt", 0640)

	err := writer.write("foobar")
	assert.Error(t, err, "Did not return error if directory doesn't exist")
}
