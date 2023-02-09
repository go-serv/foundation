package reflect

import (
	"github.com/mesh-master/foundation/pkg/z"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
)

type service struct {
	desc      protoreflect.ServiceDescriptor
	methods   []z.MethodReflectionInterface
	extValues extValueMap
}

func (s *service) FullName() string {
	return string(s.Descriptor().FullName())
}

func (s *service) Descriptor() protoreflect.ServiceDescriptor {
	return s.desc
}

func (s *service) Methods() []z.MethodReflectionInterface {
	return s.methods
}

func (s *service) Get(key *protoimpl.ExtensionInfo) (interface{}, bool) {
	return s.extValues.get(key)
}

func (s *service) Has(key *protoimpl.ExtensionInfo) bool {
	return s.extValues.has(key)
}

func (s *service) Bool(key *protoimpl.ExtensionInfo) bool {
	return s.extValues.bool(key)
}

func (s *service) Shadow(svcExt *protoimpl.ExtensionInfo, mExt *protoimpl.ExtensionInfo, mRef z.MethodReflectionInterface) (v interface{}, has bool) {
	if mRef.Has(mExt) {
		return mRef.Get(mExt)
	} else {
		return s.Get(svcExt)
	}
}
