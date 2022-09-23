package codec

import (
	"github.com/cosmos/cosmos-sdk/codec"
	ctypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/simapp/params"
	"github.com/cosmos/cosmos-sdk/std"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	common "github.com/kaifei-bianjie/common-parser/codec"
	recordtypes "gitlab.bianjie.ai/cschain/cschain/modules/ibc/applications/record/types"
	bcos "gitlab.bianjie.ai/cschain/cschain/modules/ibc/light-clients/bcos/types"
	brochain "gitlab.bianjie.ai/cschain/cschain/modules/ibc/light-clients/brochain/types"
	fabric "gitlab.bianjie.ai/cschain/cschain/modules/ibc/light-clients/fabric/types"
	tendermint "gitlab.bianjie.ai/cschain/cschain/modules/ibc/light-clients/tendermint/types"
	wutong "gitlab.bianjie.ai/cschain/cschain/modules/ibc/light-clients/wutong/types"
)

// 初始化账户地址前缀
func MakeEncodingConfig() {
	amino := codec.NewLegacyAmino()
	interfaceRegistry := ctypes.NewInterfaceRegistry()
	brochain.RegisterInterfaces(interfaceRegistry)
	fabric.RegisterInterfaces(interfaceRegistry)
	tendermint.RegisterInterfaces(interfaceRegistry)
	wutong.RegisterInterfaces(interfaceRegistry)
	bcos.RegisterInterfaces(interfaceRegistry)
	recordtypes.RegisterInterfaces(interfaceRegistry)
	moduleBasics := module.NewBasicManager(common.AppModules...)
	moduleBasics.RegisterInterfaces(interfaceRegistry)
	moduleBasics.RegisterLegacyAminoCodec(amino)
	std.RegisterInterfaces(interfaceRegistry)
	std.RegisterLegacyAminoCodec(amino)
	sdk.RegisterInterfaces(interfaceRegistry)
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	txCfg := tx.NewTxConfig(marshaler, tx.DefaultSignModes)
	common.Encodecfg = params.EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig:          txCfg,
		Amino:             amino,
	}
}
