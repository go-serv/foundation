package descriptor

import "google.golang.org/protobuf/runtime/protoimpl"

type ServiceDescriptorInterface interface {
	Get(key *protoimpl.ExtensionInfo) (interface{}, bool)
	AddServiceProtoExt(ext *protoimpl.ExtensionInfo)
	AddMethodProtoExt(ext *protoimpl.ExtensionInfo)
	Populate()
	FindMethodDescriptorByName(fullName string) *methodDescriptor
}

type MethodDescriptorInterface interface {
	Get(key *protoimpl.ExtensionInfo) (interface{}, bool)
}

type MessageDescriptorInterface interface {
	MethodName() string
}
