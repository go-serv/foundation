package server

type policyKind int

const (
	LoosePolicy policyKind = iota + 1
	StrictPolicy
)

type ipMask string

type permission struct {
	methodName  string
	clientRoles []string
	// If true rule will be saved into a persistent storage.
	persist bool
}

type policy struct {
	kind        policyKind   `json:"kind"`
	methodPerms []permission `json:"permissions"`
	whitelist   []ipMask     `json:"whitelist"`
	blacklist   []ipMask     `json:"blacklist"`
}
