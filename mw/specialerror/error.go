package specialerror

import "go-tools/errs"

var handlers []func(err error) errs.CodeError

func AddErrHandler(f func(err error) errs.CodeError) (err error) {
	if f == nil {
		return errs.New("nil handle")
	}
	handlers = append(handlers, f)
	return nil
}

func AddReplace(target error, codeError errs.CodeError) error {
	handler := func(err error) errs.CodeError {
		if err == target {
			return codeError
		}
		return nil
	}

	if err := AddErrHandler(handler); err != nil {
		return err
	}

	return nil
}

func ErrCode(err error) errs.CodeError {
	if codeErr, ok := err.(errs.CodeError); ok {
		return codeErr
	}
	for i := 0; i < len(handlers); i++ {
		if codeErr := handlers[i](err); codeErr != nil {
			return codeErr
		}
	}
	return nil
}
