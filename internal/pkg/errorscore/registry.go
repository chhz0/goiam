package errorscore

import (
	"fmt"
	"regexp"
	"strconv"
	"sync"

	"github.com/novalagung/gubrak"
)

type ErrorsCoder interface {
	Code() int
	HttpStauts() int
	Message() string
	Any() any
}

type errorsNo struct {
	errorCode int
	Http      int
	Msg       string
	AnyData   any
}

func (e *errorsNo) Code() int { return e.errorCode }

func (e *errorsNo) HttpStauts() int { return e.Http }

func (e *errorsNo) Message() string { return e.Msg }

func (e *errorsNo) Any() any { return e.AnyData }

var codeRegistry = make(map[int]ErrorsCoder, 0)
var codeMutex = sync.Mutex{}

func registerE(errorsCoder ErrorsCoder) {
	codeMutex.Lock()
	defer codeMutex.Unlock()
	codeRegistry[errorsCoder.Code()] = errorsCoder
}

func mustRegister(errorsCoder ErrorsCoder) {
	codeMutex.Lock()
	defer codeMutex.Unlock()

	if _, ok := codeRegistry[errorsCoder.Code()]; ok {
		panic(fmt.Sprintf("errorscore.mustregistry err: code %d already registered", errorsCoder.Code()))
	}

	codeRegistry[errorsCoder.Code()] = errorsCoder
}

var allowHttpStauts = []int{200, 401, 403, 404, 500}

func registerCode(code int, httpStauts int, msg string, refs ...string) {

	found, _ := gubrak.Includes(allowHttpStauts, httpStauts)
	if !found {
		panic(fmt.Sprintf("errorscore.registerCode err: http status %d not allowed, must be %+v", httpStauts, allowHttpStauts))
	}

	var reference string
	if len(refs) > 0 {
		reference = refs[0]
	}

	mustRegister(&errorsNo{
		errorCode: code,
		Http:      httpStauts,
		Msg:       msg,
		AnyData:   reference,
	})
}

// ParseErrorsToCoder 解析错误码
// 注意：err 必须是 errors.WithCode 错误，因为 ParseErrorsToCoder
// 是按照 error.Error() 返回的 string 来解析的，
// errors.WithCode 中的 Error() 返回的格式为: "code(1001): error message"
func ParseErrorToCoder(err error) ErrorsCoder {
	if err == nil {
		return nil
	}

	codeStr := getCodeFromStr(err.Error())

	codeInt, err := strconv.Atoi(codeStr)
	if err != nil {
		return nil
	}

	errorsCode, ok := codeRegistry[codeInt]
	if !ok {
		return nil
	}
	return errorsCode
}

func getCodeFromStr(codeStr string) string {
	re := regexp.MustCompile(`code\((\d+)\)`)

	matches := re.FindStringSubmatch(codeStr)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}
