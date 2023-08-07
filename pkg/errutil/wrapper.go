package errutil

type ErrorWrapper func(error) error

func (fn ErrorWrapper) wrap(err error) error {
	switch fn {
	case nil:
		return err
	default:
		return fn(err)
	}
}
