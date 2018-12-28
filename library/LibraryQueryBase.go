package library

import (
    "context"
    "github.com/eugeis/gee/eh"
    "github.com/looplab/eventhorizon"
)
type BookQueryRepository struct {
    repo eventhorizon.ReadRepo `json:"repo" eh:"optional"`
    context context.Context `json:"context" eh:"optional"`
}

func NewBookQueryRepository(repo eventhorizon.ReadRepo, context context.Context) (ret *BookQueryRepository) {
    ret = &BookQueryRepository{
        repo: repo,
        context: context,
    }
    return
}

func (o *BookQueryRepository) FindByTitle(title string) (ret *Book, err error) {
    err = eh.QueryNotImplemented("findBookByTitle")
    return
}

func (o *BookQueryRepository) FindByPattern(pattern string) (ret *Book, err error) {
    err = eh.QueryNotImplemented("findBookByPattern")
    return
}

func (o *BookQueryRepository) FindAll() (ret []*Book, err error) {
    var result []eventhorizon.Entity
	if result, err = o.repo.FindAll(o.context); err == nil {
        ret = make([]*Book, len(result))
		for i, e := range result {
            ret[i] = e.(*Book)
		}
    }
    return
}

func (o *BookQueryRepository) FindById(id eventhorizon.UUID) (ret *Book, err error) {
    var result eventhorizon.Entity
	if result, err = o.repo.Find(o.context, id); err == nil {
        ret = result.(*Book)
    }
    return
}

func (o *BookQueryRepository) CountAll() (ret int, err error) {
    var result []*Book
	if result, err = o.FindAll(); err == nil {
        ret = len(result)
    }
    return
}

func (o *BookQueryRepository) CountById(id eventhorizon.UUID) (ret int, err error) {
    var result *Book
	if result, err = o.FindById(id); err == nil && result != nil {
        ret = 1
    }
    return
}

func (o *BookQueryRepository) ExistAll() (ret bool, err error) {
    var result int
	if result, err = o.CountAll(); err == nil {
        ret = result > 0
    }
    return
}

func (o *BookQueryRepository) ExistById(id eventhorizon.UUID) (ret bool, err error) {
    var result int
	if result, err = o.CountById(id); err == nil {
        ret = result > 0
    }
    return
}









