package usecase

import (
	"errors"

	"github.com/onainadapdap1/golang-crud-redis/domain"
	"github.com/onainadapdap1/golang-crud-redis/model"
)

type novelUseCase struct {
	novelRepo domain.NovelRepo
}

func NewNovelUseCase(novelRepo domain.NovelRepo) domain.NovelUseCase {
	return &novelUseCase{
		novelRepo: novelRepo,
	}
}

// CreateNovel implements domain.NovelUseCase.
func (n *novelUseCase) CreateNovel(createNovel model.Novel) error {
	err := n.novelRepo.CreateNovel(createNovel)
	return err
}

// GetNovelById implements domain.NovelUseCase.
func (n *novelUseCase) GetNovelById(id int) (model.Novel, error) {
	novel, err := n.novelRepo.GetNovelById(id)
	if err != nil {
		return model.Novel{}, errors.New("internal server error" + err.Error())
	}
	return novel, nil
}
