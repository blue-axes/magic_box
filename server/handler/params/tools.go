package params

type (
	Base64EncodingReq struct {
		Action  string `json:"action" validate:"oneof=encode decode"`
		Mode    string `json:"mode"  validate:"oneof='std' 'raw-std' 'raw-url' 'url'"`
		Payload string `json:"payload"  validate:"required"`
	}
	Base64EncodingResp struct {
		Result string `json:"result"`
	}

	JsonPrettyReq struct {
		Indent  string `json:"indent"`
		Payload string `json:"payload" validate:"required"`
	}
	JsonPrettyResp = Base64EncodingResp

	UrlReq struct {
		Action  string `json:"action" validate:"oneof=encode decode"`
		Payload string `json:"payload" validate:"required"`
	}
	UrlResp = Base64EncodingResp
)
