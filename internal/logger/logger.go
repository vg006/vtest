// logger/logger.go
package logger

import (
	"fmt"
	"sync"

	asset "github.com/vg006/vtest/internal/assets"
	"github.com/vg006/vtest/internal/types"
)

type Logger struct {
	MsgChan chan types.Msg
	wg      sync.WaitGroup
	mu      sync.Mutex
	closed  bool
}

func New(bufSize int) *Logger {
	return &Logger{
		MsgChan: make(chan types.Msg, bufSize),
		closed:  false,
	}
}

func (l *Logger) Exec() {
	l.wg.Add(1)
	go func() {
		defer l.wg.Done()
		select {
		case msg := <-l.MsgChan:
			switch msg.Type {
			case types.MsgInfo:
				fmt.Println(asset.Info + asset.Msg.Render(msg.Message))
			case types.MsgDone:
				fmt.Println(asset.Done + asset.Msg.Render(msg.Message))
			case types.MsgError:
				fmt.Println(asset.Error + asset.Msg.Render(msg.Message))
			case types.MsgExit:
				fmt.Println(asset.Exit + asset.Msg.Render(msg.Message))
			}
		}
	}()
}

func (l *Logger) send(msg types.Msg) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.closed {
		fmt.Println("Logger is shut down. Message dropped:", msg)
		return
	}
	l.MsgChan <- msg
}

func (l *Logger) LogInfo(message string) {
	l.send(types.Msg{Type: types.MsgInfo, Message: message})
}

func (l *Logger) LogError(message string) {
	l.send(types.Msg{Type: types.MsgError, Message: message})
}

func (l *Logger) LogDone(message string) {
	l.send(types.Msg{Type: types.MsgDone, Message: message})
}

func (l *Logger) Exit(msg string) {
	if msg == "" {
		msg = "Terminating..."
	}
	l.send(types.Msg{Type: types.MsgExit, Message: msg})
	l.mu.Lock()
	if !l.closed {
		close(l.MsgChan)
		l.closed = true
	}
	l.mu.Unlock()
	l.wg.Wait()
}
