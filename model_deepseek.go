package main

import (
	"context"
	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/components/model"
)

// createChatModel 创建一个 ChatModel 实例
func createChatModel(ctx context.Context) (model.ChatModel, error) {
	chatModel, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
		Model:   "deepseek-chat",
		BaseURL: "https://api.deepseek.com",
		APIKey:  "api-key",
	})

	if err != nil {
		return nil, err
	}
	return chatModel, nil
}
