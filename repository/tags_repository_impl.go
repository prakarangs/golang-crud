package repository

import (
	"golang-crud-gin/model"

	"gorm.io/gorm"
)

// TBC: https://www.youtube.com/watch?v=cLEXgca3FM8&ab_channel=lemoncode21 at 8:30
type TagsRepositoryImpl struct {
	Db *gorm.DB
}

func NewTagsREpositoryImpl(Db *gorm.DB) TagsRepository {
	return &TagsRepositoryImpl{Db: Db}
}

// Delete implements TagsRepository
func (t *TagsRepositoryImpl) Delete(tagsId int) {
	panic("unimplemented")
}

// FindAll implements TagsRepository
func (*TagsRepositoryImpl) FindAll() []model.Tags {
	panic("unimplemented")
}

// FindById implements TagsRepository
func (*TagsRepositoryImpl) FindById(tagsId int) (tags model.Tags, err error) {
	panic("unimplemented")
}

// Save implements TagsRepository
func (*TagsRepositoryImpl) Save(tags model.Tags) {
	panic("unimplemented")
}

// Update implements TagsRepository
func (*TagsRepositoryImpl) Update(tags model.Tags) {
	panic("unimplemented")
}
