module github.com/rabihyper

go 1.21

require (
	github.com/cherry-game/cherry v1.3.12
	github.com/cherry-game/cherry/components/gops v1.3.12
	github.com/urfave/cli/v2 v2.27.1
	google.golang.org/protobuf v1.32.0
)

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/google/gops v0.3.28 // indirect
	github.com/gorilla/websocket v1.5.1 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.17.4 // indirect
	github.com/lestrrat-go/strftime v1.0.6 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/nats-io/nats.go v1.31.0 // indirect
	github.com/nats-io/nkeys v0.4.6 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/xrash/smetrics v0.0.0-20231213231151-1d8dd44e695e // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.26.0 // indirect
	golang.org/x/crypto v0.17.0 // indirect
	golang.org/x/net v0.19.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
)

replace (
	github.com/cherry-game/cherry => ../
	github.com/cherry-game/cherry/components/cron => ../components/cron
	github.com/cherry-game/cherry/components/data-config => ../components/data-config
	github.com/cherry-game/cherry/components/gin => ../components/gin
	github.com/cherry-game/cherry/components/gops => ../components/gops
	github.com/cherry-game/cherry/components/gorm => ../components/gorm
)
