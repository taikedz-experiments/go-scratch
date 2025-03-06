package main

import "errors"

type ErrorOne struct {
    message string
}
func (e *ErrorOne) Error() string {
    return e.message
}
func NewErrorOne(message string) ErrorOne {
    return ErrorOne{message}
}


func main() {
    tryerror(true)
    tryerror(false)
}

func tryerror(custom bool) {
    e := throwup(custom)
    switch (*e).(type) {
        case ErrorOne:
            print("my error\n")
        default:
            print("generic error\n")
    }
}

func throwup(custom bool) error {
    if custom {
        return &NewErrorOne("blarg")
    } else {
        return errors.New("whoopsie")
    }
}
