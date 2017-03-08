package re

func Do(n int, f func(int) error) error {
	var e error
	for i := 0; i < n; i += 1 {
		if e = f(i); e == nil {
			return e
		}
	}
	return e
}
