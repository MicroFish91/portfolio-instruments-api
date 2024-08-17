package testcase

type PostTestCase struct {
	Title              string
	Payload            any
	ExpectedStatusCode int
}

type GetTestCase struct {
	Title              string
	ParameterId        int
	Route              string
	ReplacementToken   string
	ExpectedStatusCode int
}
