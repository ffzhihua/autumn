## TITLE
autumn go framework

## build
`make`

## clean
`make clean`

##初次安装
go get -u github.com/golang/dep/cmd/dep
若$GOPATH/bin不在PATH下，则需要将生成的dep文件从$GOPATH/bin移动至$GOBIAN下

##验证
dep

##初始化
dep init -v

##优先从$GOPATH初始化
dep init -gopath -v

dep ensure -update

