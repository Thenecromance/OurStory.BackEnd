package response

const (
	Continue = iota + 100
	SwitchingProtocols
	Processing
	EarlyHints
)

const (
	OK = iota + 200
	Created
	Accepted
	NonAuthoritativeInformation
	NoContent
	ResetContent
	PartialContent
	MultiStatus
	AlreadyReported
	IMUsed = 226
)

const (
	MultipleChoices = iota + 300
	MovedPermanently
	Found
	SeeOther
	NotModified
	UseProxy
	SwitchProxy
	TemporaryRedirect
	PermanentRedirect
)

const (
	BadRequest = iota + 400
	Unauthorized
	PaymentRequired
	Forbidden
	NotFound
	MethodNotAllowed
	NotAcceptable
	ProxyAuthenticationRequired
	RequestTimeout
	Conflict
	Gone
	LengthRequired
	PreconditionFailed
	PayloadTooLarge
	URITooLong
	UnsupportedMediaType
	RangeNotSatisfiable
	ExpectationFailed
	ImATeapot
	MisdirectedRequest = 402 + iota
	UnprocessableEntity
	Locked
	FailedDependency
	TooEarly
	UpgradeRequired
	PreconditionRequired = 403 + iota
	TooManyRequests
	RequestHeaderFieldsTooLarge = 431
	UnavailableForLegalReasons  = 451
)

const (
	InternalServerError = iota + 500
	NotImplemented
	BadGateway
	ServiceUnavailable
	GatewayTimeout
	HTTPVersionNotSupported
	VariantAlsoNegotiates
	InsufficientStorage
	LoopDetected
	NotExtended                   = 510
	NetworkAuthenticationRequired = 511
)
