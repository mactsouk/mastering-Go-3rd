package coverage

import "testing"

func Test_f1(t *testing.T) {
	f1()
}

func Test_f2(t *testing.T) {
	_ = f2(123)
}
