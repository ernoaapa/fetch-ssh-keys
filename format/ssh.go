package format

import (
	"bytes"
	"fmt"
)

// ssh produces output in .ssh/authorized_keys compatible format
func ssh(keysByUsername map[string][]string) string {
	stringBuffer := bytes.NewBufferString("")

	for username, keys := range keysByUsername {
		for _, key := range keys {
			stringBuffer.WriteString(fmt.Sprintf("%s %s\n", key, username))
		}
	}
	return stringBuffer.String()
}
