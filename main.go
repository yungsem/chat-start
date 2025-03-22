package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/eino/schema"
	"log/slog"
)

// 我是一个注塑行业的设备工程师，有一个问题困扰我很久了，就是我们在对设备进行工艺参数点检的时候，效率非常低下，平均每台设备每次点检需要手动录入40多个工艺参数的时，非常耗时，也容易录错，请帮我解决此问题

func main() {
	// ctx
	ctx := context.Background()

	// 创建 ChatTemplate
	template := createTemplate()

	// 创建 ChatModel
	model, err := createChatModel(ctx)
	if err != nil {
		slog.Error("创建 ChatModel 发生错误：", err)
		return
	}

	// 定义 ChatHistory
	var chatHistory []*schema.Message

	// 实现对话
	for {
		fmt.Println("请输入你的问题：")

		// 读取用户输入
		var question string
		n, err := fmt.Scanln(&question)
		if err != nil {
			slog.Error("读取用户输入发生错误：", err)
			return
		}
		if n <= 0 {
			slog.Error("用户没有输入任何内容")
			return
		}

		// 创建 Message
		message, err := createMessage(ctx, template, chatHistory, question)
		if err != nil {
			slog.Error("创建 Message 发生错误：", err)
			return
		}

		// Message 提交到 ChatModel
		stream, err := model.Stream(ctx, message)
		if err != nil {
			slog.Error("ChatModel 处理 Message 发生错误：", err)
			return
		}

		fullMessage := reportStream(stream)
		assistantMessage := schema.AssistantMessage(fullMessage, nil)
		chatHistory = append(chatHistory, assistantMessage)
	}
}
