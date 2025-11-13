package errs

var DefaultCodeRelation = newCodeRelation()

type CodeError interface {
	Code() int
	Msg() string
	Detail() string
	WithDetail(detail string) CodeError
	Error
}

func NewCodeError(code int, msg string) CodeError {
	return &codeError{
		code: code,
		msg:  msg,
	}
}

type codeError struct {
	code   int
	msg    string
	detail string
}

func (e *codeError) Code() int {
	return e.code
}

func (e *codeError) Msg() string {
	return e.msg
}

func (e *codeError) Detail() string {
	return e.detail
}

func (e *codeError) WithDetail(detail string) CodeError {
	var d string
	if e.detail == "" {
		d = detail
	} else {
		d = e.detail + detail
	}
	return e.WithDetail(d)
}

func (e *codeError) Error() string {
	return e.msg
}

func (e *codeError) Is(err error) bool {
	codeErr, ok := err.(CodeError)
	if !ok {
		if err == nil && e == nil {
			return true
		}
		return false
	}
	if e == nil {
		return false
	}
	code := codeErr.Code()
	if e.code == code {
		return true
	}
	return DefaultCodeRelation.Is(e.code, code)
}

func (e *codeError) Wrap() error {
	return Wrap(e)
}

func (e *codeError) WrapMsg(msg string, kv ...interface{}) error {
	return WrapMsg(e, msg, kv...)
}

func (e *codeError) WrapLocal(kv ...interface{}) error {
	tmp := map[string]interface{}{}
	for i := 0; i < len(kv); i += 2 {
		tmp[kv[i].(string)] = kv[i+1]
	}
	message, err := i18n.Locale.Localize(&i18n.LocalizeConfig{
		MessageID:    e.msg,
		TemplateData: tmp,
	})
	if err != nil {
		return Wrap(e)
	}
	e.msg = message
	return WrapMsg(e, message)
}
