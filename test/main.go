package main

import (
	"compress/flate"
	"os"
	"runtime/pprof"
	"sync"

	graylog "github.com/johnnyluo/logrus-graylog-hook"
	log "github.com/sirupsen/logrus"
)

func main() {
	f, err := os.Create("cpu.out")
	if nil != err {
		panic(err)
	}
	defer f.Close()
	if err := pprof.StartCPUProfile(f); nil != err {
		panic(err)
	}
	defer pprof.StopCPUProfile()
	m := map[string]interface{}{
		"hello": 123,
	}

	h := graylog.NewGraylogHook("127.0.0.1:12201",
		graylog.WithHost("myhost"),
		graylog.WithCompressType(graylog.CompressGzip),
		graylog.WithCompressLevel(flate.NoCompression),
		graylog.WithExtra(m))
	log.AddHook(h)
	defer h.Flush()
	// for i := 0; i < 10000; i++ {
	// 	log.Infof("HelloWorld:%d", i)
	// }
	// for i := 0; i < 10000; i++ {
	// 	log.WithError(errors.New("my error")).Infoln("error")
	// }
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go writeLogInParallel(i, wg)
	}
	wg.Wait()
	mf, err := os.Create("memory.out")
	if nil != err {
		panic(err)
	}
	defer mf.Close()
	if err := pprof.WriteHeapProfile(mf); nil != err {
		panic(err)
	}
}

func writeLogInParallel(idx int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		log.WithField("goroutine", idx).WithField("counter", i).Infoln("Hi")
	}
}
