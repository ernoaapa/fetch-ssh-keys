package format

import log "github.com/Sirupsen/logrus"

// Build builds output in given formatName format
func Build(formatName string, keysByUsername map[string][]string) string {
	switch formatName {
	case "ssh":
		return ssh(keysByUsername)
	}
	log.Fatalf("Invalid output format name: %s", formatName)
	return ""
}
