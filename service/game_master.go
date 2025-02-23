package service

import (
	"fmt"
	"log"

	"github.com/DeveloperGerald/TurtleSoup/model"
	deepseek "github.com/DeveloperGerald/TurtleSoup/pkg/deep_seek"
	"github.com/DeveloperGerald/TurtleSoup/repository"
)

const GameMasterSystemContent = "你是一名专业的海龟汤主持人，基于海龟汤的题目与汤底，你需要与用户进行互动。用户回答符合汤底则回复‘是’，用户回答不符合汤底回复‘否’，用户回答与题目无关则回复‘与题目无关’，用户猜中汤底(需要相对完整的猜中整个故事)则回复‘bingo!’。其余回答一律回复‘请围绕题目作答’，如果用户的问题无法用是、否、与题目无关来回复，则回复‘请围绕题目作答’"

func GiveAnswer(storyID string, userAnswer string) (string, error) {
	// get story
	story, err := repository.GetStoryByID(storyID)
	if err != nil {
		return "", err
	}

	messages := make([]deepseek.Message, 0)
	messages = append(messages, deepseek.Message{Role: deepseek.RoleSystem, Content: GameMasterSystemContent})
	messages = append(messages, deepseek.Message{Role: deepseek.RoleAssistant, Content: fmt.Sprintf("题目：%s", story.Riddle)})
	messages = append(messages, deepseek.Message{Role: deepseek.RoleAssistant, Content: fmt.Sprintf("汤底：%s", story.FullStory)})
	messages = append(messages, deepseek.Message{Role: deepseek.RoleUser, Content: userAnswer})

	// ask deepseek
	chatCompletion, err := deepseek.GetClient().CompleteChat(deepseek.CompoleteChatInput{
		Model:          deepseek.ModelDeepseekChat,
		MaxTokens:      50,
		ResponseFormat: deepseek.ResponseFormat{Type: deepseek.ResponseFormatTypeText},
		Stream:         false,
		Temperature:    1,
		Messages:       messages,
	})
	if err != nil {
		return "", err
	}

	if len(chatCompletion.Choices) < 1 {
		return "", fmt.Errorf("no choices in chat completion")
	}
	if len(chatCompletion.Choices) > 1 {
		return "", fmt.Errorf("more than one choices in chat completion")
	}

	answer := chatCompletion.Choices[0].Message.Content
	switch answer {
	case "是":
		break
	case "否":
		break
	case "与题目无关":
		break
	case "bingo!":
		break
	case "请围绕题目作答":
		break
	default:
		log.Printf("story: %s, user answer: %s, unexpected answer: %s", storyID, userAnswer, answer)
		answer = "系统异常，请调整您的描述"
	}

	return answer, nil
}

func GetRandomStory() (*model.Story, error) {
	return repository.GetRandomStory()
}
