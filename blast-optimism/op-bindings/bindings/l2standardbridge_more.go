// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const L2StandardBridgeStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L2/L2StandardBridge.sol:L2StandardBridge\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"src/L2/L2StandardBridge.sol:L2StandardBridge\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"src/L2/L2StandardBridge.sol:L2StandardBridge\",\"label\":\"spacer_0_2_20\",\"offset\":2,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":1003,\"contract\":\"src/L2/L2StandardBridge.sol:L2StandardBridge\",\"label\":\"spacer_1_0_20\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_address\"},{\"astId\":1004,\"contract\":\"src/L2/L2StandardBridge.sol:L2StandardBridge\",\"label\":\"deposits\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_mapping(t_address,t_mapping(t_address,t_uint256))\"},{\"astId\":1005,\"contract\":\"src/L2/L2StandardBridge.sol:L2StandardBridge\",\"label\":\"messenger\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_contract(CrossDomainMessenger)1007\"},{\"astId\":1006,\"contract\":\"src/L2/L2StandardBridge.sol:L2StandardBridge\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"4\",\"type\":\"t_array(t_uint256)46_storage\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)46_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[46]\",\"numberOfBytes\":\"1472\",\"base\":\"t_uint256\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_contract(CrossDomainMessenger)1007\":{\"encoding\":\"inplace\",\"label\":\"contract CrossDomainMessenger\",\"numberOfBytes\":\"20\"},\"t_mapping(t_address,t_mapping(t_address,t_uint256))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(address =\u003e uint256))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_address,t_uint256)\"},\"t_mapping(t_address,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_uint256\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L2StandardBridgeStorageLayout = new(solc.StorageLayout)

var L2StandardBridgeDeployedBin = "0x6080604052600436106101125760003560e01c8063662a633a116100a55780638f601f6611610074578063a3a7954811610059578063a3a7954814610428578063c89701a21461023f578063e11013dd1461043b57600080fd5b80638f601f66146103b7578063927ede2d146103fd57600080fd5b8063662a633a1461033b5780637f46ddb21461034e5780638129fc1c14610382578063870876231461039757600080fd5b806336c717c1116100e157806336c717c11461023f5780633cb747bf14610298578063540abf73146102c557806354fd4d50146102e557600080fd5b80630166a07a146101e657806309fc8843146102065780631635f5fd1461021957806332b7006d1461022c57600080fd5b366101e157333b156101ab576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f4100000000000000000060648201526084015b60405180910390fd5b6101df73deaddeaddeaddeaddeaddeaddeaddeaddead000033333462030d406040518060200160405280600081525061044e565b005b600080fd5b3480156101f257600080fd5b506101df6102013660046124cd565b610529565b6101df61021436600461257e565b6108e7565b6101df6102273660046125d1565b6109be565b6101df61023a366004612644565b610e2d565b34801561024b57600080fd5b507f00000000000000000000000000000000000000000000000000000000000000005b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b3480156102a457600080fd5b5060035461026e9073ffffffffffffffffffffffffffffffffffffffff1681565b3480156102d157600080fd5b506101df6102e0366004612698565b610f07565b3480156102f157600080fd5b5061032e6040518060400160405280600581526020017f312e342e3000000000000000000000000000000000000000000000000000000081525081565b60405161028f9190612785565b6101df6103493660046124cd565b610f4c565b34801561035a57600080fd5b5061026e7f000000000000000000000000000000000000000000000000000000000000000081565b34801561038e57600080fd5b506101df610fbf565b3480156103a357600080fd5b506101df6103b2366004612798565b6111a8565b3480156103c357600080fd5b506103ef6103d236600461281b565b600260209081526000928352604080842090915290825290205481565b60405190815260200161028f565b34801561040957600080fd5b5060035473ffffffffffffffffffffffffffffffffffffffff1661026e565b6101df610436366004612798565b61127c565b6101df610449366004612854565b6112c0565b7fffffffffffffffffffffffff215221522152215221522152215221522153000073ffffffffffffffffffffffffffffffffffffffff87160161049d576104988585858585611309565b610521565b60008673ffffffffffffffffffffffffffffffffffffffff1663c01e1bd66040518163ffffffff1660e01b8152600401602060405180830381865afa1580156104ea573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061050e91906128b7565b905061051f878288888888886114ef565b505b505050505050565b60035473ffffffffffffffffffffffffffffffffffffffff16331480156106185750600354604080517f6e296e45000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000008116931691636e296e459160048083019260209291908290030181865afa1580156105dc573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061060091906128b7565b73ffffffffffffffffffffffffffffffffffffffff16145b6106ca576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604160248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20746865206f7468657220627269646760648201527f6500000000000000000000000000000000000000000000000000000000000000608482015260a4016101a2565b6106d387611837565b15610821576106e28787611899565b610794576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604a60248201527f5374616e646172644272696467653a2077726f6e672072656d6f746520746f6b60448201527f656e20666f72204f7074696d69736d204d696e7461626c65204552433230206c60648201527f6f63616c20746f6b656e00000000000000000000000000000000000000000000608482015260a4016101a2565b6040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8581166004830152602482018590528816906340c10f1990604401600060405180830381600087803b15801561080457600080fd5b505af1158015610818573d6000803e3d6000fd5b505050506108a3565b73ffffffffffffffffffffffffffffffffffffffff8088166000908152600260209081526040808320938a168352929052205461085f908490612903565b73ffffffffffffffffffffffffffffffffffffffff8089166000818152600260209081526040808320948c16835293905291909120919091556108a39085856119b9565b61051f878787878787878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611a8d92505050565b333b15610976576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f4100000000000000000060648201526084016101a2565b6109b93333348686868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061130992505050565b505050565b60035473ffffffffffffffffffffffffffffffffffffffff1633148015610aad5750600354604080517f6e296e45000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000008116931691636e296e459160048083019260209291908290030181865afa158015610a71573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a9591906128b7565b73ffffffffffffffffffffffffffffffffffffffff16145b610b5f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604160248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20746865206f7468657220627269646760648201527f6500000000000000000000000000000000000000000000000000000000000000608482015260a4016101a2565b82341115610bef576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603a60248201527f5374616e646172644272696467653a20616d6f756e742073656e7420646f657360448201527f206e6f74206d6174636820616d6f756e7420726571756972656400000000000060648201526084016101a2565b3073ffffffffffffffffffffffffffffffffffffffff851603610c94576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f5374616e646172644272696467653a2063616e6e6f742073656e6420746f207360448201527f656c66000000000000000000000000000000000000000000000000000000000060648201526084016101a2565b60035473ffffffffffffffffffffffffffffffffffffffff90811690851603610d3f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602860248201527f5374616e646172644272696467653a2063616e6e6f742073656e6420746f206d60448201527f657373656e67657200000000000000000000000000000000000000000000000060648201526084016101a2565b610d8185853485858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611b1b92505050565b6000610d9e855a3460405180602001604052806000815250611bbc565b905080610521576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f5374616e646172644272696467653a20455448207472616e736665722066616960448201527f6c6564000000000000000000000000000000000000000000000000000000000060648201526084016101a2565b333b15610ebc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f4100000000000000000060648201526084016101a2565b610f00853333878787878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061044e92505050565b5050505050565b61051f87873388888888888080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506114ef92505050565b73ffffffffffffffffffffffffffffffffffffffff8716158015610f99575073ffffffffffffffffffffffffffffffffffffffff861673deaddeaddeaddeaddeaddeaddeaddeaddead0000145b15610fb057610fab85858585856109be565b61051f565b61051f86888787878787610529565b600054600390610100900460ff16158015610fe1575060005460ff8083169116105b61106d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016101a2565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001660ff8316176101001790556110bb734200000000000000000000000000000000000007611bd6565b6040517f4c802f3800000000000000000000000000000000000000000000000000000000815273430000000000000000000000000000000000000290634c802f389061111590309060019060009061dead90600401612949565b600060405180830381600087803b15801561112f57600080fd5b505af1158015611143573d6000803e3d6000fd5b5050600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055505060405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a150565b333b15611237576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f4100000000000000000060648201526084016101a2565b61052186863333888888888080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506114ef92505050565b610521863387878787878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061044e92505050565b6113033385348686868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061130992505050565b50505050565b823414611398576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603e60248201527f5374616e646172644272696467653a206272696467696e6720455448206d757360448201527f7420696e636c7564652073756666696369656e74204554482076616c7565000060648201526084016101a2565b6113a485858584611cb4565b60035460405173ffffffffffffffffffffffffffffffffffffffff90911690633dbb202b9085907f0000000000000000000000000000000000000000000000000000000000000000907f1635f5fd0000000000000000000000000000000000000000000000000000000090611423908b908b9086908a906024016129a6565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e086901b90921682526114b6929188906004016129ef565b6000604051808303818588803b1580156114cf57600080fd5b505af11580156114e3573d6000803e3d6000fd5b50505050505050505050565b6114f887611837565b15611646576115078787611899565b6115b9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604a60248201527f5374616e646172644272696467653a2077726f6e672072656d6f746520746f6b60448201527f656e20666f72204f7074696d69736d204d696e7461626c65204552433230206c60648201527f6f63616c20746f6b656e00000000000000000000000000000000000000000000608482015260a4016101a2565b6040517f9dc29fac00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff868116600483015260248201859052881690639dc29fac90604401600060405180830381600087803b15801561162957600080fd5b505af115801561163d573d6000803e3d6000fd5b505050506116da565b61166873ffffffffffffffffffffffffffffffffffffffff8816863086611d55565b73ffffffffffffffffffffffffffffffffffffffff8088166000908152600260209081526040808320938a16835292905220546116a6908490612a34565b73ffffffffffffffffffffffffffffffffffffffff8089166000908152600260209081526040808320938b16835292905220555b6116e8878787878786611db3565b60035460405173ffffffffffffffffffffffffffffffffffffffff90911690633dbb202b907f0000000000000000000000000000000000000000000000000000000000000000907f0166a07a0000000000000000000000000000000000000000000000000000000090611769908b908d908c908c908c908b90602401612a4c565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e085901b90921682526117fc929187906004016129ef565b600060405180830381600087803b15801561181657600080fd5b505af115801561182a573d6000803e3d6000fd5b5050505050505050505050565b6000611863827f1d1d8b6300000000000000000000000000000000000000000000000000000000611e41565b806118935750611893827fec4fc8e300000000000000000000000000000000000000000000000000000000611e41565b92915050565b60006118c5837f1d1d8b6300000000000000000000000000000000000000000000000000000000611e41565b1561196e578273ffffffffffffffffffffffffffffffffffffffff1663c01e1bd66040518163ffffffff1660e01b8152600401602060405180830381865afa158015611915573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061193991906128b7565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16149050611893565b8273ffffffffffffffffffffffffffffffffffffffff1663d6c0b2c46040518163ffffffff1660e01b8152600401602060405180830381865afa158015611915573d6000803e3d6000fd5b60405173ffffffffffffffffffffffffffffffffffffffff83166024820152604481018290526109b99084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152611e64565b8373ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff167fb0444523268717a02698be47d0803aa7468c00acbed2f8bd93a0459cde61dd89868686604051611b0593929190612aa7565b60405180910390a4610521868686868686611f70565b8373ffffffffffffffffffffffffffffffffffffffff1673deaddeaddeaddeaddeaddeaddeaddeaddead000073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fb0444523268717a02698be47d0803aa7468c00acbed2f8bd93a0459cde61dd89868686604051611ba893929190612aa7565b60405180910390a461130384848484611ff8565b600080600080845160208601878a8af19695505050505050565b600054610100900460ff16611c6d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016101a2565b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b8373ffffffffffffffffffffffffffffffffffffffff1673deaddeaddeaddeaddeaddeaddeaddeaddead000073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f73d170910aba9e6d50b102db522b1dbcd796216f5128b445aa2135272886497e868686604051611d4193929190612aa7565b60405180910390a461130384848484612065565b60405173ffffffffffffffffffffffffffffffffffffffff808516602483015283166044820152606481018290526113039085907f23b872dd0000000000000000000000000000000000000000000000000000000090608401611a0b565b8373ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff167f73d170910aba9e6d50b102db522b1dbcd796216f5128b445aa2135272886497e868686604051611e2b93929190612aa7565b60405180910390a46105218686868686866120c4565b6000611e4c8361213c565b8015611e5d5750611e5d83836121a0565b9392505050565b6000611ec6826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff1661226f9092919063ffffffff16565b8051909150156109b95780806020019051810190611ee49190612ae5565b6109b9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016101a2565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167fd59c65b35445225835c83f50b6ede06a7be047d22e357073e250d9af537518cd868686604051611fe893929190612aa7565b60405180910390a4505050505050565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f31b2166ff604fc5672ea5df08a78081d2bc6d746cadce880747f3643d819e83d8484604051612057929190612b07565b60405180910390a350505050565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f2849b43074093a05396b6f2a937dee8565b15a48a7b3d4bffb732a5017380af58484604051612057929190612b07565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167f7ff126db8024424bbfd9826e8ab82ff59136289ea440b04b39a0df1b03b9cabf868686604051611fe893929190612aa7565b6000612168827f01ffc9a7000000000000000000000000000000000000000000000000000000006121a0565b80156118935750612199827fffffffff000000000000000000000000000000000000000000000000000000006121a0565b1592915050565b604080517fffffffff000000000000000000000000000000000000000000000000000000008316602480830191909152825180830390910181526044909101909152602080820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01ffc9a700000000000000000000000000000000000000000000000000000000178152825160009392849283928392918391908a617530fa92503d91506000519050828015612258575060208210155b80156122645750600081115b979650505050505050565b606061227e8484600085612286565b949350505050565b606082471015612318576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016101a2565b73ffffffffffffffffffffffffffffffffffffffff85163b612396576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016101a2565b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516123bf9190612b20565b60006040518083038185875af1925050503d80600081146123fc576040519150601f19603f3d011682016040523d82523d6000602084013e612401565b606091505b50915091506122648282866060831561241b575081611e5d565b82511561242b5782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101a29190612785565b73ffffffffffffffffffffffffffffffffffffffff8116811461248157600080fd5b50565b60008083601f84011261249657600080fd5b50813567ffffffffffffffff8111156124ae57600080fd5b6020830191508360208285010111156124c657600080fd5b9250929050565b600080600080600080600060c0888a0312156124e857600080fd5b87356124f38161245f565b965060208801356125038161245f565b955060408801356125138161245f565b945060608801356125238161245f565b93506080880135925060a088013567ffffffffffffffff81111561254657600080fd5b6125528a828b01612484565b989b979a50959850939692959293505050565b803563ffffffff8116811461257957600080fd5b919050565b60008060006040848603121561259357600080fd5b61259c84612565565b9250602084013567ffffffffffffffff8111156125b857600080fd5b6125c486828701612484565b9497909650939450505050565b6000806000806000608086880312156125e957600080fd5b85356125f48161245f565b945060208601356126048161245f565b935060408601359250606086013567ffffffffffffffff81111561262757600080fd5b61263388828901612484565b969995985093965092949392505050565b60008060008060006080868803121561265c57600080fd5b85356126678161245f565b94506020860135935061267c60408701612565565b9250606086013567ffffffffffffffff81111561262757600080fd5b600080600080600080600060c0888a0312156126b357600080fd5b87356126be8161245f565b965060208801356126ce8161245f565b955060408801356126de8161245f565b9450606088013593506126f360808901612565565b925060a088013567ffffffffffffffff81111561254657600080fd5b60005b8381101561272a578181015183820152602001612712565b838111156113035750506000910152565b6000815180845261275381602086016020860161270f565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000611e5d602083018461273b565b60008060008060008060a087890312156127b157600080fd5b86356127bc8161245f565b955060208701356127cc8161245f565b9450604087013593506127e160608801612565565b9250608087013567ffffffffffffffff8111156127fd57600080fd5b61280989828a01612484565b979a9699509497509295939492505050565b6000806040838503121561282e57600080fd5b82356128398161245f565b915060208301356128498161245f565b809150509250929050565b6000806000806060858703121561286a57600080fd5b84356128758161245f565b935061288360208601612565565b9250604085013567ffffffffffffffff81111561289f57600080fd5b6128ab87828801612484565b95989497509550505050565b6000602082840312156128c957600080fd5b8151611e5d8161245f565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082821015612915576129156128d4565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff85811682526080820190600386106129785761297861291a565b8560208401526002851061298e5761298e61291a565b84604084015280841660608401525095945050505050565b600073ffffffffffffffffffffffffffffffffffffffff8087168352808616602084015250836040830152608060608301526129e5608083018461273b565b9695505050505050565b73ffffffffffffffffffffffffffffffffffffffff84168152606060208201526000612a1e606083018561273b565b905063ffffffff83166040830152949350505050565b60008219821115612a4757612a476128d4565b500190565b600073ffffffffffffffffffffffffffffffffffffffff80891683528088166020840152808716604084015280861660608401525083608083015260c060a0830152612a9b60c083018461273b565b98975050505050505050565b73ffffffffffffffffffffffffffffffffffffffff84168152826020820152606060408201526000612adc606083018461273b565b95945050505050565b600060208284031215612af757600080fd5b81518015158114611e5d57600080fd5b82815260406020820152600061227e604083018461273b565b60008251612b3281846020870161270f565b919091019291505056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(L2StandardBridgeStorageLayoutJSON), L2StandardBridgeStorageLayout); err != nil {
		panic(err)
	}

	layouts["L2StandardBridge"] = L2StandardBridgeStorageLayout
	deployedBytecodes["L2StandardBridge"] = L2StandardBridgeDeployedBin
}
