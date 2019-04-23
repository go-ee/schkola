package finance

import (
    "github.com/go-ee/schkola/person"
    "fmt"
    "github.com/go-ee/utils"
    "github.com/google/uuid"
    "time"
)
func NewExpensesByPropNames(count int) []*Expense {
	items := make([]*Expense, count)
	for i := 0; i < count; i++ {
		items[i] = NewExpenseByPropNames(i)
	}
	return items
}

func NewExpenseByPropNames(intSalt int) (ret *Expense)  {
    ret = NewExpense()
    ret.Purpose = NewExpensePurposeByPropNames(intSalt)
    ret.Amount = float64(intSalt)
    ret.Profile = person.NewProfileByPropNames(intSalt)
    ret.Date = utils.PtrTime(time.Now())
    ret.Id = uuid.New()
    return
}


func NewExpensePurposesByPropNames(count int) []*ExpensePurpose {
	items := make([]*ExpensePurpose, count)
	for i := 0; i < count; i++ {
		items[i] = NewExpensePurposeByPropNames(i)
	}
	return items
}

func NewExpensePurposeByPropNames(intSalt int) (ret *ExpensePurpose)  {
    ret = NewExpensePurpose()
    ret.Name = fmt.Sprintf("Name %v", intSalt)
    ret.Description = fmt.Sprintf("Description %v", intSalt)
    ret.Id = uuid.New()
    return
}


func NewFeesByPropNames(count int) []*Fee {
	items := make([]*Fee, count)
	for i := 0; i < count; i++ {
		items[i] = NewFeeByPropNames(i)
	}
	return items
}

func NewFeeByPropNames(intSalt int) (ret *Fee)  {
    ret = NewFee()
    ret.Student = person.NewProfileByPropNames(intSalt)
    ret.Amount = float64(intSalt)
    ret.Kind = NewFeeKindByPropNames(intSalt)
    ret.Date = utils.PtrTime(time.Now())
    ret.Id = uuid.New()
    return
}


func NewFeeKindsByPropNames(count int) []*FeeKind {
	items := make([]*FeeKind, count)
	for i := 0; i < count; i++ {
		items[i] = NewFeeKindByPropNames(i)
	}
	return items
}

func NewFeeKindByPropNames(intSalt int) (ret *FeeKind)  {
    ret = NewFeeKind()
    ret.Name = fmt.Sprintf("Name %v", intSalt)
    ret.Amount = float64(intSalt)
    ret.Description = fmt.Sprintf("Description %v", intSalt)
    ret.Id = uuid.New()
    return
}







