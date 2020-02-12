package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	stopCh := make(chan struct{})
	wg := sync.WaitGroup{}

	wg.Add(1)
	go singnalSelect(stopCh, &wg)

	wg.Add(1)
	go loop(stopCh, &wg)

	wg.Add(1)
	go loop(stopCh, &wg)

	wg.Add(1)
	go func(stopCh chan struct{}, wg *sync.WaitGroup) {
		defer func() {
			println("inner loop defer")
			wg.Done()
		}()

		for {
			println("(goroutine) inner loop...")
			time.Sleep(2 * time.Second)
			select {
			case <-stopCh:
				println("(goroutine) innner loop stop request received.")
				return
			default:
			}
		}

	}(stopCh, &wg)

	wg.Wait()
	println("(main) all loops done.")
}

func singnalSelect(stopCh chan struct{}, wg *sync.WaitGroup) {
	defer func() {
		println("signal defer")
		wg.Done()
	}()
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh,
		syscall.SIGTERM,
		syscall.SIGINT)
	defer signal.Stop(sigCh)
	for {
		select {
		case sig := <-sigCh:
			switch sig {
			case syscall.SIGTERM:
				fmt.Println("SIGTERM")
			case syscall.SIGINT:
				fmt.Println("SIGINT")
				close(stopCh)
			}
		case <-stopCh:
			fmt.Println("Done")
			return
		}
	}
}

func loop(stopCh chan struct{}, wg *sync.WaitGroup) {
	// リソースの解放が必要であれば、ここでする
	defer func() {
		println("loop defer")
		wg.Done()
	}()

	// 無限ループ
	for {
		select {
		case <-stopCh:
			println("(goroutine) stop request received.")
			return
		default:
			println("(goroutine) loop...")
			time.Sleep(1 * time.Second)
		}
	}
}

// package main

// import (
//     "fmt"
//     "os"
//     "os/signal"
//     "syscall"
//     "sync"
// )

// func main() {
//     wg := sync.WaitGroup{}
//     doneCh := make(chan struct{}, 1)

//     wg.Add(1)
//     go func() {
//         defer wg.Done()

//         sigCh := make(chan os.Signal, 1)
//         signal.Notify(sigCh,
//             syscall.SIGTERM,
//             syscall.SIGINT)
//         defer signal.Stop(sigCh)

//         for {
//             select {
//             case sig := <-sigCh:
//                 switch(sig) {
//                 case syscall.SIGTERM:
//                     fmt.Println("SIGTERM")
//                 case syscall.SIGINT:
//                     fmt.Println("SIGINT")
//                 }
//             case <-doneCh:
//                 fmt.Println("Done")
//                 return
//             }
//         }
//     }()

//     // 何かする
//     wg.Add(1)
//     go func() {
//     	defer wg.Done()

//     }

//     // 終了
//     doneCh <- struct{}{}
//     wg.Wait()
// }

// package main

// import (
//     "time"
//     "sync"
// )

// func main() {
//     stopCh := make(chan struct{})
//     wg := sync.WaitGroup{}

//     wg.Add(1)
//     go loop(stopCh, &wg)

//     wg.Add(1)
//     go loop(stopCh, &wg)

//     time.Sleep(3 * time.Second)
//     println("(main) request stop.")
//     close(stopCh)

//     // loopが完了するまで待つ
//     wg.Wait()
//     println("(main) all loops done.")
// }

// // stopCh: 外部からこのloopを止めたい という通知をするchannel
// // wg: このloopが完了したら Done() を呼び出す
// func loop(stopCh chan struct{}, wg *sync.WaitGroup) {
//     // リソースの解放が必要であれば、ここでする
//     defer func() { wg.Done() }()

//     // 無限ループ
//     for {
//         println("(goroutine) loop...")
//         time.Sleep(1 * time.Second)

//         select {
//         case <- stopCh:
//             println("(goroutine) stop request received.")
//             return
//         default:
//             // default: 重要！！
//             // これがないと、stopChを待ち続けてしまうので、loopが続行できなくなる
//         }
//     }
// }

// package main

// import (
//     "fmt"
//     "os"
//     "os/signal"
//     "syscall"
// )

// func SelectSignal(signal_chan chan os.Signal, exit_chan chan int) {
//     for {
//             fmt.Println("inloop")
//             s := <-signal_chan
//             switch s {
//             // kill -SIGHUP XXXX
//             case syscall.SIGHUP:
//                 fmt.Println("hungup")

//             // kill -SIGINT XXXX or Ctrl+c
//             case syscall.SIGINT:
//                 fmt.Println("Warikomi")
//                 exit_chan <- 0

//             // kill -SIGTERM XXXX
//             case syscall.SIGTERM:
//                 fmt.Println("force stop")
//                 exit_chan <- 0

//             // kill -SIGQUIT XXXX
//             case syscall.SIGQUIT:
//                 fmt.Println("stop and core dump")
//                 exit_chan <- 0

//             default:
//                 fmt.Println("Unknown signal.")
//                 exit_chan <- 1
//             }
//         }
// }

// func main() {
//     signal_chan := make(chan os.Signal, 1)
//     signal.Notify(signal_chan,
//         syscall.SIGHUP,
//         syscall.SIGINT,
//         syscall.SIGTERM,
//         syscall.SIGQUIT)
//     exit_chan := make(chan int)
//     go SelectSignal(signal_chan, exit_chan);
//     fmt.Println("loop1")
//     code := <-exit_chan
//     fmt.Println("loop2")
//     os.Exit(code)
// }
