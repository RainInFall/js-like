package object

// template type Object(First, Second, FirstArray, SecondArray)
type First int
type Second string
type FirstArray []First
type SecondArray []Second

//Object has js-like functions
type Object map[First]Second

/********Self defination Begin*******************/

func (obj Object) Less(i, j First) bool {
	return obj[i] < obj[j]
}

/********Self defination End*******************/

/*
Keys return array of keys of the Obecjt
*/
func (obj Object) Keys() FirstArray {
	keys := make(FirstArray, 0, len(obj))
	for key := range obj {
		keys = append(keys, key)
	}
	return FirstArray(keys)
}
