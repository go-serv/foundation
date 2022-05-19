package ancillary

import "strings"

type GrpcDotNotation string

func (dot GrpcDotNotation) MethodName() string {
	s := string(dot)
	if strings.HasPrefix(s, "/") { /* /service/name notation */
		parts := strings.Split(s, "/")
		dotName := parts[1] + "." + parts[2]
		return dotName
	} else {
		return string(dot)
	}
}
