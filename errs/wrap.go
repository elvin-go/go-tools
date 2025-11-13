package errs

import "github.com/pkg/errors"

func Unwrap(err error) error {
	for err != nil {
		unwarp, ok := err.(interface{ Unwrap() error })
		if !ok {
			break
		}
		err = unwarp.Unwrap()
	}
	return err
}

func Wrap(err error) error {
	return errors.WithStack(err)
}

func WrapMsg(err error, msg string, kv ...any) error {
	if err == nil {
		return nil
	}
	withMessage := errors.WithMessage(err, toString(msg, kv))
	return errors.WithStack(withMessage)
}
