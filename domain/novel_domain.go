package domain

import "github.com/onainadapdap1/golang-crud-redis/model"

type NovelRepo interface {
	CreateNovel(createNovel model.Novel) error
	GetNovelById(id int) (model.Novel, error)
}

type NovelUseCase interface {
	CreateNovel(createNovel model.Novel) error
	GetNovelById(id int) (model.Novel, error)
}