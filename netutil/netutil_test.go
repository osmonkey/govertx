package netutil

import (
	"testing"
)

func Test_getNextFree(t *testing.T) {
	u := NetUtil{}
	port, err := u.GetNextFree()
	if err != nil {
		t.Log(port)
	} else {
		t.Error(err.Error())
		t.Fail()
	}
}
