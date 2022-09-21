package service

import (
	"github.com/go-serv/foundation/internal/service"
	"google.golang.org/protobuf/runtime/protoimpl"
)

func AddProtoExtension(ext *protoimpl.ExtensionInfo) {
	service.Reflection().AddProtoExtension(ext)
}
