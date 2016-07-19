package object

// template type Object(First, Second, FirstArray, SecondArray)
type First int
type Second string
type FirstArray []int
type SecondArray []string

//Object has js-like functions
type Object map[First]Second

/*
Keys return array of keys of the Obecjt
*/
func (obj Object) Keys() FirstArray {
	keys := make([]int, 0, len(obj))
	for key := range obj {
		keys = append(keys, int(key))
	}
	return FirstArray(keys)
}
