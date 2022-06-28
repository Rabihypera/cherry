package cherryHandler

import (
	"context"
	cfacade "github.com/cherry-game/cherry/facade"
	clog "github.com/cherry-game/cherry/logger"
	cmsg "github.com/cherry-game/cherry/net/message"
	cproto "github.com/cherry-game/cherry/net/proto"
	csession "github.com/cherry-game/cherry/net/session"
	"reflect"
	"runtime/debug"
)

type (
	ExecutorLocal struct {
		cfacade.IApplication
		groupIndex    int
		Session       *csession.Session
		Msg           *cmsg.Message
		HandlerFn     *cfacade.HandlerFn
		Ctx           context.Context
		BeforeFilters []FilterFn
		AfterFilters  []FilterFn
		PrintLog      bool
	}
)

func NewExecutorLocal(session *csession.Session, msg *cmsg.Message) ExecutorLocal {
	return ExecutorLocal{
		Session: session,
		Msg:     msg,
	}
}

func (p *ExecutorLocal) Index() int {
	return p.groupIndex
}

func (p *ExecutorLocal) SetIndex(index int) {
	p.groupIndex = index
}

func (p *ExecutorLocal) Invoke() {
	defer func() {
		if rev := recover(); rev != nil {
			clog.Warnf("recover in Local. %s", string(debug.Stack()))
			clog.Warnf("msg = [%+v]", p.Msg)
		}
	}()

	for _, filter := range p.BeforeFilters {
		if filter(p.Ctx, p.Session, p.Msg) == false {
			return
		}
	}

	argsLen := len(p.HandlerFn.InArgs)
	if argsLen < 2 || argsLen > 3 {
		clog.Warnf("[Route = %v] method in args error.", p.Msg.Route)
		clog.Warnf("func(session,request) or func(ctx,session,request)")
		return
	}

	var params []reflect.Value
	var ret []reflect.Value

	if p.HandlerFn.IsRaw {
		params = make([]reflect.Value, argsLen)
		params[0] = reflect.ValueOf(p.Ctx)
		params[1] = reflect.ValueOf(p.Session)
		params[2] = reflect.ValueOf(p.Msg)

		ret = p.HandlerFn.Value.Call(params)
		if p.PrintLog {
			p.Session.Debugf("[local] [groupIndex = %d, route = %s, mid = %d, req = %+v, rsp = %+v]",
				p.groupIndex,
				p.Msg.Route,
				p.Msg.ID,
				p.Msg.Data,
				printRet(ret),
			)
		}
	} else {
		val, err := p.unmarshalData(argsLen - 1)
		if err != nil {
			p.Session.Errorf("err = %v, msg = %v", err, p.Msg)
			return
		}

		if argsLen == 2 {
			params = make([]reflect.Value, argsLen)
			params[0] = reflect.ValueOf(p.Session)
			params[1] = reflect.ValueOf(val)
		} else {
			params = make([]reflect.Value, argsLen)
			params[0] = reflect.ValueOf(p.Ctx)
			params[1] = reflect.ValueOf(p.Session)
			params[2] = reflect.ValueOf(val)
		}

		ret = p.HandlerFn.Value.Call(params)
		if p.PrintLog {
			p.Session.Debugf("[local] [groupIndex = %d, route = %s, mid = %d, req = %+v, rsp = %+v]",
				p.groupIndex,
				p.Msg.Route,
				p.Msg.ID,
				val,
				printRet(ret),
			)
		}
	}

	if p.Msg.Type == cmsg.Request {
		retLen := len(ret)

		if retLen == 2 {
			if ret[0].IsNil() {
				if face := ret[1].Interface(); face != nil {
					if code, ok := face.(int32); ok {
						rsp := &cproto.Response{
							Code: code,
						}

						p.Session.ResponseMID(p.Msg.ID, rsp, true)
					} else {
						p.Session.Warn(face)
					}
				} else {
					p.Session.Warnf("ret value type error. [type = %+v]", ret)
				}
			} else {
				p.Session.ResponseMID(p.Msg.ID, ret[0].Interface())
			}
		}
	}

	for _, filter := range p.AfterFilters {
		if !filter(p.Ctx, p.Session, p.Msg) {
			break
		}
	}
}

func (p *ExecutorLocal) unmarshalData(index int) (interface{}, error) {
	in2 := p.HandlerFn.InArgs[index]

	var val interface{}
	val = reflect.New(in2.Elem()).Interface()

	err := p.Unmarshal(p.Msg.Data, val)
	if err != nil {
		return nil, err
	}

	return val, err
}

func (p *ExecutorLocal) String() string {
	return p.Msg.Route
}
