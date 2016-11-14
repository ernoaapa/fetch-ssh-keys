package format

import (
	"testing"

	log "github.com/Sirupsen/logrus"

	"github.com/stretchr/testify/assert"
)

func TestSsh(t *testing.T) {
	log.SetLevel(log.DebugLevel)

	keys := map[string][]string{
		"ernoaapa": []string{
			"ssh-rsa AAAAB3NzsshPublicKeyBlah",
		},
	}

	result := ssh(keys)

	assert.Equal(t, "ssh-rsa AAAAB3NzsshPublicKeyBlah ernoaapa\n", result, "Returned invalid ssh output")
}
