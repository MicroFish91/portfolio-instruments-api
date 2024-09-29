package constants

import "time"

const (
	LOCALS_REQ_PARAMS = "requestParams"
	LOCALS_REQ_BODY   = "requestBody"
	LOCALS_REQ_QUERY  = "requestQuery"
	LOCALS_REQ_USER   = "requestUser"
	LOCALS_REQ_START  = "requestStart"
	LOCALS_REQ_ID     = "requestId"
	LOCALS_LOGGER     = "contextLogger"

	TIMEOUT_MEDIUM = 5 * time.Second
	TIMEOUT_LONG   = 10 * time.Second
)
