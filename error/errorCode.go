package error

var (
	Success            = NewError(00000000, "success")
	ServerError        = NewError(10000001, "internal error")
	InvalidParams      = NewError(10000002, "invalid params")
	UnauthorizedRights = NewError(10000003, "invalid author")
	AuthorizedExpire   = NewError(10000004, "auth expire")
)
