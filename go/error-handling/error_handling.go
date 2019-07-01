package erratum

//Use opens the resource and handles the errors
func Use(o ResourceOpener, input string) (e error) {
	opener, err := o()
	for err != nil {
		if _, ok := err.(TransientError); !ok {
			return err
		}
		opener, err = o()
	}
	defer func() {
		if r := recover(); r != nil {
			switch r.(type) {
			case FrobError:
				e = r.(FrobError).inner
				opener.Defrob(r.(FrobError).defrobTag)
			case error:
				e = r.(error)
			default:
				panic(r)
			}
		}
		opener.Close()
	}()
	opener.Frob(input)
	return nil
}
