package v20190819

import (
	"encoding/json"
	"reflect"
)

type (
	ProductInfo struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}
	GoodsDetail struct {
		TimeUnit         string        `json:"timeUnit"`
		TimeSpan         int64         `json:"timeSpan"`
		ProductCode      string        `json:"productCode"`
		SubProductCode   string        `json:"subProductCode"`
		Pid              int64         `json:"pid"`
		ClusterId        int64         `json:"clusterId"`
		AutoRenewFlag    int64         `json:"autoRenewFlag"`
		ProductInfo      []ProductInfo `json:"productInfo"`
		InstanceName     string        `json:"instanceName"`
		VpcId            string        `json:"vpcId"`
		SubnetId         string        `json:"subnetId"`
		MsgRetentionTime int64         `json:"msgRetentionTime"`
		ExData           map[string]interface{} 
	}
	ProductsDetail struct {
		CVM           int      `json:"cvm"`
		Tags          []string `json:"tags"`
		BrokerVersion string   `json:"brokerVersion"`
		IPType        int      `json:"ipType"`
		InstanceType  string   `json:"instanceType"`
		MultiZoneFlag bool     `json:"multiZoneFlag"`
		ZoneIdList    []int    `json:"zoneIdList"`
	}
)

func (gd *GoodsDetail) MarshalJSON() ([]byte, error) {
	var configAsMap = make(map[string]interface{})
	objValue := reflect.ValueOf(*gd)
	objType := objValue.Type()

	for i := 0; i < objValue.NumField(); i++ {
		field := objType.Field(i)
		value := objValue.Field(i).Interface()
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" {
			continue
		}
		configAsMap[jsonTag] = value
	}
	
	// ExData
	if gd.ExData != nil {
		for k, v := range gd.ExData {
			configAsMap[k] = v
		}
	}
	return json.Marshal(configAsMap)
}