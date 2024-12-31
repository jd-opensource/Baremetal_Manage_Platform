package util

func Bool2Int8(b bool) int8 {
	if b {
		return 1
	}
	return 0
}

func Int82Bool(i int8) bool {
	return i > 0
}
