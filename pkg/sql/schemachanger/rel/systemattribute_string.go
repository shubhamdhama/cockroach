// Code generated by "stringer -type systemAttribute"; DO NOT EDIT.

package rel

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Type-63]
	_ = x[Self-62]
}

const _systemAttribute_name = "SelfType"

var _systemAttribute_index = [...]uint8{0, 4, 8}

func (i systemAttribute) String() string {
	i -= 62
	if i < 0 || i >= systemAttribute(len(_systemAttribute_index)-1) {
		return "systemAttribute(" + strconv.FormatInt(int64(i+62), 10) + ")"
	}
	return _systemAttribute_name[_systemAttribute_index[i]:_systemAttribute_index[i+1]]
}
