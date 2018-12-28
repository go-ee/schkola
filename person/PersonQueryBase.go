package person

import (
    "context"
    "github.com/eugeis/gee/eh"
    "github.com/looplab/eventhorizon"
)
type ChurchQueryRepository struct {
    repo eventhorizon.ReadRepo `json:"repo" eh:"optional"`
    context context.Context `json:"context" eh:"optional"`
}

func NewChurchQueryRepository(repo eventhorizon.ReadRepo, context context.Context) (ret *ChurchQueryRepository) {
    ret = &ChurchQueryRepository{
        repo: repo,
        context: context,
    }
    return
}

func (o *ChurchQueryRepository) FindAll() (ret []*Church, err error) {
    var result []eventhorizon.Entity
	if result, err = o.repo.FindAll(o.context); err == nil {
        ret = make([]*Church, len(result))
		for i, e := range result {
            ret[i] = e.(*Church)
		}
    }
    return
}

func (o *ChurchQueryRepository) FindById(id eventhorizon.UUID) (ret *Church, err error) {
    var result eventhorizon.Entity
	if result, err = o.repo.Find(o.context, id); err == nil {
        ret = result.(*Church)
    }
    return
}

func (o *ChurchQueryRepository) CountAll() (ret int, err error) {
    var result []*Church
	if result, err = o.FindAll(); err == nil {
        ret = len(result)
    }
    return
}

func (o *ChurchQueryRepository) CountById(id eventhorizon.UUID) (ret int, err error) {
    var result *Church
	if result, err = o.FindById(id); err == nil && result != nil {
        ret = 1
    }
    return
}

func (o *ChurchQueryRepository) ExistAll() (ret bool, err error) {
    var result int
	if result, err = o.CountAll(); err == nil {
        ret = result > 0
    }
    return
}

func (o *ChurchQueryRepository) ExistById(id eventhorizon.UUID) (ret bool, err error) {
    var result int
	if result, err = o.CountById(id); err == nil {
        ret = result > 0
    }
    return
}


type GraduationQueryRepository struct {
    repo eventhorizon.ReadRepo `json:"repo" eh:"optional"`
    context context.Context `json:"context" eh:"optional"`
}

func NewGraduationQueryRepository(repo eventhorizon.ReadRepo, context context.Context) (ret *GraduationQueryRepository) {
    ret = &GraduationQueryRepository{
        repo: repo,
        context: context,
    }
    return
}

func (o *GraduationQueryRepository) FindAll() (ret []*Graduation, err error) {
    var result []eventhorizon.Entity
	if result, err = o.repo.FindAll(o.context); err == nil {
        ret = make([]*Graduation, len(result))
		for i, e := range result {
            ret[i] = e.(*Graduation)
		}
    }
    return
}

func (o *GraduationQueryRepository) FindById(id eventhorizon.UUID) (ret *Graduation, err error) {
    var result eventhorizon.Entity
	if result, err = o.repo.Find(o.context, id); err == nil {
        ret = result.(*Graduation)
    }
    return
}

func (o *GraduationQueryRepository) CountAll() (ret int, err error) {
    var result []*Graduation
	if result, err = o.FindAll(); err == nil {
        ret = len(result)
    }
    return
}

func (o *GraduationQueryRepository) CountById(id eventhorizon.UUID) (ret int, err error) {
    var result *Graduation
	if result, err = o.FindById(id); err == nil && result != nil {
        ret = 1
    }
    return
}

func (o *GraduationQueryRepository) ExistAll() (ret bool, err error) {
    var result int
	if result, err = o.CountAll(); err == nil {
        ret = result > 0
    }
    return
}

func (o *GraduationQueryRepository) ExistById(id eventhorizon.UUID) (ret bool, err error) {
    var result int
	if result, err = o.CountById(id); err == nil {
        ret = result > 0
    }
    return
}


type ProfileQueryRepository struct {
    repo eventhorizon.ReadRepo `json:"repo" eh:"optional"`
    context context.Context `json:"context" eh:"optional"`
}

func NewProfileQueryRepository(repo eventhorizon.ReadRepo, context context.Context) (ret *ProfileQueryRepository) {
    ret = &ProfileQueryRepository{
        repo: repo,
        context: context,
    }
    return
}

func (o *ProfileQueryRepository) FindByEmail(email string) (ret *Profile, err error) {
    err = eh.QueryNotImplemented("findProfileByEmail")
    return
}

func (o *ProfileQueryRepository) FindByPhone(phone string) (ret *Profile, err error) {
    err = eh.QueryNotImplemented("findProfileByPhone")
    return
}

func (o *ProfileQueryRepository) FindAll() (ret []*Profile, err error) {
    var result []eventhorizon.Entity
	if result, err = o.repo.FindAll(o.context); err == nil {
        ret = make([]*Profile, len(result))
		for i, e := range result {
            ret[i] = e.(*Profile)
		}
    }
    return
}

func (o *ProfileQueryRepository) FindById(id eventhorizon.UUID) (ret *Profile, err error) {
    var result eventhorizon.Entity
	if result, err = o.repo.Find(o.context, id); err == nil {
        ret = result.(*Profile)
    }
    return
}

func (o *ProfileQueryRepository) CountAll() (ret int, err error) {
    var result []*Profile
	if result, err = o.FindAll(); err == nil {
        ret = len(result)
    }
    return
}

func (o *ProfileQueryRepository) CountById(id eventhorizon.UUID) (ret int, err error) {
    var result *Profile
	if result, err = o.FindById(id); err == nil && result != nil {
        ret = 1
    }
    return
}

func (o *ProfileQueryRepository) ExistAll() (ret bool, err error) {
    var result int
	if result, err = o.CountAll(); err == nil {
        ret = result > 0
    }
    return
}

func (o *ProfileQueryRepository) ExistById(id eventhorizon.UUID) (ret bool, err error) {
    var result int
	if result, err = o.CountById(id); err == nil {
        ret = result > 0
    }
    return
}









