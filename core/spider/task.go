package spider

import (
	"io"

	"github.com/huangxiaobo/gospider/core/log"
)

type FetchTask struct {
	Url    string
	Parser Parser
}

func (t *FetchTask) OnSuccess(selector string) {
	if t.Parser == nil {
		log.Warn("task's parser is empty")
		return
	}
	go t.Parser.Parse(selector)
}

// Task Response

type TaskResponse struct {
	Body io.ReadCloser

	CSS (string)
}
