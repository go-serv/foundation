package net

import (
	i "github.com/go-serv/service/internal"
	"github.com/go-serv/service/internal/ancillary"
	cc "github.com/go-serv/service/internal/grpc/codec"
)

type headerFlags i.HeaderFlags32Type

const (
	Encryption headerFlags = 1 << iota
)

type codec struct {
	i.CodecInterface
}

type dataframe struct {
	i.DataFrameInterface
	shmObjName   string
	shmBlockSize uint32
}

func (c *codec) NewDataFrame() i.DataFrameInterface {
	df := new(dataframe)
	df.DataFrameInterface = cc.NewDataFrame()
	return df
}

func (df *dataframe) Parse(b []byte, hook func(*ancillary.NetReader) error) error {
	return df.DataFrameInterface.Parse(b, func(netr *ancillary.NetReader) (err error) {
		df.shmObjName, err = netr.ReadString()
		if err != nil {
			return
		}
		df.shmBlockSize, err = ancillary.GenericNetReader[uint32](netr)
		if err != nil {
			return
		}
		return
	})
}

func (df *dataframe) Compose([]byte) (out []byte, err error) {
	netw := ancillary.NewNetWriter()
	netw.WriteString(df.shmObjName)
	err = ancillary.GenericNetWriter[uint32](netw, df.shmBlockSize)
	if err != nil {
		return
	}
	header := netw.Bytes()
	out, err = df.DataFrameInterface.Compose(header)
	return
}

func (df *dataframe) ParseHook(netr *ancillary.NetReader) (err error) {
	df.shmObjName, err = netr.ReadString()
	if err != nil {
		return
	}
	df.shmBlockSize, err = ancillary.GenericNetReader[uint32](netr)
	if err != nil {
		return
	}
	return
}

func (df *dataframe) SharedMemObjectName() string {
	return df.shmObjName
}

func (df *dataframe) WithSharedMemObjectName(name string) {
	df.shmObjName = name
}

func (df *dataframe) SharedMemBlockSize() uint32 {
	return df.shmBlockSize
}

func (df *dataframe) WithSharedMemBlockSize(size uint32) {
	df.shmBlockSize = size
}
