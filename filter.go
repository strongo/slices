package slices // TODO: Move to a dedicated GitHub repository

func FilterInt64s(source, remove []int64) (result []int64, changed bool) {
	result = source[:0]
	for _, s := range source {
		for _, r := range remove {
			if s == r {
				changed = true
				goto next
			}
		}
		result = append(result, s)
	next:
	}
	return
}

func RemoveStrings(source, remove []string) (result []string, removed int) {
	result = make([]string, 0, len(source)-len(remove))
sourceRange:
	for _, s := range source {
		for _, r := range remove {
			if s == r {
				removed++
				continue sourceRange
			}
		}
		result = append(result, s)
	}
	return
}

func IsInStringSlice(v string, s []string) bool {
	for _, sv := range s {
		if sv == v {
			return true
		}
	}
	return false
}

func IsInInt64Slice(v int64, s []int64) bool {
	for _, sv := range s {
		if sv == v {
			return true
		}
	}
	return false
}
