package re

func Do(f func(int) (bool, error)) error {
	for i := 1; ; i += 1 {
		b, e := f(i)
		if e == nil || !b {
			return e
		}
	}
}
