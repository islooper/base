package util

// rcp 响应接口
type RpcResponder interface {
	SetCode(ret int32)
	SetMsg(msg string)
}

type RpcError interface {
	GetCode() int32
	GetMsg() string
}

// 定义错误
type ErrInfo struct {
	Code int32  // 错误码
	Msg  string // 展示给用户看的
	Err  error  // 保存内部错误信息
}

func (e *ErrInfo) GetCode() int32 {
	return e.Code
}

func (e *ErrInfo) GetMsg() string {
	return e.Msg
}

func (e *ErrInfo) Is(info *ErrInfo) bool {
	if info == nil {
		return false
	}
	if e.Code == info.Code {
		return true
	} else {
		return false
	}
}

func (e *ErrInfo) IsErrNot() bool {
	return e.Is(ErrNot)
}

//ret=0 成功。
//noinspection ALL
var (
	ErrNot    = &ErrInfo{0, "success", nil}
	ErrUnknow = &ErrInfo{1, "unknown error", nil}

	ErrParam    = &ErrInfo{2, "param error", nil}
	ErrDataBase = &ErrInfo{3, "database error", nil}
	ErrToken    = &ErrInfo{4, "token error", nil}
	ErrRpc      = &ErrInfo{5, "rpc call error", nil}
)

// 从 info 创建一个新的 ErrInfo 类型的对象。
// 当 msg 不为空,则用 msg 替换原 msg
// 当 err 不为 nil,则用 err 替换 原 err
//noinspection ALL
func NewErrInfo(info *ErrInfo, msg string, err error) *ErrInfo {
	errInfo := &ErrInfo{
		Code: info.Code,
		Msg:  info.Msg,
		Err:  info.Err,
	}
	if msg != "" {
		errInfo.Msg = msg
	}

	if err != nil {
		errInfo.Err = err
	}
	return errInfo
}

func WriteRpcRsp(rspPtr interface{}, rpcError RpcError, data map[string]interface{}) {
	if nil == data {
		data = make(map[string]interface{})
	}
	if rpcError == nil {
		panic("rpcError is nil")
	}
	data["Code"] = rpcError.GetCode()
	data["Msg"] = rpcError.GetMsg()
	SetStructVals(rspPtr, data)
}

//noinspection ALL
func WriteRpcRspWithMsg(rspPtr interface{}, rpcError RpcError, msg string, data map[string]interface{}) {
	if nil == data {
		data = make(map[string]interface{})
	}
	if rpcError == nil {
		panic("rpcError is nil")
	}
	data["Code"] = rpcError.GetCode()
	data["Msg"] = msg
	SetStructVals(rspPtr, data)
}
