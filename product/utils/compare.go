package utils

func CompareMaps(a map[string]string, b map[string]string) bool {
	a_keys := make([]string, 0, len(a))
	for k := range a {
		a_keys = append(a_keys, k)
	}

	b_keys := make([]string, 0, len(b))
	for k := range b {
		b_keys = append(b_keys, k)
	}

	if len(a_keys) != len(b_keys) {
		return false
	}

	for _, a_key := range a_keys {
		if a[a_key] != b[a_key] {
			return false
		}
	}

	return true
}
