package govtx

type Message struct {
	receivedBody []byte
	SendBody     []byte
	cause        error
}

func (m *Message) Reply(reply []byte) {
	m.receivedBody = reply
}

func (m *Message) Fail(err error) {
	m.cause = err
}

func (m *Message) Result() AsyncResult {
	return AsyncResult{m.receivedBody, m.cause}
}
