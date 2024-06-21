package expressionparser

func OutOfRange[T any](items []T, index int) bool {
	return len(items)-1 < index

}
