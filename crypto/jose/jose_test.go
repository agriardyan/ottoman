package jose_test

import (
	"testing"

	"github.com/bukalapak/ottoman/crypto/jose"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		_, err := jose.New("", rsaPrivateKey)
		assert.NotNil(t, err)

		_, err = jose.New(rsaPublicKey, "")
		assert.NotNil(t, err)

		_, err = jose.New("", "")
		assert.NotNil(t, err)
	})

	t.Run("NewSignature", func(t *testing.T) {
		_, err := jose.NewSignature("", rsaPrivateKey)
		assert.NotNil(t, err)

		_, err = jose.NewSignature(rsaPublicKey, "")
		assert.NotNil(t, err)
	})

	t.Run("NewEncryption", func(t *testing.T) {
		_, err := jose.NewEncryption("", rsaPrivateKey)
		assert.NotNil(t, err)

		_, err = jose.NewEncryption(rsaPublicKey, "")
		assert.NotNil(t, err)
	})
}

func TestSignature(t *testing.T) {
	b := []byte(`{"foo":"bar"}`)
	n, err := jose.New(rsaPublicKey, rsaPrivateKey)
	assert.Nil(t, err)

	token, err := n.Encode(b)
	assert.Nil(t, err)
	assert.NotNil(t, token)

	data, err := n.Decode(token)
	assert.Nil(t, err)
	assert.Equal(t, b, data)

	out, err := n.Decode("x")
	assert.NotNil(t, err)
	assert.Nil(t, out)

	data2, err := jose.Decode(rsaPublicKey, token)
	assert.Nil(t, err)
	assert.Equal(t, b, data2)

	data3, err := jose.Decode("", token)
	assert.NotNil(t, err)
	assert.Nil(t, data3)
}

func TestEncryption(t *testing.T) {
	b := []byte(`{"foo":"bar"}`)
	n, err := jose.New(rsaPublicKey, rsaPrivateKey)
	assert.Nil(t, err)

	token, err := n.Encrypt(b)
	assert.Nil(t, err)
	assert.NotNil(t, token)

	data, err := n.Decrypt(token)
	assert.Nil(t, err)
	assert.Equal(t, b, data)

	out, err := n.Decrypt("x")
	assert.NotNil(t, err)
	assert.Nil(t, out)

	data2, err := jose.Decrypt(rsaPrivateKey, token)
	assert.Nil(t, err)
	assert.Equal(t, b, data2)

	data3, err := jose.Decrypt("", token)
	assert.NotNil(t, err)
	assert.Nil(t, data3)
}
