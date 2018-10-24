package validate

import "regexp"

const (
	MOBILE_REGEX            = "^(13[0-9]|14[0-9]|15[0-9]|17[0-9]|18[0-9])\\d{8}$"
	IDCARD_REGEX            = "^(\\d{15}$|^\\d{18}$|^\\d{17}(\\d|X|x))$"
	ETHER_WALLET_ADDR_REGEX = "^0x[0-9a-fA-F]{40}$"
	EMAIL_REGEX             = "^([a-z0-9_\\.-]+)@([\\da-z\\.-]+)\\.([a-z\\.]{2,6})$"
)

func IsEmail(email string) bool  {
	reg := regexp.MustCompile(EMAIL_REGEX)
	return reg.MatchString(email)
}

//检测以太坊钱包地址
func IsEtherWalletAddr(addr string) bool {
	reg := regexp.MustCompile(ETHER_WALLET_ADDR_REGEX)
	return reg.MatchString(addr)
}

//检测手机号
func IsMobile(mobile string) bool {
	reg := regexp.MustCompile(MOBILE_REGEX)
	return reg.MatchString(mobile)
}

//检测身份证号
func IsIdCard(idNo string) bool {
	reg := regexp.MustCompile(IDCARD_REGEX)
	return reg.MatchString(idNo)
}

func IsFloat(f string, n string) bool {
	float_regex	:= "^[0-9]{1,8}(.[0-9]{"+n+"})?$"
	reg := regexp.MustCompile(float_regex)
	return reg.MatchString(f)
}
