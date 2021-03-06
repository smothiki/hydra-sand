package jwk

import (
	"crypto/x509"

	"github.com/ory/hydra/rand/sequence"
	"github.com/pkg/errors"
	"github.com/square/go-jose"
)

type HS256Generator struct {
	Length int
}

func (g *HS256Generator) Generate(id string) (*jose.JSONWebKeySet, error) {
	if g.Length < 12 {
		g.Length = 12
	}

	if id == "" {
		id = "shared"
	}

	key, err := sequence.RuneSequence(g.Length, []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789,.-;:_#+*!§$%&/()=?}][{<>"))
	if err != nil {
		return nil, errors.Errorf("Could not generate key because %s", err)
	}

	return &jose.JSONWebKeySet{
		Keys: []jose.JSONWebKey{
			{
				Key:          []byte(string(key)),
				KeyID:        id,
				Certificates: []*x509.Certificate{},
			},
		},
	}, nil
}
