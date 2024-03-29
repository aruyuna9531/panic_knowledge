package panic_recover

import (
	"fmt"
	"log"
	"runtime"
)

func PanicRecoverTrace() {
	r := recover()
	if r == nil {
		// 无事发生
		return
	}

	buf := fmt.Sprintf("Panic called, message: [%v], Trace: \n", r)
	// 下面其实跟debug.StackTrace()差不多 不过那个分行看着好别扭
	traceSkip := 1
	for {
		pc, file, line, ok := runtime.Caller(traceSkip)
		if !ok {
			break
		}
		buf += fmt.Sprintf("\t%d. %s %s:%d \n", traceSkip, runtime.FuncForPC(pc).Name(), file, line)
		traceSkip++
	}
	log.Println(buf)
}
