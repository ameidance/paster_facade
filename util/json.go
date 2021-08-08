package util

import (
	"github.com/cloudwego/kitex/pkg/klog"
	jsoniter "github.com/json-iterator/go"
)

func GetJsonMapFromStruct(obj interface{}) map[string]interface{} {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	data, err := json.Marshal(obj)
	if err != nil {
		klog.Errorf("[GetJsonMapFromStruct] marshal failed. err:%v", err)
		return nil
	}
	m := make(map[string]interface{})
	err = json.Unmarshal(data, &m)
	if err != nil {
		klog.Errorf("[GetJsonMapFromStruct] unmarshal failed. err:%v", err)
		return nil
	}
	return m
}
