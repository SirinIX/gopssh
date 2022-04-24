package direct

import (
	"encoding/json"
	"fmt"

	"cmd-scaffold/log"
	"cmd-scaffold/pkg/http"
)

const (
	URI_PRODUCT_DEPLOYED_GET = "/api/v2/product?productName=%s&parentProductName=DTinsight&clusterId=%d"
)

type RequestGetDeployedProduct struct {
	EmClusterId int    `json:"clusterId"`
	ProductName string `json:"productName"`
}

type ResponseGetDeployedProductData struct {
	Count float64              `json:"count"`
	List  []ProductInformation `json:"list"`
}

type ProductInformation struct {
	Id             int    `json:"id"`
	CanUpgrade     bool   `json:"can_upgrade"`
	ProductName    string `json:"product_name"`
	ProductVersion string `json:"product_version"`
	Status         string `json:"status"`
}

func (r *RequestBuilder) GetDeployedProduct(request *RequestGetDeployedProduct) (*ResponseGetDeployedProductData, error) {
	uri := fmt.Sprintf(r.EmUrl+URI_PRODUCT_DEPLOYED_GET, request.ProductName, request.EmClusterId)

	// Do request
	res, err := r.HttpClient.GetRequest(
		uri,
		map[string]string{
			HeaderKeyAccept: HeaderValueAccept,
		},
		map[string]string{
			CookieKeyDTStack: CookieValueDTStack,
		},
	)
	if err != nil {
		return nil, err
	}

	return toResponseGetDeployedProductData(res)
}

func toResponseGetDeployedProductData(res http.HttpResponseData) (*ResponseGetDeployedProductData, error) {
	resJson, err := json.Marshal(res)
	if err != nil {
		log.Error("failed to marshal response, error: %v", err)
		return nil, err
	}

	resData := &ResponseGetDeployedProductData{}
	if err := json.Unmarshal(resJson, resData); err != nil {
		log.Error("failed to unmarshal response to struct, error: %v", err)
		return nil, err
	}
	log.Info("successfully unmarshal response to struct, data: %+v", resData)

	return resData, nil
}
