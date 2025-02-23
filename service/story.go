package service

import (
	"github.com/DeveloperGerald/TurtleSoup/model"
	"github.com/DeveloperGerald/TurtleSoup/repository"
)

func CreateStory(title, riddle, fullStory, owner string) (*model.Story, error) {
	story := model.Story{
		Title:     title,
		Riddle:    riddle,
		FullStory: fullStory,
		Owner:     owner,
	}
	created, err := repository.CreateStory(story)
	if err != nil {
		return nil, err
	}

	return created, nil
}
