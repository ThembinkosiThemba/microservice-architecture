package shop_test

import (
	"testing"

	"monolith-microservice/pkg/common/price"
	"monolith-microservice/pkg/orders/domain/orders"
	"monolith-microservice/pkg/orders/infrastructure/shop"
	"monolith-microservice/pkg/shop/interfaces/private/intraprocess"

	"github.com/stretchr/testify/assert"
)

func TestOrderProductFromShopProduct(t *testing.T) {
	shopProduct := intraprocess.Product{
		"123",
		"name",
		"desc",
		price.NewPriceP(42, "EUR"),
	}
	orderProduct, err := shop.OrderProductFromIntraprocess(shopProduct)
	assert.NoError(t, err)

	expectedOrderProduct, err := orders.NewProduct("123", "name", price.NewPriceP(42, "EUR"))
	assert.NoError(t, err)

	assert.EqualValues(t, expectedOrderProduct, orderProduct)
}
