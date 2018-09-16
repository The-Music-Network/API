package errs

import (
	"runtime"
	"strings"
	"errors"
	"log"
	"bytes"
	"fmt"
)

type Error struct {
	root   error
	code   int
	frames []frame
}

type frame struct {
	fnc  string
	line int
	file string
}

func (f *frame) String() string {
	return fmt.Sprintf("%s:%d\t(%s)\n", f.file, f.line, f.fnc)
}

func New(statusCode int, message string) *Error {
	return &Error{
		root:   errors.New(message),
		frames: []frame{newFrame()},
		code: statusCode,
	}
}

func (this *Error) Root() error {
	return this.root
}

// This method implements the error interface defined in the Go SDK.
func (err *Error) Error() string {
	b := bytes.NewBuffer(nil)
	fmt.Fprintln(b, err.root)
	for _, f := range err.frames {
		fmt.Fprint(b, f.String())
	}
	return b.String()
}

func (this *Error) Status() int {
	return this.code
}

func Stack(err error) *Error {
	log.Println("err:", err)
	if err == nil {
		return nil
	}

	newFrame := newFrame()

	if err_, ok := err.(*Error); ok && err_ != nil {
		err_.frames = append(err_.frames, newFrame)
		return err_
	}

	newError := &Error{
		root:   err,
		frames: []frame{newFrame},
	}
	log.Println("newError:", newError)

	return newError
}

func newFrame() frame {
	pc, file, line, _ := runtime.Caller(2)
	func_ := runtime.FuncForPC(pc)
	funcChunks := strings.Split(func_.Name(), "/")
	funcName := funcChunks[len(funcChunks)-1]

	f := frame{
		fnc:  funcName,
		line: line,
		file: file,
	}

	return f
}
