package eleven

type Eleven struct {
	secret string
}

func New(secret string) *Eleven {
	return &Eleven{secret: secret}
}

func (e *Eleven) SetSecret(secret string) {
	e.secret = secret
}
