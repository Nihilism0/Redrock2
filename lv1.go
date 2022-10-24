package main

import (
	"fmt"
)

//提问题

type Question struct {
	Id      int64  //用户名
	title   string //问题标题
	content string //内容
}

//结构体制造机器

func NewQuestion(id int64, title string, content string) Question {
	return Question{
		Id:      id,
		title:   title,
		content: content,
	}
}

//输出方法

func (q Question) Output() {
	fmt.Printf("ID为%v的GO语言学习者提出了一个\"%v\"的问题,内容是%v\n", q.Id, q.title, q.content)
}

func main() {
	Question1 := NewQuestion(123456, "How to learn go", "I really dont know!")
	Question2 := NewQuestion(654321, "What is interface", "I am kidding")
	Question1.Output()
	Question2.Output()
}
