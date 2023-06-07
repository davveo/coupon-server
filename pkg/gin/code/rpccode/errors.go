package rpccode

import (
	_errors "github.com/davveo/market-coupon/pkg/pkgerror"
)

var errors Errors

type Errors interface {
	Render(code int) ErrInfo
	Register(code int, fields ...string)
	ParseErrInfo(err error) (int, ErrInfo)
}

type set struct {
	list map[int]ErrInfo
}

func (s *set) Render(code int) ErrInfo {
	return s.list[code]
}

func newErrorSet() Errors {
	s := &set{
		list: make(map[int]ErrInfo),
	}
	return s
}

// Register(300001, "msg", "title", "reference")
// 注册 错误码
func (s *set) Register(code int, fields ...string) {
	fieldsSz := len(fields)
	ei := ErrInfo{}
	switch {
	case fieldsSz > 2:
		ei.Reference = fields[2]
		fallthrough
	case fieldsSz > 1:
		ei.Title = fields[1]
		fallthrough
	case fieldsSz > 0:
		ei.Msg = fields[0]
	}
	s.list[code] = ei
	register(code, ei.Msg, ei.Reference)
}

func init() {
	errors = newErrorSet()
}

func (s *set) ParseErrInfo(err error) (int, ErrInfo) {
	coder := _errors.ParseCoder(err)
	errInfo := Render(coder.Code())
	if errInfo.Msg == "" {
		return coder.Code(), ErrInfo{
			Title: "Uncaught exception",
			Msg:   err.Error(),
		}
	}
	return coder.Code(), Render(coder.Code())
}
