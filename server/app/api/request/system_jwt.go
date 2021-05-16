package request

type CustomClaims struct {
	AdminId          uint    `gconv:"user_id"`
	AdminUuid        string  `gconv:"user_uuid"`
	AdminNickname    string  `gconv:"user_nickname"`
	Exp              float64 `gconv:"exp"`
	OrigIat          float64 `gconv:"orig_iat"`
	AdminAuthorityId string  `gconv:"user_authority_id"`
}
