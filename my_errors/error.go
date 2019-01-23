package my_errors

type Error interface { //在 error 中划分出 Error
	Error() string //继承 error 接口
	ErrorCode() int
	GetOrigin() error //返回原始错误
}

type MyError struct { //实现了 Error 的错误1000
	msg    string //附加信息
	origin error
}

func New(msg string, origin error) Error {
	return MyError{msg, origin}

}

func (this MyError) Error() string {
	if len(this.msg) == 0 {
		this.msg = "未知错误"
	}

	return this.msg
}

func (this MyError) ErrorCode() int {
	return 101
}

func (this MyError) GetOrigin() error {
	return this.origin
}

//实现了Error 的错误1002
type Error404 struct {
	MyError
}

func (err Error404) Error() string { //重写了该方法
	return "路由错误"
}

func (err Error404) ErrorCode() int {
	return 404
}

//实现了Error 的错误202
type NotLoginError struct {
	MyError
}

func (err NotLoginError) Error() string { //重写方法
	return "请先登陆"
}

func (err NotLoginError) ErrorCode() int {
	return 202
}

//实现了Error 的错误303
type UnactivatedError struct {
	MyError
}

func (err UnactivatedError) Error() string { //重写方法
	return "您的帐户尚未激活"
}

func (err UnactivatedError) ErrorCode() int {
	return 303
}
