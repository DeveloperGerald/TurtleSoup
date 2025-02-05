package repository

import (
	"fmt"

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
