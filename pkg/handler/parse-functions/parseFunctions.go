package parse_functions

import (
	"errors"
	"strconv"
	"strings"
)

var timeControlError = errors.New("time control parse error")

func ParseTimeControl(timeControl string) (int, int, error) {
	control := strings.Split(timeControl, "+")
	minutes, err1 := strconv.Atoi(control[0])
	increment, err2 := strconv.Atoi(control[1])
	if err1 != nil || err2 != nil {
		return 0, 0, timeControlError
	}
	return minutes, increment, nil
}
