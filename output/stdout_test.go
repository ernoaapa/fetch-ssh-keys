package output

import (
	"testing"

	log "github.com/Sirupsen/logrus"
)

func TestStdout(t *testing.T) {
	log.SetLevel(log.DebugLevel)

	writer := &StdoutWriter{}
	writer.write("whoop")
}
