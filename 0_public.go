package testX

import (
	"fmt"
	"reflect"
)

func Equal(exp interface{}, real interface{}) {
	if reflect.DeepEqual(exp, real) {
		return
	}
	panic(fmt.Sprintf(`
------------------------------
exp != real
------------------------------
Expect: %v [%T] 
Real  : %v [%T]
------------------------------`, exp, exp, real, real))
}

func Err(err error, info ...interface{}) {
	if err == nil {
		return
	}
	if len(info) == 0 {
		panic(fmt.Sprintf(`
------------------------------
err != nil
------------------------------
Error : %v
------------------------------`, err))
	}
	panic(fmt.Sprintf(`Error
------------------------------
Error : %v
Msg   : %s
------------------------------`, err, fmt.Sprint(info...)))
}

func Ok(ok bool, info ...interface{}) {
	if ok {
		return
	}
	if len(info) == 0 {
		panic(fmt.Sprintf(`
------------------------------
ok == false
------------------------------`))
	}
	panic(fmt.Sprintf(`
------------------------------
ok == false
------------------------------
Msg   : %s
------------------------------`, fmt.Sprint(info...)))
}
