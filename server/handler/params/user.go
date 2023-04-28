package params

type (
	LoginReq struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
	LoginResp struct {
		AccessToken string `json:"access_token"`
		ExpireIn    int64  `json:"expire_in"`
	}
	RefreshAccessTokenReq struct {
		AccessToken string `json:"accessToken" validate:"required"`
	}
	RefreshAccessTokenResp struct {
		AccessToken string `json:"access_token"`
		ExpireIn    int    `json:"expire_in"`
	}
)
