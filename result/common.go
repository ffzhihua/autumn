package result

//成功
type Success struct {

}

//失败
type Fail struct {

}

type Login struct {
	Token string    `json:"token"`
}
