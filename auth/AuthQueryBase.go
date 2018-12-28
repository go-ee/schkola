package auth

import (
    "context"
    "github.com/looplab/eventhorizon"
)
type AccountQueryRepository struct {
    repo eventhorizon.ReadRepo `json:"repo" eh:"optional"`
    context context.Context `json:"context" eh:"optional"`
}

func NewAccountQueryRepository(repo eventhorizon.ReadRepo, context context.Context) (ret *AccountQueryRepository) {
    ret = &AccountQueryRepository{
        repo: repo,
        context: context,
    }
    return
}

func (o *AccountQueryRepository) FindAll() (ret []*Account, err error) {
    var result []eventhorizon.Entity
	if result, err = o.repo.FindAll(o.context); err == nil {
        ret = make([]*Account, len(result))
		for i, e := range result {
            ret[i] = e.(*Account)
		}
    }
    return
}

func (o *AccountQueryRepository) FindById(id eventhorizon.UUID) (ret *Account, err error) {
    var result eventhorizon.Entity
	if result, err = o.repo.Find(o.context, id); err == nil {
        ret = result.(*Account)
    }
    return
}

func (o *AccountQueryRepository) CountAll() (ret int, err error) {
    var result []*Account
	if result, err = o.FindAll(); err == nil {
        ret = len(result)
    }
    return
}

func (o *AccountQueryRepository) CountById(id eventhorizon.UUID) (ret int, err error) {
    var result *Account
	if result, err = o.FindById(id); err == nil && result != nil {
        ret = 1
    }
    return
}

func (o *AccountQueryRepository) ExistAll() (ret bool, err error) {
    var result int
	if result, err = o.CountAll(); err == nil {
        ret = result > 0
    }
    return
}

func (o *AccountQueryRepository) ExistById(id eventhorizon.UUID) (ret bool, err error) {
    var result int
	if result, err = o.CountById(id); err == nil {
        ret = result > 0
    }
    return
}









