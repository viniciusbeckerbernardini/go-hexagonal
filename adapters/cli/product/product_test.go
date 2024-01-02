package cli_test

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	cli "github.com/viniciusbeckerbernardini/go-hexagonal/adapters/cli/product"
	mock_application "github.com/viniciusbeckerbernardini/go-hexagonal/application/mocks"
	"testing"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "Product 1"
	productPrice := 10.0
	productStatus := "enabled"
	productId := "abc"

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(productMock).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(productMock).Return(productMock, nil).AnyTimes()

	resultExpected := fmt.Sprintf("Product ID %s with the name %s has been created with the price %f and status %s", productId, productName, productPrice, productStatus)
	result, err := cli.Run(service, "create", "", productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	result, err = cli.Run(service, "enable", productId, "", 10)
	resultExpected = fmt.Sprintf("Product %s has been enabled", productName)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	result, err = cli.Run(service, "disable", productId, "", 0)
	resultExpected = fmt.Sprintf("Product %s has been disabled", productName)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

}
