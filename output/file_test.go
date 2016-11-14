package output

import (
	"io/ioutil"
	"testing"

	log "github.com/Sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestFile(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	log.Println("hwllllwe")

	file, createErr := ioutil.TempFile("", "example")
	assert.NoError(t, createErr, "Unable to create temp file")

	writer := NewFileWriter(file.Name())

	writer.write("foobar")
	writer.write("foobar-second-time")

	fileBytes, readErr := ioutil.ReadFile(file.Name())
	assert.NoError(t, readErr, "Unable to read temp file")

	assert.Equal(t, "foobar-second-time", string(fileBytes), "FileWriter didnt wrote expected output to file")
}
