package repository

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/DeveloperGerald/TurtleSoup/model"
	"github.com/google/uuid"
)

func CreateStory(story model.Story) (*model.Story, error) {
	if story.ID == "" {
		newUUID := uuid.New()
		story.ID = newUUID.String()
	}
	err := db.Create(&story).Error
	if err != nil {
		return nil, fmt.Errorf("create story error: %v", err)
	}

	return &story, nil
}

func GetStoryByID(id string) (*model.Story, error) {
	var stories []model.Story
	err := db.Where("id = ?", id).Find(&stories).Error
	if err != nil {
		return nil, fmt.Errorf("get story by id %s error: %v", id, err)
	}

	if len(stories) > 1 {
		return nil, fmt.Errorf("get more than one story by id %s", id)
	}
	if len(stories) < 1 {
		return nil, fmt.Errorf("find no stories by id %s", id)
	}

	story := stories[0]
	return &story, nil
}

func GetRandomStory() (*model.Story, error) {
	var count int
	if err := db.Model(&model.Story{}).Count(&count).Error; err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, fmt.Errorf("no stories found")
	}

	source := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(source)
	offset := rand.Intn(count) // 生成一个 [0, count) 的随机数

	// 3. 获取随机记录
	var story model.Story
	if err := db.Offset(offset).Limit(1).Find(&story).Error; err != nil {
		return nil, err
	}

	return &story, nil
}
