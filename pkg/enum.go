package pkg

// SearchType 查询类型
type SearchType int

const (
	SearchAccount     SearchType = iota + 1 //账户
	SearchBlock                             //区块
	SearchAddress                           //地址
	SearchTransaction                       //交易Hash
	SearchMessage                           //消息
)

func (s SearchType) ToString() string {
	switch s {
	case SearchAccount:
		return "账户"
	case SearchBlock:
		return "区块"
	case SearchAddress:
		return "地址"
	case SearchTransaction:
		return "交易Hash"
	case SearchMessage:
		return "消息"
	}
	return ""
}

// BlockQueryType 查询类型
type QueryType int

const (
	BlockQueryHeight QueryType = iota + 1
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
