package shut

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	asset "github.com/vg006/vtest/internal/assets"
)

var (
	ShutdownChan = make(chan struct{})
	once         sync.Once
)

func GracefulExit() {
	once.Do(func() {
		fmt.Println("[SHUTDOWN] Initiating global shutdown...")
		close(ShutdownChan)
	})
}

func ListenForOSSignals() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-sigChan
		fmt.Println(asset.Exit + asset.Msg.Render(sig.String()))
		GracefulExit()
	}()
}
