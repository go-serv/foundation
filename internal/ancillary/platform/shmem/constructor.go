package shmem

func NewForRead(objname string, len uint32, cap uint32) *blockInfo {
	b := new(blockInfo)
	b.objname = objname
	b.len = len
	b.cap = cap
	return b
}
