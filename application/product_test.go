package application_test

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"github.com/viniciusbeckerbernardini/go-hexagonal/application"
	"testing"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Jorbrotota"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "O preço não é negativo né cupãnxers", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Jorbrotota"
	product.Status = application.DISABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "Só produto zeradinho né cupinxera", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Jorbrotota"
	product.Status = application.DISABLED
	product.Price = 10

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "O preço não pode ser negativo", err.Error())
}
