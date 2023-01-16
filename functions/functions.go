package functions

func addValues(v1 interface{}, v2 interface{}) interface{} {
	var result interface{}
	switch v1.(type) {
	case int:
		switch v2.(type) {
		case int:
			result = v1.(int) + v2.(int)
		case float64:
			result = float64(v1.(int)) + v2.(float64)
		}
	case float64:
		switch v2.(type) {
		case int:
			result = v1.(float64) + float64(v2.(int))
		case float64:
			result = v1.(float64) + v2.(float64)
		}
	}
	return result
}

func Sum(arg ...interface{}) interface{} {

	var tot interface{}
	tot = 0
	for i := range arg {
		tot = addValues(tot, arg[i])
	}

	return tot
}

func SSum(arg ...interface{}) interface{} {

	var tot float64
	for i := range arg {
		v := arg[i]

		switch v.(type) {
		case int:
			tot += float64(v.(int))
		case float64:
			tot += v.(float64)
		}

	}

	return tot
}
