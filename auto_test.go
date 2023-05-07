package testX

import (
	"errors"
	"fmt"
	"testing"
)

func TestOk(t *testing.T) {
	Ok(true)
	err := runWithRecover(func() {
		Ok(false)
	})
	if err == nil {
		panic("not ok")
	}
}

func TestEqual(t *testing.T) {
	Equal(1, 1)
	err := runWithRecover(func() {
		Equal(1, 2)
	})
	Ok(err != nil)
}

func TestErr(t *testing.T) {
	Err(nil)
	err := runWithRecover(func() {
		Err(errors.New("1"))
	})
	Ok(err != nil)
}

func runWithRecover(callback func()) (err error) {
	defer func() {
		panicInfo := recover()
		if panicInfo == nil {
			return
		}
		err = errors.New(fmt.Sprintf("%v", panicInfo))
	}()
	callback()
	return
}
