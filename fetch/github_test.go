package fetch

import (
	"os"
	"testing"

	log "github.com/Sirupsen/logrus"

	"github.com/stretchr/testify/assert"
)

func TestFetchOrganisationKeys(t *testing.T) {
	log.SetLevel(log.DebugLevel)

	keys, err := GitHubOrganisationKeys("devopsfinland", GithubFetchParams{
		// Use token if it's available to avoid hitting API rate limits with the tests...
		Token:             os.Getenv("GITHUB_TOKEN"),
		PublicMembersOnly: true,
	})

	assert.NoError(t, err, "Fetch GitHub keys returned error")
	assert.True(t, len(keys) > 0, "should return SSH at least one public key")
	assert.True(t, len(keys["ernoaapa"]) > 0, "should return ernoaapa public SSH key")
	assert.True(t, len(keys["ernoaapa"][0]) > 0, "should not return empty key for ernoaapa")
}

func TestFetchUserKeys(t *testing.T) {
	log.SetLevel(log.DebugLevel)

	keys, err := GitHubUsers([]string{"ernoaapa", "arnested"}, os.Getenv("GITHUB_TOKEN"))

	assert.NoError(t, err, "Fetch GitHub keys returned error")
	assert.Equal(t, 2, len(keys), "should return SSH keys for both users")
	assert.True(t, len(keys["ernoaapa"]) > 0, "should return ernoaapa public SSH key")
	assert.True(t, len(keys["ernoaapa"][0]) > 0, "should not return empty key for ernoaapa")
}
