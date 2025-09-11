package httpstatus

const (
	// Success
	OK                  = 200
	Created             = 201
	Accepted            = 202
	NonAutoritativeInfo = 203
	NoContent           = 204
	ResetContent        = 205
	PartialContent      = 206
	Multi               = 207
	AlreadyReported     = 208
	IMUsed              = 209

	// Client Error
	BadRequest                  = 400
	Unauthorized                = 401
	Forbidden                   = 403
	NotFound                    = 404
	MethodNotAllowed            = 405
	NotAcceptable               = 406
	RequestTimeout              = 408
	Conflict                    = 409
	Gone                        = 410
	LengthRequired              = 411
	RequestEntityTooLarge       = 413
	RequestURITooLong           = 414
	UnsupportedMediaType        = 415
	RangeNotSatisfiable         = 416
	Locked                      = 423
	UpgradeRequired             = 426
	TooManyRequests             = 429
	RequestHeaderFieldsTooLarge = 431

	// Server Errror
	InternalServerError     = 500
	NotImplemented          = 501
	BadGateway              = 502
	ServiceUnavailable      = 503
	GatewayTimeout          = 504
	HTTPVersionNotSupported = 505
	InsufficientStorage     = 507
)
