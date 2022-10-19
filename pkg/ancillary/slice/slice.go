package slice

func InsertBefore[elType any](in []elType, newEl elType, idx int) []elType {
	result := make([]elType, len(in)+1)
	copy(result, in[:idx])
	result[idx] = newEl
	copy(result[idx+1:], in[idx:])
	return result
}

func InsertAfter[elType any](in []elType, newEl elType, idx int) []elType {
	if idx >= len(in)-1 {
		result := append(in, newEl)
		return result
	}
	result := make([]elType, len(in)+1)
	copy(result, in[:idx+1])
	result[idx+1] = newEl
	copy(result[idx+2:], in[idx+1:])
	return result
}

// Prepend adds a new element to the beginning of the slice.
func Prepend[elType any](newEl elType, dst []elType) []elType {
	return append([]elType{newEl}, dst...)
}
