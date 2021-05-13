package ddk

import (
	"github.com/yancyzhou/pinduoduo-sdk/common"
	"github.com/yancyzhou/pinduoduo-sdk/util"
)

type GoodsZsUnitGenerateResponse struct {
	MobileShortUrl           string `json:"mobile_short_url"`
	MobileUrl                string `json:"mobile_url"`
	MultiGroupMobileShortUrl string `json:"multi_group_mobile_short_url"`
	MultiGroupMobileUrl      string `json:"multi_group_mobile_url"`
	MultiGroupShortUrl       string `json:"multi_group_short_url"`
	MultiGroupUrl            string `json:"multi_group_url"`
	ShortUrl                 string `json:"short_url"`
	Url                      string `json:"url"`
}

type GoodsZsUnitUrlGenResult struct {
	GoodsZsUnitGenerateResponse GoodsZsUnitGenerateResponse `json:"goods_zs_unit_generate_response"`
	common.CommonResult
}

type GoodsZsUnitUrlGenParams struct {
	Pid              string `json:"pid"`
	CustomParameters string `json:"custom_parameters"`
	SourceUrl        string `json:"source_url"`
}

func (this *DuoduoKe) GoodsZsUnitUrlGen(p *GoodsZsUnitUrlGenParams) (*GoodsZsUnitUrlGenResult, error) {
	apiType := `pdd.ddk.goods.zs.unit.url.gen`
	params, paramsURL := util.FormatURLParams(p)
	url := this.GetURL(apiType, "", params, paramsURL)
	var result GoodsZsUnitUrlGenResult
	err := util.HttpPOST(url, nil, &result)
	if err != nil {
		return nil, err
	}
	err = common.CheckErrCode(result.CommonResult)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
