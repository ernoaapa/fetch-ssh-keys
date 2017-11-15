package utils

// MergeKeys merges key maps together to single map
func MergeKeys(keySets ...map[string][]string) map[string][]string {
	result := make(map[string][]string)

	for _, userKeys := range keySets {
		for username, keys := range userKeys {
			if _, ok := result[username]; !ok {
				result[username] = []string{}
			}

			result[username] = append(result[username], keys...)
		}
	}

	return result
}
