package test

import (
	"fmt"
	"testing"

	"github.com/yancyzhou/pinduoduo-sdk/common"
	"github.com/yancyzhou/pinduoduo-sdk/ddk"
)

func TestMemeber(t *testing.T) {
	pdd := ddk.NewDuoduoKe(common.NewContext("5a5813b638314d1792b711e116fc65da", "e5f7992cd3f0b7652733932235214488c3b1d9e6"))
	var query ddk.GoodsSearchParams
	s := "iphone"
	pid := "15983074_207918426"
	query.Keyword = &s
	query.Pid = &pid
	c := "{'uid':'vincent'}"
	var q ddk.MemberAuthorityQueryParams
	q.Pid = &pid
	q.CustomParameters = &c
	query.CustomParameters = &c
	result, err := pdd.GoodsSearch(&query)
	fmt.Println(result.GoodsSearchResponse.GoodsList[0].GoodsID, result.GoodsSearchResponse.GoodsList[0].GoodsSign)
	authorityQuery, err := pdd.MemberAuthorityQuery(&q)
	if err != nil {
		fmt.Println("err", err)
	}
	if authorityQuery.AuthorityQueryResponse.Bind == 1 {
		fmt.Println("Authority is Success")
	} else {
		var goodsq ddk.RpPromUrlGenerateParams
		ChannelType := 10
		goodsq.ChannelType = &ChannelType
		goodsq.PIdList = &[]string{pid}
		goodsq.CustomParameters = &c
		generate, err := pdd.RpPromUrlGenerate(&goodsq)
		if err != nil {
			fmt.Println("GoodsPromotionUrlGenerate err", err)
		}
		fmt.Println(generate.RpPromotionUrlGenerateResponse.UrlList)
	}
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
		return
	}
}
