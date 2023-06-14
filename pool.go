package asyncpool

// PoolConf 协程池配置
type PoolConf struct {
	GoNum   int //goroutine的数量
	ChanLen int //channel的长度
}

type Pool struct {
	Ch     chan Task //任务队列
	Logger Logger
}

// 给协程池添加任务
func (pool *Pool) AddTask(t Task) error {
	if t == nil {
		return ErrTaskNull
	}
	pool.Ch <- t
	return nil
}

// 处理协程池任务
func (pool *Pool) HandleTask() (err error) {
	defer func() {
		err := recover()
		if err != nil {
			pool.Logger.Error("HandleTask panic %s", err)
		}
	}()
	task := <-pool.Ch
	pool.Logger.Info("HandleTask task:%+v", task)
	err = task.Handle()
	return err
}

type FuncPoolOption struct {
	f func(options *Pool)
}

func (fpo *FuncPoolOption) Apply(po *Pool) {
	fpo.f(po)
}

func NewFuncPoolOption(f func(options *Pool)) *FuncPoolOption {
	return &FuncPoolOption{
		f: f,
	}
}

func WithLogger(logger Logger) PoolOpt {
	return NewFuncPoolOption(func(o *Pool) {
		o.Logger = logger
	})
}

type PoolOpt interface {
	Apply(*Pool)
}

func NewPool(conf *PoolConf, opts ...PoolOpt) *Pool {
	pool := &Pool{
		Ch:     make(chan Task, conf.ChanLen),
		Logger: NewDefaultLogger(),
	}
	for _, opt := range opts {
		opt.Apply(pool)
	}

	for i := 0; i < conf.GoNum; i++ {
		go func() {
			for {
				pool.HandleTask()
			}
		}()
	}
	return pool
}
