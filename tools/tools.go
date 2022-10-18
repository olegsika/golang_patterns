package tools

import "errors"

func IsEqual(err, target error) bool {
	if err == nil || target == nil {
		return errors.Is(err, target)
	}

	return err.Error() == target.Error()
}
