package typeconv

import "strconv"

/**
	//int 转化为string
	s := strconv.Itoa(i)

	//强制转化为int64后使用FormatInt
	s := strconv.FormatInt(int64(i), 10)

	//string 转为int
	i, err := strconv.Atoi(s)

	//int64 转 string，第二个参数为基数
	s := strconv.FormatInt(i64, 10)

	// string 转换为 int64
	//第二参数为基数，后面为位数，可以转换为int32，int64等
	i64, err := strconv.ParseInt(s, 10, 64)


	// flaot 转为string 最后一位是位数设置float32或float64
	s1 := strconv.FormatFloat(v, 'E', -1, 32)
	//string 转 float 同样最后一位设置位数
	v, err := strconv.ParseFloat(s, 32)
	v, err := strconv.atof32(s)


	b, err := strconv.ParseBool("true") // string 转bool
	s := strconv.FormatBool(true) // bool 转string

	var a interface{}
	var b string
	a = "123"
	b = a.(string)
 */
func main() {




}

func IntToString(i int) string{

	s := strconv.Itoa(i)
	return s
}

func Int64ToString(i int) string  {
	s := strconv.FormatInt(int64(i), 10)
	return s
}