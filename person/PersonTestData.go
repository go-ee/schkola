package person

import (
	"fmt"

	"github.com/google/uuid"
)

func BuildChurches(count int) []*Church {
	items := make([]*Church, count)
	for i := 0; i < count; i++ {
		id := uuid.New()
		items[i] = &Church{
			Name: fmt.Sprintf("Name %v", i),
			//Address *Address `json:"address" eh:"optional"`
			//Pastor *PersonName `json:"pastor" eh:"optional"`
			//Contact *Contact `json:"contact" eh:"optional"`
			Association: fmt.Sprintf("Association %v", i),
			Id:          id,
		}
	}
	return items
}
