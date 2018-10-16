// +build deploy_key

package fetch

import (
	"os"
	"testing"

	log "github.com/Sirupsen/logrus"

	"github.com/stretchr/testify/assert"
)

func TestFetchDeployKeys(t *testing.T) {
	log.SetLevel(log.DebugLevel)

	keys, err := GitHubDeployKeys([]string{"arnested/fetch-ssh-keys"}, os.Getenv("GITHUB_TOKEN"))

	assert.NoError(t, err, "Fetch GitHub keys returned error")
	assert.Equal(t, 1, len(keys), "should return SSH keys for 'arnested/fetch-ssh-keys'")
	assert.True(t, len(keys["arnested/fetch-ssh-keys"]) > 0, "should return 'arnested/fetch-ssh-keys' deploy SSH key")
	assert.True(t, len(keys["arnested/fetch-ssh-keys"][0]) > 0, "should not return empty key for 'arnested/fetch-ssh-keys'")
}
