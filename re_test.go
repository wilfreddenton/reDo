package re

import (
	"fmt"
	"testing"
)

func test1(i int) error {
	if i == 0 {
		return fmt.Errorf("error")
	}
	return nil
}

func test2(i int) error {
	return fmt.Errorf("error")
}

func TestDo(t *testing.T) {
	count := 0
	n := 2
	err := Do(n, func(i int) error {
		count += 1
		return test1(i)
	})
	if err != nil {
		t.Errorf("err should be nil")
	}
	if count != n {
		t.Errorf("count should be %d", n)
	}

	count = 0
	n = 3
	err = Do(n, func(i int) error {
		count += 1
		return test2(i)
	})
	if err == nil {
		t.Error("err should not be nil")
	}
	if count != n {
		t.Errorf("count should be %d", 3)
	}
}
