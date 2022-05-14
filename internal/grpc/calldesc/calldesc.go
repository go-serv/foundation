//
//

package calldesc

import "context"

type callDesc struct {
	context.Context
}

type callDescServer struct {
	callDesc
}

type callDescClient struct {
	callDesc
}
