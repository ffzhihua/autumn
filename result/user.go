package result

type UserInfo struct {
	Uid         int     `json:"uid"`
	Username    string  `json:"username"`
	LastLogin   int64   `json:"last_login"`
	MobileValid int     `json:"mobile_valid"`
	EmailValid  int     `json:"email_valid"`
	NameValid   int     `json:"name_valid"`
	GaValid     int     `json:"ga_valid"`
	RegTime     int64   `json:"reg_time"`
}


type ForgotVerify struct {
	Token string    `json:"token"`
}

type MsgList struct {
	MsgId       int     `json:"msg_id"`
	Title       string  `json:"title"`
	Content     string  `json:"content"`
	Status      int     `json:"status"`
	Created     int64   `json:"created_at"`
}