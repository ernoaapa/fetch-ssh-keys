package format

import (
	"testing"

	log "github.com/Sirupsen/logrus"

	"github.com/stretchr/testify/assert"
)

func TestSshWithoutComment(t *testing.T) {
	log.SetLevel(log.DebugLevel)

	keys := map[string][]string{
		"ernoaapa": {
			"ssh-rsa AAAAB3NzsshPublicKeyBlah",
		},
	}

	result := ssh(keys, "")

	assert.Equal(t, "ssh-rsa AAAAB3NzsshPublicKeyBlah ernoaapa\n", result, "Returned invalid ssh output")
}

func TestSshWithComment(t *testing.T) {
	log.SetLevel(log.DebugLevel)

	keys := map[string][]string{
		"ernoaapa": {
			"ssh-rsa AAAAB3NzsshPublicKeyBlah",
		},
	}

	result := ssh(keys, "Generated file")

	assert.Equal(t, "# Generated file\nssh-rsa AAAAB3NzsshPublicKeyBlah ernoaapa\n# Generated file\n", result, "Returned invalid ssh output")
}
