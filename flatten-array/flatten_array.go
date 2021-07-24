package flatten

func Flatten(input interface{}) []interface{} {
	switch e := input.(type) {
	case []interface{}:
		output := []interface{}{}
		for _, item := range e {
			output = append(output, Flatten(item)...)
		}
		return output
	case nil:
		return []interface{}{}
	}
	return []interface{}{input}
}
