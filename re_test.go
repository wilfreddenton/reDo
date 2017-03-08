package re

import (
	"fmt"
	"testing"
)

func test1(i int) error {
	if i < 2 {
		return fmt.Errorf("error")
	}
	return nil
}

func test2(i int) error {
	return fmt.Errorf("error")
}

func TestDo(t *testing.T) {
	count := 0
	err := Do(func(i int) (bool, error) {
		count += 1
		err := test1(i)
		return i < 2, err
	})
	if err != nil {
		t.Errorf("err should be nil")
	}
	if count != 2 {
		t.Errorf("count should be 2")
	}

	count = 0
	err = Do(func(i int) (bool, error) {
		count += 1
		err = test2(i)
		return i < 3, err
	})
	if err == nil {
		t.Error("err should not be nil")
	}
	if count != 3 {
		t.Errorf("count should be 3")
	}
}
