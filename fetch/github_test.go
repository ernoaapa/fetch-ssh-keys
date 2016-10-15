package fetch

import (
	"testing"

	log "github.com/Sirupsen/logrus"

	"github.com/stretchr/testify/assert"
)

func TestFetchPublicKeys(t *testing.T) {
	log.SetLevel(log.DebugLevel)

	keys, err := GitHubKeys("devopsfinland", GithubFetchParams{PublicMembersOnly: true})

	assert.NoError(t, err, "Fetch GitHub keys returned error")
	assert.True(t, len(keys) > 0, "should return SSH at least one public key")
	assert.True(t, len(keys["ernoaapa"]) > 0, "should return ernoaapa public SSH key")
	assert.True(t, len(keys["ernoaapa"][0]) > 0, "should not return empty key for ernoaapa")
}
