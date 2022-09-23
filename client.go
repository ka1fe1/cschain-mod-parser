package cschain_mod_parser

import (
	. "github.com/kaifei-bianjie/common-parser"
	"github.com/kaifei-bianjie/cschain-mod-parser/codec"
	"github.com/kaifei-bianjie/cschain-mod-parser/modules/ibc"
)

type MsgClient struct {
	Ibc Client
}

func NewMsgClient() MsgClient {
	codec.MakeEncodingConfig()
	return MsgClient{
		Ibc: ibc.NewClient(),
	}
}
