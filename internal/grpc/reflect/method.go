package reflect

import "google.golang.org/protobuf/runtime/protoimpl"

type methodReflection struct {
	name      string
	shortName string
	options   optionsMap
}

func (m methodReflection) AddExtension(ext *protoimpl.ExtensionInfo) {
	m.extInfo = append(m.extInfo, ext)
}
