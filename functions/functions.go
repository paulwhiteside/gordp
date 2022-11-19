package functions

func Sum(arg ...interface{}) interface{} {

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
