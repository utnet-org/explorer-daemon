package pkg

// SearchType 查询类型
type SearchType int

const (
	SearchAccount     SearchType = iota + 1 //账户
	SearchBlockHeight                       //区块
	SearchBlockHash                         //区块
	SearchAddress                           //地址
	SearchTransaction                       //交易Hash
	SearchMessage                           //消息
	SearchChip                              //芯片
)

func (s SearchType) ToString() string {
	switch s {
	case SearchAccount:
		return "账户"
	case SearchBlockHeight:
		return "区块"
	case SearchBlockHash:
		return "区块"
	case SearchAddress:
		return "地址"
	case SearchTransaction:
		return "交易Hash"
	case SearchMessage:
		return "消息"
	default:
		panic("unhandled default case")
	}
	return ""
}

// BlockQueryType 查询类型
type BlockQueryType int

const (
	BlockQueryHeight BlockQueryType = iota + 1
	BlockQueryHash
	BlockQueryFinal
)

type ChunkQueryType int

const (
	ChunkQueryBlockHeight ChunkQueryType = iota + 1
	ChunkQueryBlockHash
	ChunkQueryChunkHash
)

// BlockChangeRpcType
type BlockChangeRpcType int

const (
	BlockChangeRpcFinal BlockChangeRpcType = iota + 1
	BlockChangeRpcHeight
	BlockChangeRpcHash
)
