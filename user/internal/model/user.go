package user

type UserProfile struct {
	UID      int64  `json:"uid"`
	NickName string `json:"nick_name"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Age      string `json:"age"`
	Birthday string `json:"birthday"`
}
