package govtx

type AsyncResult struct {
	Result []byte
	Cause  error
}

func (as *AsyncResult) Succeeded() bool {
	return as.Result != nil
}

func (as *AsyncResult) Failed() bool {
	return as.Cause != nil
}
