//go:generate mockgen -package=mock_test -destination=./mocks/tes1/test_mock.go  -source=test2.go temp.com/temp/test
//go:generate mockgen -package=mock_test -destination=./mocks/tes2/test_mock.go  -source=test2.go temp.com/temp/test
//go:generate mockgen -package=mock_test -destination=./mocks/tes3/test_mock.go  -source=test2.go temp.com/temp/test
//go:generate mockgen -package=mock_test -destination=./mocks/tes4/test_mock.go  -source=test2.go temp.com/temp/test
//go:generate echo Hello, Go Generate!
package test

import (
	"testing"
)

func Test_test2(t *testing.T) {

}
