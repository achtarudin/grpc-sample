package port

type HelloPort interface {
	ValidateSayHelloRequest(name string) string
	ValidateSayManyHellosRequest(name string) string
}
