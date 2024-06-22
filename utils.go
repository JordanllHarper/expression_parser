package main

func OutOfRange[T any](items []T, index int) bool {
	return index > len(items)-1
}
