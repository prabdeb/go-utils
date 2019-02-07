package parser

// Parser is the main callable struct
type Parser struct {
	deployment Deployment
}

// New the client
func New() Parser {
	return Parser{}
}
