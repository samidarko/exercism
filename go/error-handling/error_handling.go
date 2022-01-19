package erratum

import "fmt"

func Use(opener ResourceOpener, input string) (err error) {
	resource, err := opener()
	if err != nil {
		if _, ok := err.(TransientError); ok {
			return Use(opener, input)
		}
		return err
	}
	defer resource.Close()
	defer func() {
		if deferErr := recover(); deferErr != nil {
			if frobError, ok := deferErr.(FrobError); ok {
				resource.Defrob(frobError.defrobTag)
				err = frobError.inner
				return
			}
			err = fmt.Errorf("%s", deferErr)
		}
	}()
	resource.Frob(input)
	return
}
