package logger

import "fmt"

func Wrapf(err error, format string, args ...interface{}) error {
	e := stack{
		prev:   err,
		this:   fmt.Sprintf(format, args...),
		caller: GetCaller(2),
	}
	return e
}

type IErrorWithCaller interface {
	Caller() Caller
}

type IErrorWithPrev interface {
	Prev() error
}

type stack struct {
	error
	IErrorWithCaller
	IErrorWithPrev
	prev   error
	this   string
	caller Caller
}

func (s stack) Error() string {
	return s.this
}

func (s stack) Caller() Caller {
	return s.caller
}

func (s stack) Prev() error {
	return s.prev
}

func (s stack) Format(f fmt.State, verb rune) {
	//verb: %s for last error only, %e for first error only, %v for full stack
	//%+e/%+s/%+v to also show file and line nr
	var e string
	plus := f.Flag(int('+')) //true when formatting with %+v or %+s or %+e
	switch verb {
	case 's':
		if plus {
			c := s.caller
			e = fmt.Sprintf("%s(%d): %s", c.PackageFile(), c.Line(), s.this)
		} else {
			e = s.this
		}
	case 'e':
		//rewind to first error
		var err error
		err = s
		for {
			errIsStack, ok := err.(IErrorWithPrev)
			if !ok {
				break
			}
			err = errIsStack.Prev()
		}
		if plus {
			if errWithCaller, ok := err.(IErrorWithCaller); ok {
				c := errWithCaller.Caller()
				e = fmt.Sprintf("%s(%d): %s", c.PackageFile(), c.Line(), err.Error())
			} else {
				e = err.Error()
			}
		} else {
			e = err.Error()
		}
	default: //also %v to show stack
		var err error
		err = s
		count := 0
		for err != nil {
			if count > 0 {
				e += "; "
			}
			count++
			//append this to the stack:
			if plus {
				if errWithCaller, ok := err.(IErrorWithCaller); ok {
					c := errWithCaller.Caller()
					e += fmt.Sprintf("%s(%d): %s", c.PackageFile(), c.Line(), err.Error())
				} else {
					e += err.Error()
				}
			} else {
				e += err.Error()
			}

			//go one level up
			if errWithPrev, ok := err.(IErrorWithPrev); ok {
				err = errWithPrev.Prev()
			} else {
				break //no prev
			}
		}
	}
	f.Write([]byte(e))
}
