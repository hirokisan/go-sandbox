package sample

import (
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/hirokisan/go-sandbox/swagger/gen/restapi/operations/items"
)

// API実装を記述する
// todo: もっと良いやり方を考える
func ListItems(params items.ListItemsParams) middleware.Responder {
	return middleware.NotImplemented("operation items.ListItems has not yet been implemented")
}
