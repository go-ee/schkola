package person

import (
	"strings"
	"ee/schkola/shared"
)

func (o *genders) ParseGenderGerman(name string, defaultValue *Gender) (ret *Gender, ok bool) {
	if strings.EqualFold(strings.ToLower(name), "männlich") {
		ret = Genders().Male()
		ok = true
	} else if strings.EqualFold(strings.ToLower(name), "weiblich") {
		ret = Genders().Female()
		ok = true
	} else {
		ret = defaultValue
	}
	return
}

func (o *maritalStates) ParseMaritalStateGerman(name string, defaultValue *MaritalState) (ret *MaritalState, ok bool) {
	if strings.EqualFold(strings.ToLower(name), "ledig") {
		ret = MaritalStates().Single()
		ok = true
	} else if strings.EqualFold(strings.ToLower(name), "verheiratet") {
		ret = MaritalStates().Married()
		ok = true
	} else {
		ret = defaultValue
	}
	return
}

func PersonNameParse(value string) (ret *shared.PersonName, ok bool) {
	str := strings.Trim(value, " ")
	if len(str) > 0 {
		ok = true
		firstLast := strings.Split(str, " ")
		if len(firstLast) >= 2 {
			ret = &shared.PersonName{First: firstLast[0], Last: strings.Join(firstLast[1:], " ")}
		} else if len(firstLast) == 1 {
			ret = &shared.PersonName{First: firstLast[0]}
		} else {
			ret = &shared.PersonName{First: str}
		}
	}
	return
}
