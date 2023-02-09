package slice

func InsertBefore[typ any](in []typ, newEl typ, idx int) []typ {
	result := make([]typ, len(in)+1)
	copy(result, in[:idx])
	result[idx] = newEl
	copy(result[idx+1:], in[idx:])
	return result
}

func InsertAfter[typ any](in []typ, newEl typ, idx int) []typ {
	if idx >= len(in)-1 {
		result := append(in, newEl)
		return result
	}
	result := make([]typ, len(in)+1)
	copy(result, in[:idx+1])
	result[idx+1] = newEl
	copy(result[idx+2:], in[idx+1:])
	return result
}

// Prepend adds a new element to the beginning of the slice.
func Prepend[typ any](newEl typ, dst []typ) []typ {
	return append([]typ{newEl}, dst...)
}
