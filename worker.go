package async

// Goroutine safe
type Worker struct {
	closeChan 		chan bool
	chanJob 		chan Job
	chanRet 		chan Job
}

// Create a goroutine safe job worker
func NewWorker(chanSize int) *Worker {
	mgr := &Worker{
		closeChan: make(chan bool, 0),
		chanJob: make(chan Job, chanSize),
		chanRet: make(chan Job, chanSize),
	}

	go mgr.work()
	return mgr
}

func (mgr *Worker) work() {
	for {
		select {
		case <- mgr.closeChan:
			return
		case job := <- mgr.chanJob:
			job.DoIt()
			mgr.chanRet <- job
		case ret := <- mgr.chanRet:
			ret.Cb()
		}
	}
}

func (mgr *Worker) Dispose()  {
	mgr.closeChan <- true
}

func (mgr *Worker) AddJob(job Job)  {
	mgr.chanJob <- job
}

// Asynchronous call without returns
func (mgr *Worker) AsynCall0(f func([]interface{}), cb func(), args ...interface{}) {
	job := &Job0{
		args: args,
		cb:   cb,
		f:    f,
	}
	mgr.AddJob(job)
}

// Asynchronous call with a error return
func (mgr *Worker) AsynCall1(f func([]interface{}) error, cb func(error), args ...interface{}) {
	job := &Job1{
		args: args,
		cb:   cb,
		f:    f,
	}
	mgr.AddJob(job)
}

// Asynchronous call with two returns
func (mgr *Worker) AsynCall2(f func([]interface{}) (interface{},error), cb func(interface{},error), args ...interface{}) {
	job := &Job2{
		args: args,
		cb:   cb,
		f:    f,
	}
	mgr.AddJob(job)
}

// Asynchronous call with n returns
func (mgr *Worker) AsynCallN(f func([]interface{}) ([]interface{},error), cb func([]interface{},error), args ...interface{}) {
	job := &JobN{
		args: args,
		cb:   cb,
		f:    f,
	}
	mgr.AddJob(job)
}