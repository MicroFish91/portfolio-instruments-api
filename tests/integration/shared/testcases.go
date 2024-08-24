package shared

type PostTestCase struct {
	Title              string
	Payload            any
	Route              string
	ReplacementToken   string
	ExpectedStatusCode int
}

type GetTestCase struct {
	Title              string
	ParameterId        int
	Route              string
	ReplacementToken   string
	ExpectedStatusCode int
	ExpectedResponse   any
}

type PutTestCase struct {
	Title              string
	ParameterId        int
	Route              string
	ReplacementToken   string
	Payload            any
	ExpectedStatusCode int
}
