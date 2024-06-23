package res

import "github.com/yohanesmario/online-book-store/internal/api/common/errorcode"

type ErrorData struct {
	Code    errorcode.ErrorCode `json:"code"`
	Message string              `json:"message"`
}
