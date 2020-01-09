package output

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestStdout(t *testing.T) {
	log.SetLevel(log.DebugLevel)

	writer := &StdoutWriter{}
	writer.write("whoop")
}
