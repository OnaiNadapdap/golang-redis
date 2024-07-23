package repo

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/onainadapdap1/golang-crud-redis/domain"
	"github.com/onainadapdap1/golang-crud-redis/model"
	"gorm.io/gorm"
)

type novelRepo struct {
	db  *gorm.DB
	rdb *redis.Client
}

// GetNovelById implements domain.NovelRepo.
func (n *novelRepo) GetNovelById(id int) (model.Novel, error) {
	var novels model.Novel
	var ctx = context.Background()
	
	// first check data is available in redis
	result, err := n.rdb.Get(ctx, "novel"+strconv.Itoa(id)).Result()
	if err != nil {
		return novels, err
	}
	// if data is available in redis, decode it from json, and return it
	if len(result) > 0 {
		err := json.Unmarshal([]byte(result), &novels)
		return novels, err
	}

	// if data was not available in redis, got it from database
	err = n.db.Model(model.Novel{}).Select("id", "name", "author", "description").Where("id=?", id).Find(&novels).Error
	if err != nil {
		return model.Novel{}, err
	}
	// encode slice into json before saving into redis
	jsonBytes, err := json.Marshal(novels)
	if err != nil {
		return novels, err
	}
	jsonString := string(jsonBytes)

	// set json-encode value in redis
	err = n.rdb.Set(ctx, "novel"+strconv.Itoa(id), jsonString, 24*time.Hour).Err()
	if err != nil {
		return novels, err
	}

	return novels, nil
}

func (n *novelRepo) CreateNovel(createNovel model.Novel) error {
	if err := n.db.Create(&createNovel).Error; err != nil {
		return errors.New("internal server error: cannot create novel")
	}

	return nil
}

func NewNovelRepo(db *gorm.DB, rdb *redis.Client) domain.NovelRepo {
	return &novelRepo{
		db:  db,
		rdb: rdb,
	}
}
