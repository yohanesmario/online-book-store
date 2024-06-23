package res

type BaseResponse[TData any] struct {
	Data TData `json:"data"`
}
