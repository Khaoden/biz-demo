package routes

import (
	"context"
	"net/http"

	"github.com/baiyutang/gomall/app/frontend/infra/rpc"
	"github.com/baiyutang/gomall/app/frontend/kitex_gen/product"
	frontendutils "github.com/baiyutang/gomall/app/frontend/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/hertz-contrib/sessions"
)

func RegisterHome(h *server.Hertz) {
	productClient := rpc.ProductClient
	h.GET("/", func(ctx context.Context, c *app.RequestContext) {
		p, err := productClient.ListProducts(ctx, &product.ListProductsReq{})
		if err != nil {
			klog.Error(err)
		}
		var items []*product.Product
		if p != nil {
			items = p.Products
		}
		session := sessions.Default(c)
		userId := session.Get(frontendutils.UserIdKey)
		c.HTML(http.StatusOK, "home", frontendutils.WarpResponse(ctx, c, utils.H{
			"title":    "Hot sale",
			"cart_num": 10,
			"items":    items,
			"user_id":  userId,
		}))
	})
}
