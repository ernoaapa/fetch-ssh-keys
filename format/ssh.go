package format

import (
	"bytes"
	"fmt"
)

// ssh produces output in .ssh/authorized_keys compatible format
func ssh(keysByUsername map[string][]string, comment string) string {
	stringBuffer := bytes.NewBufferString("")

	if len(comment) > 0 {
		stringBuffer.WriteString(fmt.Sprintf("# %s\n", comment))
	}

	for username, keys := range keysByUsername {
		for _, key := range keys {
			stringBuffer.WriteString(fmt.Sprintf("%s %s\n", key, username))
		}
	}

	if len(comment) > 0 {
		stringBuffer.WriteString(fmt.Sprintf("# %s\n", comment))
	}

	return stringBuffer.String()
}
