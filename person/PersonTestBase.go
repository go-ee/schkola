package person

import (
    "fmt"
    "github.com/go-ee/utils"
    "github.com/google/uuid"
    "time"
)
func NewChurchesByPropNames(count int) []*Church {
	items := make([]*Church, count)
	for i := 0; i < count; i++ {
		items[i] = NewChurchByPropNames(i)
	}
	return items
}

func NewChurchByPropNames(intSalt int) (ret *Church)  {
    ret = NewChurch()
    ret.Name = fmt.Sprintf("Name %v", intSalt)
    ret.Address = NewAddressByPropNames(intSalt)
    ret.Pastor = NewPersonNameByPropNames(intSalt)
    ret.Contact = NewContactByPropNames(intSalt)
    ret.Association = fmt.Sprintf("Association %v", intSalt)
    ret.Id = uuid.New()
    return
}


func NewGraduationsByPropNames(count int) []*Graduation {
	items := make([]*Graduation, count)
	for i := 0; i < count; i++ {
		items[i] = NewGraduationByPropNames(i)
	}
	return items
}

func NewGraduationByPropNames(intSalt int) (ret *Graduation)  {
    ret = NewGraduation()
    ret.Name = fmt.Sprintf("Name %v", intSalt)
    ret.Level = GraduationLevels().Unknown()
    ret.Id = uuid.New()
    return
}


func NewProfilesByPropNames(count int) []*Profile {
	items := make([]*Profile, count)
	for i := 0; i < count; i++ {
		items[i] = NewProfileByPropNames(i)
	}
	return items
}

func NewProfileByPropNames(intSalt int) (ret *Profile)  {
    ret = NewProfile()
    ret.Gender = Genders().Unknown()
    ret.Name = NewPersonNameByPropNames(intSalt)
    ret.BirthName = fmt.Sprintf("BirthName %v", intSalt)
    ret.Birthday = utils.PtrTime(time.Now())
    ret.Address = NewAddressByPropNames(intSalt)
    ret.Contact = NewContactByPropNames(intSalt)
    ret.PhotoData = []byte(fmt.Sprintf("PhotoData %v", intSalt))
    ret.Photo = fmt.Sprintf("Photo %v", intSalt)
    ret.Family = NewFamilyByPropNames(intSalt)
    ret.Church = NewChurchInfoByPropNames(intSalt)
    ret.Education = NewEducationByPropNames(intSalt)
    ret.Id = uuid.New()
    return
}






func NewAddresssByPropNames(count int) []*Address {
	items := make([]*Address, count)
	for i := 0; i < count; i++ {
		items[i] = NewAddressByPropNames(i)
	}
	return items
}

func NewAddressByPropNames(intSalt int) (ret *Address)  {
    ret = NewAddress()
    ret.Street = fmt.Sprintf("Street %v", intSalt)
    ret.Suite = fmt.Sprintf("Suite %v", intSalt)
    ret.City = fmt.Sprintf("City %v", intSalt)
    ret.Code = fmt.Sprintf("Code %v", intSalt)
    ret.Country = fmt.Sprintf("Country %v", intSalt)
    return
}


func NewChurchInfosByPropNames(count int) []*ChurchInfo {
	items := make([]*ChurchInfo, count)
	for i := 0; i < count; i++ {
		items[i] = NewChurchInfoByPropNames(i)
	}
	return items
}

func NewChurchInfoByPropNames(intSalt int) (ret *ChurchInfo)  {
    ret = NewChurchInfo()
    ret.Church = fmt.Sprintf("Church %v", intSalt)
    ret.Member = false
    ret.Services = fmt.Sprintf("Services %v", intSalt)
    return
}


func NewContactsByPropNames(count int) []*Contact {
	items := make([]*Contact, count)
	for i := 0; i < count; i++ {
		items[i] = NewContactByPropNames(i)
	}
	return items
}

func NewContactByPropNames(intSalt int) (ret *Contact)  {
    ret = NewContact()
    ret.Phone = fmt.Sprintf("Phone %v", intSalt)
    ret.Email = fmt.Sprintf("Email %v", intSalt)
    ret.Cellphone = fmt.Sprintf("Cellphone %v", intSalt)
    return
}


func NewEducationsByPropNames(count int) []*Education {
	items := make([]*Education, count)
	for i := 0; i < count; i++ {
		items[i] = NewEducationByPropNames(i)
	}
	return items
}

func NewEducationByPropNames(intSalt int) (ret *Education)  {
    ret = NewEducation()
    ret.Graduation = NewGraduationByPropNames(intSalt)
    ret.Other = fmt.Sprintf("Other %v", intSalt)
    ret.Profession = fmt.Sprintf("Profession %v", intSalt)
    return
}


func NewFamilysByPropNames(count int) []*Family {
	items := make([]*Family, count)
	for i := 0; i < count; i++ {
		items[i] = NewFamilyByPropNames(i)
	}
	return items
}

func NewFamilyByPropNames(intSalt int) (ret *Family)  {
    ret = NewFamily()
    ret.MaritalState = MaritalStates().Unknown()
    ret.ChildrenCount = intSalt
    ret.Partner = NewPersonNameByPropNames(intSalt)
    return
}


func NewPersonNamesByPropNames(count int) []*PersonName {
	items := make([]*PersonName, count)
	for i := 0; i < count; i++ {
		items[i] = NewPersonNameByPropNames(i)
	}
	return items
}

func NewPersonNameByPropNames(intSalt int) (ret *PersonName)  {
    ret = NewPersonName()
    ret.First = fmt.Sprintf("First %v", intSalt)
    ret.Last = fmt.Sprintf("Last %v", intSalt)
    return
}



