package functions

import "errors"

func FindKeyByValue[K comparable, V comparable](mp map[K]V, value V) (K, error) {
	for k, v := range mp {
		if v == value {
			return k, nil
		}
	}
	var key K
	return key, errors.New("no such key")
}

func AbsInt(n int) int {
	if n > 0 {
		return n
	}
	return -n
}
