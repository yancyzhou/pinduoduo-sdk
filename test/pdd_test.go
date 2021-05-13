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
	c := "{'uid':'5ff5e09241141f0001fa3865'}"
	var q ddk.GoodsZsUnitUrlGenParams
	q.CustomParameters = c
	q.Pid = pid
	q.SourceUrl = "https://p.pinduoduo.com/WFpFzQwq"
	goodsZsUnitUrlGenResult, err := pdd.GoodsZsUnitUrlGen(&q)
	if err != nil {
		fmt.Errorf("error is %t",err)

	}
	fmt.Println(goodsZsUnitUrlGenResult.GoodsZsUnitGenerateResponse)

}
