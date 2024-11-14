package middlerware

type JWT struct {
	SigningKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte("jwtSigningKey"),
	}
}
