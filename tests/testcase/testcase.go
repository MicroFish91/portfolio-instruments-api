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

type PutTestCase struct {
	Title              string
	ParameterId        int
	Route              string
	ReplacementToken   string
	Payload            any
	ExpectedStatusCode int
}
