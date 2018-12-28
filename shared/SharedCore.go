package shared

import (
	"strings"
)

func PersonNameParse(value string) (ret *PersonName, ok bool) {
	str := strings.Trim(value, " ")
	if len(str) > 0 {
		ok = true
		firstLast := strings.Split(str, " ")
		if len(firstLast) >= 2 {
			ret = &PersonName{First: firstLast[0], Last: strings.Join(firstLast[1:], " ")}
		} else if len(firstLast) == 1 {
			ret = &PersonName{First: firstLast[0]}
		} else {
			ret = &PersonName{First: str}
		}
	}
	return
}
