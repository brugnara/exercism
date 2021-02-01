package flatten

// Flatten flatten an array of interfaces, removing nil
func Flatten(x interface{}) interface{} {
	ret := []interface{}{}

	for _, v := range x.([]interface{}) {
		if xi, ok := v.([]interface{}); ok {
			ret = append(ret, Flatten(xi).([]interface{})...)
			continue
		}

		if v != nil {
			ret = append(ret, v)
		}
	}

	return ret
}
