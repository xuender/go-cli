package utils_test

type Map[K comparable, V any] map[K]V

func (p Map[K, V]) Has(key K) bool {
	_, has := p[key]

	return has
}
