package msgs

import (
	cschainibc "gitlab.bianjie.ai/cschain/cschain/modules/ibc/core/types"
)

const (
	//ibc client
	MsgTypeCreateClient = "create_client"
	MsgTypeUpdateClient = "update_client"
	MsgTypeRecvPacket   = "recv_packet"
)

type (
	//ibc
	MsgRecvPacket   = cschainibc.MsgRecvPacket
	MsgCreateClient = cschainibc.MsgCreateClient
	MsgUpdateClient = cschainibc.MsgUpdateClient
)
