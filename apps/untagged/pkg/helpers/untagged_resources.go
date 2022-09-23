package helpers

func UntaggedResources(resources map[string][]string) ([]string, error) {
	untagged := []string{}
	for key, tags := range resources {
		if len(tags) == 0 {
			untagged = append(untagged, key)
		}
	}

	return untagged, nil
}
