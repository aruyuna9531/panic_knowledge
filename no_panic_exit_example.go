package panic_recover

import (
	"fmt"
	"os"
	"os/signal"
)

var ClientSignalInput chan interface{}

func main() {
	// 其他初始化代码（这里报panic将宕机）
	Loop()
}

func Loop() {
	osSig := make(chan os.Signal)
	signal.Notify(osSig)
	for {
		select {
		case sig := <-osSig:
			// 处理系统信号（一般是正常退出程序）
			_ = sig
		case sig := <-ClientSignalInput:
			// 客户端信号（一般是业务线）
			SigFunc(sig)
		}
	}
}

func SigFunc(sig interface{}) {
	defer PanicRecoverTrace() // 会抓获DoSig里的panic，这个recover一定是新开函数并写在函数里，不能因为只有2行代码就直接写到case下面
	DoSig(sig)
}

func DoSig(sig interface{}) {
	// 业务线的深处
	i := 3
	i--
	i--
	i--
	s := 3 / i // 除数不能为0
	fmt.Println(s)
}
