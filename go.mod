module github.com/go-serv/foundation

go 1.18

replace github.com/AgentCoop/go-work => /home/pihpah/go/src/github.com/AgentCoop/go-work

replace google.golang.org/grpc => ../grpc-go

require (
	github.com/AgentCoop/go-work v0.0.2
	github.com/fatih/color v1.13.0
	github.com/monnand/dhkx v0.0.0-20180522003156-9e5b033f1ac4
	golang.org/x/sys v0.0.0-20210809222454-d867a43fc93e
	google.golang.org/grpc v1.46.0
	google.golang.org/protobuf v1.27.1
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/mattn/go-colorable v0.1.9 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	golang.org/x/net v0.0.0-20210805182204-aaa1db679c0d // indirect
	golang.org/x/text v0.3.6 // indirect
	google.golang.org/genproto v0.0.0-20210126160654-44e461bb6506 // indirect
)
