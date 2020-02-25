package async

import (
	"fmt"
	"testing"
	"time"
)

func TestAsync(t *testing.T)  {
	mgr := NewWorker(10)

	t.Log("start")
	mgr.AsynCall0(func(args []interface{}) {
		a := args[0].(int)
		b := args[1].(int)

		t.Log(a+b)
	}, func() {
		t.Log("cb")
	}, 1,3)

	mgr.AsynCall1(func(args []interface{}) error {
		a := args[0].(int)
		b := args[1].(int)

		t.Log(a+b)
		return fmt.Errorf("test error")
	}, func(err error) {
		t.Log(err)
	}, 2,3)

	mgr.AsynCall2(func(args []interface{}) (interface{}, error) {
		a := args[0].(int)
		b := args[1].(int)

		t.Log(a+b)
		return a+b, nil
	}, func(ret interface{}, err error) {
		t.Log(ret, err)
	}, 3,3)

	mgr.AsynCallN(func(args []interface{}) ([]interface{}, error) {
		a := args[0].(int)
		b := args[1].(int)

		t.Log(a+b)
		return []interface{}{a+b, "kudoo"}, nil
	}, func(rets []interface{}, err error) {
		t.Log(rets, err)
	}, 4,3)

	t.Log("end")
	time.Sleep(time.Second*5)
}
