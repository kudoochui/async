package async

type Job interface {
	DoIt()
	Cb()
}

type Job0 struct {
	args 	[]interface{}
	cb 		func()
	f 		func([]interface{})
}

func (job *Job0) DoIt()  {
	job.f(job.args)
}

func (job *Job0) Cb()  {
	if job.cb != nil {
		job.cb()
	}
}

type Job1 struct {
	args 	[]interface{}
	cb 		func(err error)
	err 	error
	f 		func([]interface{}) error
}

func (job *Job1) DoIt()  {
	job.err = job.f(job.args)
}

func (job *Job1) Cb()  {
	job.cb(job.err)
}

type Job2 struct {
	args 	[]interface{}
	cb 		func(interface{}, error)
	ret 	interface{}
	err 	error
	f 		func([]interface{}) (interface{}, error)
}

func (job *Job2) DoIt()  {
	job.ret, job.err = job.f(job.args)
}

func (job *Job2) Cb()  {
	job.cb(job.ret, job.err)
}

type JobN struct {
	args 	[]interface{}
	cb 		func([]interface{}, error)
	ret 	[]interface{}
	err 	error
	f 		func([]interface{}) ([]interface{}, error)
}

func (job *JobN) DoIt()  {
	job.ret, job.err = job.f(job.args)
}

func (job *JobN) Cb()  {
	job.cb(job.ret, job.err)
}