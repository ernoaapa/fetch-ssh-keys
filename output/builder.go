package output

import log "github.com/Sirupsen/logrus"

// Build builds output in given outputName format
func Build(outputName string, keysByUsername map[string][]string) string {
	switch outputName {
	case "ssh":
		return ssh(keysByUsername)
	}
	log.Fatalf("Invalid output name: %s", outputName)
	return ""
}
