package gate

import (
	"github.com/cherry-game/cherry"
	cherryGops "github.com/cherry-game/cherry/components/gops"
	cfacade "github.com/cherry-game/cherry/facade"
	cconnector "github.com/cherry-game/cherry/net/connector"
	"github.com/cherry-game/cherry/net/parser/pomelo"
	"github.com/rabihyper/nodes/gate/module/agent"
)

func Run(profileFilePath, nodeId string) {
	app := cherry.Configure(profileFilePath, nodeId, true, cherry.Cluster)

	netParser := buildPomeloParser(app)
	app.SetNetParser(netParser)

	app.Register(cherryGops.New())

	app.Startup()
}

func buildPomeloParser(app *cherry.AppBuilder) cfacade.INetParser {
	// 使用pomelo网络数据包解析器
	agentActor := pomelo.NewActor("rabihyper")
	//创建一个tcp监听，用于client/robot压测机器人连接网关tcp
	agentActor.AddConnector(cconnector.NewTCP(":10011"))
	//再创建一个websocket监听，用于h5客户端建立连接
	agentActor.AddConnector(cconnector.NewWS(app.Address()))
	//当有新连接创建Agent时，启动一个自定义(ActorAgent)的子actor
	agentActor.SetOnNewAgent(func(newAgent *pomelo.Agent) {
		childActor := &agent.ActorAgent{}
		newAgent.AddOnClose(childActor.OnSessionClose)
		agentActor.Child().Create(newAgent.SID(), childActor)
	})

	// 设置数据路由函数
	agentActor.SetOnDataRoute(onPomeloDataRoute)

	return agentActor
}
