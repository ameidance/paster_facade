package util

import (
	"reflect"

	"github.com/ameidance/paster_facade/constant"
	"github.com/cloudwego/kitex/pkg/klog"
)

func FillBizResp(resp interface{}, status *constant.ErrorStatus) {
	if resp == nil {
		return
	}
	if status == nil {
		status = constant.ERR_SERVICE_INTERNAL
	}

	respType := reflect.TypeOf(resp)
	respVal := reflect.ValueOf(resp)
	if respType.Kind() == reflect.Ptr {
		respVal = respVal.Elem()
	}

	respVal.FieldByName("StatusCode").SetInt(int64(status.StatusCode))
	respVal.FieldByName("StatusMessage").SetString(status.StatusMsg)
}

func IsStatusSuccess(status *constant.ErrorStatus) bool {
	return status == nil || status == constant.SUCCESS || status.StatusCode == 0
}

func CheckRpcResponse(resp interface{}, err error) *constant.ErrorStatus {
	if err != nil {
		klog.Errorf("[CheckRpcResponse] resp:%v, err:%v", resp, err)
		return constant.ERR_SERVICE_INTERNAL
	}
	if resp == nil {
		klog.Errorf("[CheckRpcResponse] resp is nil", resp)
		return constant.ERR_SERVICE_INTERNAL
	}

	v := reflect.ValueOf(resp)
	for v.Kind() == reflect.Ptr {
		v = reflect.Indirect(v)
	}
	if IsNil(v) {
		klog.Errorf("[CheckRpcResponse] resp is nil", resp)
		return constant.ERR_SERVICE_INTERNAL
	}

	statusCode := v.FieldByName("StatusCode").Interface().(int32)
	var msg string
	msgTypeV := v.FieldByName("StatusMessage")
	if !IsNil(msgTypeV) {
		msg = reflect.Indirect(msgTypeV).Interface().(string)
	}

	if statusCode != 0 {
		return &constant.ErrorStatus{
			StatusCode: statusCode,
			StatusMsg:  msg,
		}
	}
	return nil
}

func IsNil(v reflect.Value) bool {
	if v.IsValid() {
		switch v.Kind() {
		case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
			return v.IsNil()
		default:
			return false
		}
	}
	return true
}
