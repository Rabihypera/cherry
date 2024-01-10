package gate

import (
	cfacade "github.com/cherry-game/cherry/facade"
	"github.com/cherry-game/cherry/net/parser/pomelo"
	pmessage "github.com/cherry-game/cherry/net/parser/pomelo/message"
	cproto "github.com/cherry-game/cherry/net/proto"
	"github.com/rabihyper/internal/code"
	pb "github.com/rabihyper/internal/pd"
)

var (
	// 客户端连接后，必需先执行第一条协议，进行token验证后，才能进行后续的逻辑
	firstRouteName = "gate.rabihyper.hello"

	notLoginRsp = &pb.Int32{
		Value: code.PlayerDenyLogin,
	}
)

func onPomeloDataRoute(agent *pomelo.Agent, route *pmessage.Route, msg *pmessage.Message) {
	session := pomelo.BuildSession(agent, msg)

	if !session.IsBind() && msg.Route != firstRouteName {
		agent.Kick(notLoginRsp, true)
		return
	}

	if agent.NodeType() == route.NodeType() {
		targetPath := cfacade.NewChildPath(agent.NodeId(), route.HandleName(), session.Sid)
		pomelo.LocalDataRoute(agent, session, route, msg, targetPath)
	} else {
		gameNodeRoute(agent, session, route, msg)
	}
}

// gameNodeRoute 实现agent路由消息到游戏节点
func gameNodeRoute(agent *pomelo.Agent, session *cproto.Session, route *pmessage.Route, msg *pmessage.Message) {
	if !session.IsBind() {
		return
	}

}
