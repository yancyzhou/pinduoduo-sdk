package goods

import "github.com/yancyzhou/pinduoduo-sdk/common"

type Goods struct {
	*common.Context
}

func NewGoods(context *common.Context) *Goods {
	service := new(Goods)
	service.Context = context
	return service
}
