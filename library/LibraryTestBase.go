package library

import (
    "github.com/go-ee/schkola/person"
    "fmt"
    "github.com/go-ee/utils"
    "github.com/google/uuid"
    "time"
)
func NewBooksByPropNames(count int) []*Book {
	items := make([]*Book, count)
	for i := 0; i < count; i++ {
		items[i] = NewBookByPropNames(i)
	}
	return items
}

func NewBookByPropNames(intSalt int) (ret *Book)  {
    ret = NewBook()
    ret.Title = fmt.Sprintf("Title %v", intSalt)
    ret.Description = fmt.Sprintf("Description %v", intSalt)
    ret.Language = fmt.Sprintf("Language %v", intSalt)
    ret.ReleaseDate = utils.PtrTime(time.Now())
    ret.Edition = fmt.Sprintf("Edition %v", intSalt)
    ret.Category = fmt.Sprintf("Category %v", intSalt)
    ret.Author = person.NewPersonNameByPropNames(intSalt)
    ret.Location = NewLocationByPropNames(intSalt)
    ret.Id = uuid.New()
    return
}






func NewLocationsByPropNames(count int) []*Location {
	items := make([]*Location, count)
	for i := 0; i < count; i++ {
		items[i] = NewLocationByPropNames(i)
	}
	return items
}

func NewLocationByPropNames(intSalt int) (ret *Location)  {
    ret = NewLocation()
    ret.Shelf = fmt.Sprintf("Shelf %v", intSalt)
    ret.Fold = fmt.Sprintf("Fold %v", intSalt)
    return
}



