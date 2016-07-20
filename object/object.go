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

func (obj Object) LessOther(other Object, i, j First) bool {
	return obj[i] < other[j]
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

/*
Equals return true if all keys and values equal
*/
func (obj Object) Equals(other Object) bool {
	if len(obj) != len(other) {
		return false
	}
	for k := range obj {
		if obj.LessOther(other, k, k) || other.LessOther(obj, k, k) {
			return false
		}
	}
	return true
}
