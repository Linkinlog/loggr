// Code generated by "stringer -type=ItemType"; DO NOT EDIT.

package models

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Plant-1]
	_ = x[Tool-2]
	_ = x[Seed-3]
}

const _ItemType_name = "PlantToolSeed"

var _ItemType_index = [...]uint8{0, 5, 9, 13}

func (i ItemType) String() string {
	i -= 1
	if i < 0 || i >= ItemType(len(_ItemType_index)-1) {
		return "ItemType(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _ItemType_name[_ItemType_index[i]:_ItemType_index[i+1]]
}
