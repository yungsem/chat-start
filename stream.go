package main

import (
	"fmt"
	"github.com/cloudwego/eino/schema"
	"io"
	"log/slog"
)

// reportStream 流式输出 ChatModel 回复的消息，并返回完整的消息内容
func reportStream(stream *schema.StreamReader[*schema.Message]) string {
	defer stream.Close()

	// 完整的消息内容
	var fullMessage string

	for {
		// 读取消息
		message, err := stream.Recv()

		// 本轮消息读完
		if err == io.EOF {
			fmt.Println()
			return fullMessage
		}

		// 其他错误
		if err != nil {
			slog.Error("读取 ChatModel 回复的消息发生错误：", err)
			return fullMessage
		}

		// 输出消息
		fmt.Print(message.Content)

		// 分次读取到的消息，都累加到 fullMessage
		fullMessage = fullMessage + message.Content
	}
}
