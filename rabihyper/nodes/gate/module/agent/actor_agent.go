package agent

import (
	cstring "github.com/cherry-game/cherry/extend/string"
	cfacade "github.com/cherry-game/cherry/facade"
	clog "github.com/cherry-game/cherry/logger"
	cactor "github.com/cherry-game/cherry/net/actor"
	"github.com/cherry-game/cherry/net/parser/pomelo"
	cproto "github.com/cherry-game/cherry/net/proto"
	"github.com/rabihyper/internal/pd/Greeter"
	sessionKey "github.com/rabihyper/internal/session_key"
)

type (
	ActorAgent struct {
		cactor.Base
	}
)

func (p *ActorAgent) OnInit() {
	p.Local().Register("hello", p.hello)
}

func (p *ActorAgent) hello(session *cproto.Session, req *Greeter.HelloRequest) {
	agent, found := pomelo.GetAgent(p.ActorID())
	if !found {
		return
	}

	response := &Greeter.HelloResponse{
		ReplyMessage: req.Name,
	}

	agent.Response(session, response)
}

// OnSessionClose  当agent断开时，关闭对应的ActorAgent
func (p *ActorAgent) OnSessionClose(agent *pomelo.Agent) {
	session := agent.Session()
	serverId := session.GetString(sessionKey.ServerID)
	if serverId == "" {
		return
	}

	// 通知game节点关闭session
	childId := cstring.ToString(session.Uid)
	if childId != "" {
		targetPath := cfacade.NewChildPath(serverId, "player", childId)
		p.Call(targetPath, "sessionClose", nil)
	}

	// 自己退出
	p.Exit()
	clog.Infof("sessionClose path = %s", p.Path())
}
