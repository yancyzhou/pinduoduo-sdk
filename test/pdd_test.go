package test

import (
	"fmt"
	"testing"

	"github.com/yancyzhou/pinduoduo-sdk/common"
	"github.com/yancyzhou/pinduoduo-sdk/ddk"
)

func TestMemeber(t *testing.T) {
	pdd := ddk.NewDuoduoKe(common.NewContext("", ""))
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
