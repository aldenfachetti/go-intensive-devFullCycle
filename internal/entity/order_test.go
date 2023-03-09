package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_If_Its_Gets_An_Error_If_ID_Is_Blank(t *testing.T) {
	order := Order {}

	//Native way to treat an error in Go
	// err := order.Validate()
	// if err == nil {
	// 	t.Error("Expected an error, but got nil")
	// }

	//Alternative way to treat an error, using  a library
	assert.Error(t, order.Validate(), "Invalid ID")
}

func Test_If_Its_Gets_An_Error_If_Price_Is_Blank(t *testing.T) {
	order := Order {ID: "123"}
	assert.Error(t, order.Validate(), "Invalid Price")
}

func Test_If_Its_Gets_An_Error_If_Tax_Is_Blank(t *testing.T) {
	order := Order {ID: "123", Price: 10.0}
	assert.Error(t, order.Validate(), "Invalid Tax")
}

func Test_With_All_Valid_Params(t *testing.T) {
	order := Order {ID: "123", Price: 10.0, Tax: 2.0}
	assert.NoError(t, order.Validate())
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 2.0, order.Tax)
	order.CalculateFinalPrice()
	assert.Equal(t, 12.0, order.FinalPrice)
}
