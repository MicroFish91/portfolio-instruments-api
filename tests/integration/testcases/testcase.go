package testcases

type TestCase struct {
	Title              string
	ParameterId        int
	ParameterId2       int
	Payload            any
	Route              string
	ReplacementToken   string
	ExpectedStatusCode int
	ExpectedResponse   any
}
