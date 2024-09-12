package shared

type PostTestCase struct {
	Title              string
	Payload            any
	Route              string
	ReplacementToken   string
	ExpectedStatusCode int
	ExpectedResponse   any
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
	ParameterId2       int
	Route              string
	ReplacementToken   string
	Payload            any
	ExpectedStatusCode int
}

type DeleteTestCase struct {
	Title              string
	Route              string
	ParameterId        int
	ReplacementToken   string
	ExpectedStatusCode int
}
