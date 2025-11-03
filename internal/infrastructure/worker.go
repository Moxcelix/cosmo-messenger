package infrastructure

type Worker interface {
	Run()
}

type Workers []Worker

func NewWorkers() Workers {
	return Workers{}
}

func (u Workers) Run() {
	for _, worker := range u {
		worker.Run()
	}
}
