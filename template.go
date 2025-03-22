package main

import (
	"context"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
)

// createTemplate 创建一个消息模板
func createTemplate() prompt.ChatTemplate {
	return prompt.FromMessages(
		schema.FString,
		schema.SystemMessage("你是一个{role}。你需要用{style}的语气回答问题。你的目标是{target}"),
		schema.MessagesPlaceholder("chat_history", true),
		schema.UserMessage("问题：{question}"),
	)
}

// createMessage 基于指定的消息模板实例化消息
func createMessage(ctx context.Context, template prompt.ChatTemplate, chatHistory []*schema.Message, question string) ([]*schema.Message, error) {
	messages, err := template.Format(ctx, map[string]interface{}{
		"role":         "制造业设备管理专家",
		"style":        "专业、严谨、自信",
		"target":       "回答制造业设备管理相关的问题并给出专业可行的解决方案",
		"chat_history": chatHistory,
		"question":     question,
	})
	if err != nil {
		return nil, err
	}
	return messages, nil
}
