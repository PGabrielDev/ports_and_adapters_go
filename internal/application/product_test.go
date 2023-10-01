package application_test

import (
	"github.com/PGabrielDev/ports_and_adapters_go/internal/application"
	"github.com/stretchr/testify/require"

	"testing"
)

func TestProductEnable(t *testing.T) {
	product := application.Product{}
	product.Name = "Olá"
	product.Status = application.DISABLE
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.NotNil(t, "Price precisa ser maior que 0", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	produto := application.Product{}
	produto.Name = "Lapis"
	produto.Price = 60

	err := produto.Disable()
	require.Equal(t, "Prico precisa ser 0", err.Error())

	produto.Price = 0

	err = produto.Disable()

	require.Nil(t, err)
}
