package flatten

func Flatten(deepList []interface{}) (flattened []interface{}) {
	flattened = make([]interface{}, 0)
	for _, element := range deepList {
		switch i := element.(type) {
		case []interface{}:
			flattened = append(flattened, Flatten(element)...)
		}
	}
	return
}

func flatten(deepList []interface{}, outList *[]interface{}) {
	for _, element := range deepList{
		
	}
}