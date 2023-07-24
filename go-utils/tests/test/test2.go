package test

type testInterface interface {
	GetTest() int
	SetTest(num int)
}

type test2 struct {
	testInterface testInterface
}

func NewTest2(testInterface testInterface) *test2 {
	return &test2{
		testInterface: testInterface,
	}
}

func (t *test2) HowToWorkItTest(number int) int {
	t.testInterface.SetTest(number)

	return t.testInterface.GetTest()
}
