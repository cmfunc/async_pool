package asyncpool

type Task interface {
	Handle() error //任务处理器
}

type SimpleTask struct {
	TaskID     string                 `json:"task_id"`
	SendTime   int64                  `json:"send_time"`
	RecvTime   int64                  `json:"recv_time"`
	HandleTime int64                  `json:"handle_time"`
	AckTime    int64                  `json:"ack_time"`
	Header     map[string]interface{} `json:"header"`
	Body       interface{}            `json:"body"`
}

func (task *SimpleTask) Handle() error {
	return nil
}
