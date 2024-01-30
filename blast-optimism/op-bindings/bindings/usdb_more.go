// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const USDBStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L2/USDB.sol:USDB\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"src/L2/USDB.sol:USDB\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"src/L2/USDB.sol:USDB\",\"label\":\"_hashedName\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_bytes32\"},{\"astId\":1003,\"contract\":\"src/L2/USDB.sol:USDB\",\"label\":\"_hashedVersion\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_bytes32\"},{\"astId\":1004,\"contract\":\"src/L2/USDB.sol:USDB\",\"label\":\"_name\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_string_storage\"},{\"astId\":1005,\"contract\":\"src/L2/USDB.sol:USDB\",\"label\":\"_version\",\"offset\":0,\"slot\":\"4\",\"type\":\"t_string_storage\"},{\"astId\":1006,\"contract\":\"src/L2/USDB.sol:USDB\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"5\",\"type\":\"t_array(t_uint256)48_storage\"},{\"astId\":1007,\"contract\":\"src/L2/USDB.sol:USDB\",\"label\":\"_nonces\",\"offset\":0,\"slot\":\"53\",\"type\":\"t_mapping(t_address,t_struct(Counter)1024_storage)\"},{\"astId\":1008,\"contract\":\"src/L2/USDB.sol:USDB\",\"label\":\"_PERMIT_TYPEHASH_DEPRECATED_SLOT\",\"offset\":0,\"slot\":\"54\",\"type\":\"t_bytes32\"},{\"astId\":1009,\"contract\":\"src/L2/USDB.sol:USDB\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"55\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":1010,\"contract\":\"src/L2/USDB.sol:USDB\",\"label\":\"price\",\"offset\":0,\"slot\":\"104\",\"type\":\"t_uint256\"},{\"astId\":1011,\"contract\":\"src/L2/USDB.sol:USDB\",\"label\":\"pending\",\"offset\":0,\"slot\":\"105\",\"type\":\"t_uint256\"},{\"astId\":1012,\"contract\":\"src/L2/USDB.sol:USDB\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"106\",\"type\":\"t_array(t_uint256)48_storage\"},{\"astId\":1013,\"contract\":\"src/L2/USDB.sol:USDB\",\"label\":\"name\",\"offset\":0,\"slot\":\"154\",\"type\":\"t_string_storage\"},{\"astId\":1014,\"contract\":\"src/L2/USDB.sol:USDB\",\"label\":\"symbol\",\"offset\":0,\"slot\":\"155\",\"type\":\"t_string_storage\"},{\"astId\":1015,\"contract\":\"src/L2/USDB.sol:USDB\",\"label\":\"_shares\",\"offset\":0,\"slot\":\"156\",\"type\":\"t_mapping(t_address,t_uint256)\"},{\"astId\":1016,\"contract\":\"src/L2/USDB.sol:USDB\",\"label\":\"_totalShares\",\"offset\":0,\"slot\":\"157\",\"type\":\"t_uint256\"},{\"astId\":1017,\"contract\":\"src/L2/USDB.sol:USDB\",\"label\":\"_remainders\",\"offset\":0,\"slot\":\"158\",\"type\":\"t_mapping(t_address,t_uint256)\"},{\"astId\":1018,\"contract\":\"src/L2/USDB.sol:USDB\",\"label\":\"_fixed\",\"offset\":0,\"slot\":\"159\",\"type\":\"t_mapping(t_address,t_uint256)\"},{\"astId\":1019,\"contract\":\"src/L2/USDB.sol:USDB\",\"label\":\"_totalVoidAndRemainders\",\"offset\":0,\"slot\":\"160\",\"type\":\"t_uint256\"},{\"astId\":1020,\"contract\":\"src/L2/USDB.sol:USDB\",\"label\":\"_yieldMode\",\"offset\":0,\"slot\":\"161\",\"type\":\"t_mapping(t_address,t_enum(YieldMode)1023)\"},{\"astId\":1021,\"contract\":\"src/L2/USDB.sol:USDB\",\"label\":\"_allowances\",\"offset\":0,\"slot\":\"162\",\"type\":\"t_mapping(t_address,t_mapping(t_address,t_uint256))\"},{\"astId\":1022,\"contract\":\"src/L2/USDB.sol:USDB\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"163\",\"type\":\"t_array(t_uint256)41_storage\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)41_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[41]\",\"numberOfBytes\":\"1312\",\"base\":\"t_uint256\"},\"t_array(t_uint256)48_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[48]\",\"numberOfBytes\":\"1536\",\"base\":\"t_uint256\"},\"t_array(t_uint256)49_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\",\"base\":\"t_uint256\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_enum(YieldMode)1023\":{\"encoding\":\"inplace\",\"label\":\"enum YieldMode\",\"numberOfBytes\":\"1\"},\"t_mapping(t_address,t_enum(YieldMode)1023)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e enum YieldMode)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_enum(YieldMode)1023\"},\"t_mapping(t_address,t_mapping(t_address,t_uint256))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(address =\u003e uint256))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_address,t_uint256)\"},\"t_mapping(t_address,t_struct(Counter)1024_storage)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e struct CountersUpgradeable.Counter)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_struct(Counter)1024_storage\"},\"t_mapping(t_address,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_uint256\"},\"t_string_storage\":{\"encoding\":\"bytes\",\"label\":\"string\",\"numberOfBytes\":\"32\"},\"t_struct(Counter)1024_storage\":{\"encoding\":\"inplace\",\"label\":\"struct CountersUpgradeable.Counter\",\"numberOfBytes\":\"32\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var USDBStorageLayout = new(solc.StorageLayout)

var USDBDeployedBin = "0x608060405234801561001057600080fd5b506004361061020b5760003560e01c80637ae556b51161012a578063aad3ec96116100bd578063dd62ed3e1161008c578063e20ccec311610071578063e20ccec31461051e578063e78cea9214610527578063ee9a31a21461054d57600080fd5b8063dd62ed3e146104c5578063e12f3a611461050b57600080fd5b8063aad3ec9614610433578063c44b11f714610446578063d505accf1461048c578063d6c0b2c41461049f57600080fd5b806395d89b41116100f957806395d89b41146103fc5780639dc29fac14610404578063a035b1fe14610417578063a9059cbb1461042057600080fd5b80637ae556b51461039f5780637ecebe00146103c65780638129fc1c146103d957806384b0196e146103e157600080fd5b806330adf81f116101a25780634291cd11116101715780634291cd111461036957806354fd4d50146103715780635b9af12b1461037957806370a082311461038c57600080fd5b806330adf81f146102ec578063313ce567146103135780633644e5151461034c57806340c10f191461035457600080fd5b8063095ea7b3116101de578063095ea7b3146102ab57806318160ddd146102be5780631a33757d146102c657806323b872dd146102d957600080fd5b806301ffc9a714610210578063033964be1461023857806306661abd1461028457806306fdde0314610296575b600080fd5b61022361021e366004612672565b610574565b60405190151581526020015b60405180910390f35b61025f7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161022f565b609d545b60405190815260200161022f565b61029e610614565b60405161022f919061272a565b6102236102b9366004612766565b6106a2565b6102886106bc565b6102886102d4366004612790565b6106e0565b6102236102e73660046127b1565b610743565b6102887f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c981565b61033a7f000000000000000000000000000000000000000000000000000000000000000081565b60405160ff909116815260200161022f565b610288610765565b610367610362366004612766565b61076f565b005b610367610889565b61029e6108dd565b6103676103873660046127ed565b610980565b61028861039a366004612806565b61098c565b61025f7f000000000000000000000000000000000000000000000000000000000000000081565b6102886103d4366004612806565b610a3f565b610367610a6a565b6103e9610cf8565b60405161022f9796959493929190612821565b61029e610dd4565b610367610412366004612766565b610de1565b61028860685481565b61022361042e366004612766565b610eef565b610288610441366004612766565b610f05565b61047f610454366004612806565b73ffffffffffffffffffffffffffffffffffffffff16600090815260a1602052604090205460ff1690565b60405161022f9190612923565b61036761049a366004612931565b61113a565b7f000000000000000000000000000000000000000000000000000000000000000061025f565b6102886104d33660046129a4565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260a26020908152604080832093909416825291909152205490565b610288610519366004612806565b6112f9565b61028860695481565b7f000000000000000000000000000000000000000000000000000000000000000061025f565b61025f7f000000000000000000000000000000000000000000000000000000000000000081565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007fec4fc8e3000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000841682148061060c57507fffffffff00000000000000000000000000000000000000000000000000000000848116908216145b949350505050565b609a8054610621906129d7565b80601f016020809104026020016040519081016040528092919081815260200182805461064d906129d7565b801561069a5780601f1061066f5761010080835404028352916020019161069a565b820191906000526020600020905b81548152906001019060200180831161067d57829003601f168201915b505050505081565b6000336106b08185856113e3565b60019150505b92915050565b600060a054609d546068546106d19190612a53565b6106db9190612a90565b905090565b60006106ec33836114ec565b3373ffffffffffffffffffffffffffffffffffffffff167fcaa97ab28bae75adcb5a02786c64b44d0d3139aa521bf831cdfbe280ef246e36836040516107329190612923565b60405180910390a26106b63361098c565b600061075084338461167d565b61075b848484611724565b5060019392505050565b60006106db611831565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146107de576040517fea0e1ccb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff821661082b576040517fea553b3400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610835828261183b565b8173ffffffffffffffffffffffffffffffffffffffff167f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d41213968858260405161087d91815260200190565b60405180910390a25050565b6108916118be565b6108db57609d546069546040517f0f68f20e000000000000000000000000000000000000000000000000000000008152600481019290925260248201526044015b60405180910390fd5b565b60606109087f0000000000000000000000000000000000000000000000000000000000000000611958565b6109317f0000000000000000000000000000000000000000000000000000000000000000611958565b61095a7f0000000000000000000000000000000000000000000000000000000000000000611958565b60405160200161096c93929190612aa8565b604051602081830303815290604052905090565b61098981611a8d565b50565b73ffffffffffffffffffffffffffffffffffffffff8116600090815260a1602052604081205460ff16818160028111156109c8576109c86128e0565b03610a105773ffffffffffffffffffffffffffffffffffffffff83166000908152609c6020908152604080832054609e90925290912054610a099190611b5e565b9150610a39565b73ffffffffffffffffffffffffffffffffffffffff83166000908152609f602052604090205491505b50919050565b73ffffffffffffffffffffffffffffffffffffffff81166000908152603560205260408120546106b6565b600054610100900460ff1615808015610a8a5750600054600160ff909116105b80610aa45750303b158015610aa4575060005460ff166001145b610b30576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016108d2565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015610b8e57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b610c076040518060400160405280600c81526020017f5265626173696e672055534400000000000000000000000000000000000000008152506040518060400160405280600481526020017f5553444200000000000000000000000000000000000000000000000000000000815250633b9aca00611b79565b6040517f4c802f3800000000000000000000000000000000000000000000000000000000815273430000000000000000000000000000000000000290634c802f3890610c6190309060019060009061dead90600401612b1e565b600060405180830381600087803b158015610c7b57600080fd5b505af1158015610c8f573d6000803e3d6000fd5b50505050801561098957600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a150565b6000606080600080600060606001546000801b148015610d185750600254155b610d7e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f4549503731323a20556e696e697469616c697a6564000000000000000000000060448201526064016108d2565b610d86611c3b565b610d8e611ccd565b604080516000808252602082019092527f0f000000000000000000000000000000000000000000000000000000000000009b939a50919850469750309650945092509050565b609b8054610621906129d7565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610e50576040517fea0e1ccb00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216610e9d576040517f160fca8a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610ea78282611cdc565b8173ffffffffffffffffffffffffffffffffffffffff167fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca58260405161087d91815260200190565b6000610efc338484611724565b50600192915050565b60003373ffffffffffffffffffffffffffffffffffffffff8416610f55576040517f446b606800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600273ffffffffffffffffffffffffffffffffffffffff8216600090815260a1602052604090205460ff166002811115610f9157610f916128e0565b14610fc8576040517febf953c700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff81166000908152609c6020908152604080832054609e9092528220546110039190611b5e565b73ffffffffffffffffffffffffffffffffffffffff83166000908152609f6020526040812054919250906110379083612ba1565b905080851115611073576040517ff4d678b800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000806110886110838886612ba1565b611d85565b915091506110d7858383609f60008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054611dba565b6110e1888861183b565b60405187815273ffffffffffffffffffffffffffffffffffffffff89169033907f70eb43c4a8ae8c40502dcf22436c509c28d6ff421cf07c491be56984bd9870689060200160405180910390a350949695505050505050565b834211156111a4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f45524332305065726d69743a206578706972656420646561646c696e6500000060448201526064016108d2565b60007f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98888886111d38c611e7d565b60408051602081019690965273ffffffffffffffffffffffffffffffffffffffff94851690860152929091166060840152608083015260a082015260c0810186905260e001604051602081830303815290604052805190602001209050600061123b82611eb0565b9050600061124b82878787611ef8565b90508973ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146112e2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f45524332305065726d69743a20696e76616c6964207369676e6174757265000060448201526064016108d2565b6112ed8a8a8a6113e3565b50505050505050505050565b6000600273ffffffffffffffffffffffffffffffffffffffff8316600090815260a1602052604090205460ff166002811115611337576113376128e0565b1461136e576040517febf953c700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82166000908152609c6020908152604080832054609e9092528220546113a99190611b5e565b73ffffffffffffffffffffffffffffffffffffffff84166000908152609f60205260409020549091506113dc9082612ba1565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff8316611430576040517feb3b083500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff821661147d576040517f3b719e1800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff838116600081815260a2602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8216600090815260a1602052604081205460ff169060028281811115611529576115296128e0565b036115715773ffffffffffffffffffffffffffffffffffffffff84166000908152609c6020908152604080832054609e9092529091205461156a9190611b5e565b905061157d565b61157a8461098c565b90505b73ffffffffffffffffffffffffffffffffffffffff8416600090815260a16020526040902080548491907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660018360028111156115dd576115dd6128e0565b021790555073ffffffffffffffffffffffffffffffffffffffff84166000908152609f602052604090205461161485836001611f20565b6001836002811115611628576116286128e0565b03611645578060a0600082825461163f9190612ba1565b90915550505b6001846002811115611659576116596128e0565b03611676578160a060008282546116709190612a90565b90915550505b5050505050565b73ffffffffffffffffffffffffffffffffffffffff838116600090815260a260209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811461171e5780821115611711576040517f13be252b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61171e84848484036113e3565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316611771576040517f160fca8a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff82166117be576040517fea553b3400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6117c88382611cdc565b6117d2828261183b565b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516114df91815260200190565b60006106db612067565b6000816118478461098c565b6118519190612a90565b905061185f83826000611f20565b73ffffffffffffffffffffffffffffffffffffffff8316600090815260a1602052604090205460ff16600181600281111561189c5761189c6128e0565b0361171e578260a060008282546118b39190612a90565b909155505050505050565b60006118c9609d5490565b60695410806118d85750609d54155b156118e35750600090565b609d546069546118f39190612be7565b606860008282546119049190612a90565b9091555050609d546069546119199190612bfb565b6069556068546040519081527f270b316b51ab2cf3a3bb8ca4d22e76a327d05e762fcaa8bd6afaf8cfde9270b79060200160405180910390a150600190565b60608160000361199b57505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b81156119c557806119af81612c0f565b91506119be9050600a83612be7565b915061199f565b60008167ffffffffffffffff8111156119e0576119e0612b72565b6040519080825280601f01601f191660200182016040528015611a0a576020820181803683370190505b5090505b841561060c57611a1f600183612ba1565b9150611a2c600a86612bfb565b611a37906030612a90565b60f81b818381518110611a4c57611a4c612c47565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350611a86600a86612be7565b9450611a0e565b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000167fffffffffffffffffffffffffeeeeffffffffffffffffffffffffffffffffeeef330173ffffffffffffffffffffffffffffffffffffffff1614611b34576040517f3ae6ee0a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8015611b52578060696000828254611b4c9190612a90565b90915550505b611b5a6118be565b5050565b60008183606854611b6f9190612a53565b6113dc9190612a90565b600054610100900460ff16611c10576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016108d2565b611c19836120db565b611c22816121b1565b609a611c2e8482612cc5565b50609b61171e8382612cc5565b606060038054611c4a906129d7565b80601f0160208091040260200160405190810160405280929190818152602001828054611c76906129d7565b8015611cc35780601f10611c9857610100808354040283529160200191611cc3565b820191906000526020600020905b815481529060010190602001808311611ca657829003601f168201915b5050505050905090565b606060048054611c4a906129d7565b6000611ce78361098c565b905080821115611d23576040517ff4d678b800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611d31838383036000611f20565b73ffffffffffffffffffffffffffffffffffffffff8316600090815260a1602052604090205460ff166001816002811115611d6e57611d6e6128e0565b0361171e578260a060008282546118b39190612ba1565b600080606854600003611d96575091565b606854611da39084612be7565b915060685483611db39190612bfb565b9050915091565b73ffffffffffffffffffffffffffffffffffffffff84166000908152609c6020526040902054609d54611dee908590612a90565b611df89190612ba1565b609d5573ffffffffffffffffffffffffffffffffffffffff84166000908152609e602052604090205460a054611e2f908490612a90565b611e399190612ba1565b60a05573ffffffffffffffffffffffffffffffffffffffff9093166000908152609c6020908152604080832094909455609e815283822092909255609f9091522055565b73ffffffffffffffffffffffffffffffffffffffff81166000908152603560205260409020805460018101825590610a39565b60006106b6611ebd611831565b836040517f19010000000000000000000000000000000000000000000000000000000000008152600281019290925260228201526042902090565b6000806000611f0987878787612287565b91509150611f1681612376565b5095945050505050565b600080600080611f558773ffffffffffffffffffffffffffffffffffffffff16600090815260a1602052604090205460ff1690565b90506000816002811115611f6b57611f6b6128e0565b03611f8357611f7986611d85565b9094509250612052565b6001816002811115611f9757611f976128e0565b03611fa457859150612052565b6002816002811115611fb857611fb86128e0565b036120525785915081856120425773ffffffffffffffffffffffffffffffffffffffff88166000908152609c6020908152604080832054609e909252909120546120029190611b5e565b73ffffffffffffffffffffffffffffffffffffffff89166000908152609f60205260409020549091506120358883612a90565b61203f9190612ba1565b90505b61204b81611d85565b9095509350505b61205e87858585611dba565b50505050505050565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f612092612529565b61209a612582565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b600054610100900460ff16612172576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016108d2565b610989816040518060400160405280600181526020017f31000000000000000000000000000000000000000000000000000000000000008152506125b3565b600054610100900460ff16612248576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016108d2565b60685415612282576040517f4c72d1b400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606855565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156122be575060009050600361236d565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015612312573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff81166123665760006001925092505061236d565b9150600090505b94509492505050565b600081600481111561238a5761238a6128e0565b036123925750565b60018160048111156123a6576123a66128e0565b0361240d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064016108d2565b6002816004811115612421576124216128e0565b03612488576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060448201526064016108d2565b600381600481111561249c5761249c6128e0565b03610989576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f756500000000000000000000000000000000000000000000000000000000000060648201526084016108d2565b600080612534611c3b565b80519091501561254b578051602090910120919050565b600154801561255a5792915050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a4709250505090565b60008061258d611ccd565b8051909150156125a4578051602090910120919050565b600254801561255a5792915050565b600054610100900460ff1661264a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016108d2565b60036126568382612cc5565b5060046126638282612cc5565b50506000600181905560025550565b60006020828403121561268457600080fd5b81357fffffffff00000000000000000000000000000000000000000000000000000000811681146113dc57600080fd5b60005b838110156126cf5781810151838201526020016126b7565b8381111561171e5750506000910152565b600081518084526126f88160208601602086016126b4565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006113dc60208301846126e0565b803573ffffffffffffffffffffffffffffffffffffffff8116811461276157600080fd5b919050565b6000806040838503121561277957600080fd5b6127828361273d565b946020939093013593505050565b6000602082840312156127a257600080fd5b8135600381106113dc57600080fd5b6000806000606084860312156127c657600080fd5b6127cf8461273d565b92506127dd6020850161273d565b9150604084013590509250925092565b6000602082840312156127ff57600080fd5b5035919050565b60006020828403121561281857600080fd5b6113dc8261273d565b7fff00000000000000000000000000000000000000000000000000000000000000881681526000602060e08184015261285d60e084018a6126e0565b838103604085015261286f818a6126e0565b6060850189905273ffffffffffffffffffffffffffffffffffffffff8816608086015260a0850187905284810360c0860152855180825283870192509083019060005b818110156128ce578351835292840192918401916001016128b2565b50909c9b505050505050505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6003811061291f5761291f6128e0565b9052565b602081016106b6828461290f565b600080600080600080600060e0888a03121561294c57600080fd5b6129558861273d565b96506129636020890161273d565b95506040880135945060608801359350608088013560ff8116811461298757600080fd5b9699959850939692959460a0840135945060c09093013592915050565b600080604083850312156129b757600080fd5b6129c08361273d565b91506129ce6020840161273d565b90509250929050565b600181811c908216806129eb57607f821691505b602082108103610a39577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615612a8b57612a8b612a24565b500290565b60008219821115612aa357612aa3612a24565b500190565b60008451612aba8184602089016126b4565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551612af6816001850160208a016126b4565b60019201918201528351612b118160028401602088016126b4565b0160020195945050505050565b73ffffffffffffffffffffffffffffffffffffffff85811682526080820190612b4a602084018761290f565b60028510612b5a57612b5a6128e0565b84604084015280841660608401525095945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082821015612bb357612bb3612a24565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600082612bf657612bf6612bb8565b500490565b600082612c0a57612c0a612bb8565b500690565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203612c4057612c40612a24565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b601f821115612cc057600081815260208120601f850160051c81016020861015612c9d5750805b601f850160051c820191505b81811015612cbc57828155600101612ca9565b5050505b505050565b815167ffffffffffffffff811115612cdf57612cdf612b72565b612cf381612ced84546129d7565b84612c76565b602080601f831160018114612d465760008415612d105750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555612cbc565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b82811015612d9357888601518255948401946001909101908401612d74565b5085821015612dcf57878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b0190555056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(USDBStorageLayoutJSON), USDBStorageLayout); err != nil {
		panic(err)
	}

	layouts["USDB"] = USDBStorageLayout
	deployedBytecodes["USDB"] = USDBDeployedBin
}
