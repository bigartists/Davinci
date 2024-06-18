package qa

import "context"

type Service interface {
	// GetQuestions 获取问题列表，question简化结构
	GetQuestions(ctx context.Context, pager *Pager) ([]*Question, error)
	// GetQuestion 获取某个问题详情，question简化结构
	GetQuestion(ctx context.Context, questionID int64) (*Question, error)
	// PostQuestion 上传某个问题
	// ctx必须带操作人id
	PostQuestion(ctx context.Context, question *Question) error

	// QuestionLoadAuthor 问题加载Author字段
	QuestionLoadAuthor(ctx context.Context, question *Question) error
	// QuestionsLoadAuthor 批量加载Author字段
	QuestionsLoadAuthor(ctx context.Context, questions *[]*Question) error

	// QuestionLoadAnswers 单个问题加载Answers
	QuestionLoadAnswers(ctx context.Context, question *Question) error
	// QuestionsLoadAnswers 批量问题加载Answers
	QuestionsLoadAnswers(ctx context.Context, questions *[]*Question) error

	// PostAnswer 上传某个回答
	// ctx必须带操作人信息
	PostAnswer(ctx context.Context, answer *Answer) error
	// GetAnswer 获取回答
	GetAnswer(ctx context.Context, answerID int64) (*Answer, error)

	// AnswerLoadAuthor 问题加载Author字段
	AnswerLoadAuthor(ctx context.Context, question *Answer) error
	// AnswersLoadAuthor 批量加载Author字段
	AnswersLoadAuthor(ctx context.Context, questions *[]*Answer) error

	// DeleteQuestion 删除问题，同时删除对应的回答
	// ctx必须带操作人信息
	DeleteQuestion(ctx context.Context, questionID int64) error
	// DeleteAnswer 删除某个回答
	// ctx必须带操作人信息
	DeleteAnswer(ctx context.Context, answerID int64) error

	// UpdateQuestion 代表更新问题, 只会对比其中的context，title两个字段，其他字段不会对比
	// ctx必须带操作人
	UpdateQuestion(ctx context.Context, question *Question) error
}

type Pager struct {
	Total int64 // 共有多少数据，只有返回值使用
	Start int   // 	起始位置
	Size  int   // 每个页面个数
}
