package manager

import "github.com/ameidance/paster_facade/constant"

var (
	StatusMap map[constant.ErrorStatus]*constant.ErrorStatus
)

func bindStatus(key constant.ErrorStatus, val *constant.ErrorStatus) {
	StatusMap[key] = val
}

func init() {
	StatusMap = make(map[constant.ErrorStatus]*constant.ErrorStatus)

	bindStatus(*constant.SUCCESS, constant.HTTP_SUCCESS)
	bindStatus(*constant.ERR_WRONG_PASSWORD, constant.HTTP_ERR_WRONG_PASSWORD)
	bindStatus(*constant.ERR_RECORD_NOT_FOUND, constant.HTTP_ERR_RECORD_NOT_FOUND)
	bindStatus(*constant.ERR_SERVICE_INTERNAL, constant.HTTP_ERR_SERVICE_INTERNAL)
}

func ConvertToHttpStatus(status *constant.ErrorStatus) *constant.ErrorStatus {
	if status == nil {
		return constant.HTTP_ERR_SERVICE_INTERNAL
	}
	if httpStatus, ok := StatusMap[*status]; ok {
		return httpStatus
	}
	return constant.HTTP_ERR_SERVICE_INTERNAL
}
