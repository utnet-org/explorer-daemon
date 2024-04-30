package fetch

func InitFetchData() {
	// 定时执行RPC请求
	//ticker := time.NewTicker(time.Hour) // 例如，每小时执行一次
	//for range ticker.C {
	BlockDetailsByFinal()
	//BlockChangesRpc()
	//HandleNetworkInfo()
	//HandleChipQuery()
	//}
}
