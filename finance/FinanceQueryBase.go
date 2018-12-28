package finance

import (
    "context"
    "github.com/looplab/eventhorizon"
)
type ExpenseQueryRepository struct {
    repo eventhorizon.ReadRepo `json:"repo" eh:"optional"`
    context context.Context `json:"context" eh:"optional"`
}

func NewExpenseQueryRepository(repo eventhorizon.ReadRepo, context context.Context) (ret *ExpenseQueryRepository) {
    ret = &ExpenseQueryRepository{
        repo: repo,
        context: context,
    }
    return
}

func (o *ExpenseQueryRepository) FindAll() (ret []*Expense, err error) {
    var result []eventhorizon.Entity
	if result, err = o.repo.FindAll(o.context); err == nil {
        ret = make([]*Expense, len(result))
		for i, e := range result {
            ret[i] = e.(*Expense)
		}
    }
    return
}

func (o *ExpenseQueryRepository) FindById(id eventhorizon.UUID) (ret *Expense, err error) {
    var result eventhorizon.Entity
	if result, err = o.repo.Find(o.context, id); err == nil {
        ret = result.(*Expense)
    }
    return
}

func (o *ExpenseQueryRepository) CountAll() (ret int, err error) {
    var result []*Expense
	if result, err = o.FindAll(); err == nil {
        ret = len(result)
    }
    return
}

func (o *ExpenseQueryRepository) CountById(id eventhorizon.UUID) (ret int, err error) {
    var result *Expense
	if result, err = o.FindById(id); err == nil && result != nil {
        ret = 1
    }
    return
}

func (o *ExpenseQueryRepository) ExistAll() (ret bool, err error) {
    var result int
	if result, err = o.CountAll(); err == nil {
        ret = result > 0
    }
    return
}

func (o *ExpenseQueryRepository) ExistById(id eventhorizon.UUID) (ret bool, err error) {
    var result int
	if result, err = o.CountById(id); err == nil {
        ret = result > 0
    }
    return
}


type ExpensePurposeQueryRepository struct {
    repo eventhorizon.ReadRepo `json:"repo" eh:"optional"`
    context context.Context `json:"context" eh:"optional"`
}

func NewExpensePurposeQueryRepository(repo eventhorizon.ReadRepo, context context.Context) (ret *ExpensePurposeQueryRepository) {
    ret = &ExpensePurposeQueryRepository{
        repo: repo,
        context: context,
    }
    return
}

func (o *ExpensePurposeQueryRepository) FindAll() (ret []*ExpensePurpose, err error) {
    var result []eventhorizon.Entity
	if result, err = o.repo.FindAll(o.context); err == nil {
        ret = make([]*ExpensePurpose, len(result))
		for i, e := range result {
            ret[i] = e.(*ExpensePurpose)
		}
    }
    return
}

func (o *ExpensePurposeQueryRepository) FindById(id eventhorizon.UUID) (ret *ExpensePurpose, err error) {
    var result eventhorizon.Entity
	if result, err = o.repo.Find(o.context, id); err == nil {
        ret = result.(*ExpensePurpose)
    }
    return
}

func (o *ExpensePurposeQueryRepository) CountAll() (ret int, err error) {
    var result []*ExpensePurpose
	if result, err = o.FindAll(); err == nil {
        ret = len(result)
    }
    return
}

func (o *ExpensePurposeQueryRepository) CountById(id eventhorizon.UUID) (ret int, err error) {
    var result *ExpensePurpose
	if result, err = o.FindById(id); err == nil && result != nil {
        ret = 1
    }
    return
}

func (o *ExpensePurposeQueryRepository) ExistAll() (ret bool, err error) {
    var result int
	if result, err = o.CountAll(); err == nil {
        ret = result > 0
    }
    return
}

func (o *ExpensePurposeQueryRepository) ExistById(id eventhorizon.UUID) (ret bool, err error) {
    var result int
	if result, err = o.CountById(id); err == nil {
        ret = result > 0
    }
    return
}


type FeeQueryRepository struct {
    repo eventhorizon.ReadRepo `json:"repo" eh:"optional"`
    context context.Context `json:"context" eh:"optional"`
}

func NewFeeQueryRepository(repo eventhorizon.ReadRepo, context context.Context) (ret *FeeQueryRepository) {
    ret = &FeeQueryRepository{
        repo: repo,
        context: context,
    }
    return
}

func (o *FeeQueryRepository) FindAll() (ret []*Fee, err error) {
    var result []eventhorizon.Entity
	if result, err = o.repo.FindAll(o.context); err == nil {
        ret = make([]*Fee, len(result))
		for i, e := range result {
            ret[i] = e.(*Fee)
		}
    }
    return
}

func (o *FeeQueryRepository) FindById(id eventhorizon.UUID) (ret *Fee, err error) {
    var result eventhorizon.Entity
	if result, err = o.repo.Find(o.context, id); err == nil {
        ret = result.(*Fee)
    }
    return
}

func (o *FeeQueryRepository) CountAll() (ret int, err error) {
    var result []*Fee
	if result, err = o.FindAll(); err == nil {
        ret = len(result)
    }
    return
}

func (o *FeeQueryRepository) CountById(id eventhorizon.UUID) (ret int, err error) {
    var result *Fee
	if result, err = o.FindById(id); err == nil && result != nil {
        ret = 1
    }
    return
}

func (o *FeeQueryRepository) ExistAll() (ret bool, err error) {
    var result int
	if result, err = o.CountAll(); err == nil {
        ret = result > 0
    }
    return
}

func (o *FeeQueryRepository) ExistById(id eventhorizon.UUID) (ret bool, err error) {
    var result int
	if result, err = o.CountById(id); err == nil {
        ret = result > 0
    }
    return
}


type FeeKindQueryRepository struct {
    repo eventhorizon.ReadRepo `json:"repo" eh:"optional"`
    context context.Context `json:"context" eh:"optional"`
}

func NewFeeKindQueryRepository(repo eventhorizon.ReadRepo, context context.Context) (ret *FeeKindQueryRepository) {
    ret = &FeeKindQueryRepository{
        repo: repo,
        context: context,
    }
    return
}

func (o *FeeKindQueryRepository) FindAll() (ret []*FeeKind, err error) {
    var result []eventhorizon.Entity
	if result, err = o.repo.FindAll(o.context); err == nil {
        ret = make([]*FeeKind, len(result))
		for i, e := range result {
            ret[i] = e.(*FeeKind)
		}
    }
    return
}

func (o *FeeKindQueryRepository) FindById(id eventhorizon.UUID) (ret *FeeKind, err error) {
    var result eventhorizon.Entity
	if result, err = o.repo.Find(o.context, id); err == nil {
        ret = result.(*FeeKind)
    }
    return
}

func (o *FeeKindQueryRepository) CountAll() (ret int, err error) {
    var result []*FeeKind
	if result, err = o.FindAll(); err == nil {
        ret = len(result)
    }
    return
}

func (o *FeeKindQueryRepository) CountById(id eventhorizon.UUID) (ret int, err error) {
    var result *FeeKind
	if result, err = o.FindById(id); err == nil && result != nil {
        ret = 1
    }
    return
}

func (o *FeeKindQueryRepository) ExistAll() (ret bool, err error) {
    var result int
	if result, err = o.CountAll(); err == nil {
        ret = result > 0
    }
    return
}

func (o *FeeKindQueryRepository) ExistById(id eventhorizon.UUID) (ret bool, err error) {
    var result int
	if result, err = o.CountById(id); err == nil {
        ret = result > 0
    }
    return
}









