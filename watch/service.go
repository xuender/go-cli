package watch

import (
	"context"
	"errors"
	"os"
	"os/exec"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/samber/lo"
	"github.com/xuender/kit/logs"
)

type Service struct {
	watcher *fsnotify.Watcher
	cancel  context.CancelFunc
}

func NewService() *Service {
	return &Service{lo.Must1(fsnotify.NewWatcher()), nil}
}

func (p *Service) Add(path string) {
	lo.Must0(p.watcher.Add(path))
}

func (p *Service) Run(name string, args []string) {
	go p.watch()

	for {
		ctx, can := context.WithCancel(context.Background())
		com := exec.CommandContext(ctx, name, args...)
		com.Stdin = os.Stdin
		com.Stdout = os.Stdout
		p.cancel = can

		lo.Must0(com.Start())

		if err := com.Wait(); err.Error() == "signal: killed" || errors.Is(err, context.Canceled) {
			time.Sleep(time.Second)
		} else {
			return
		}
	}
}

func (p *Service) watch() {
	for {
		select {
		case event, ok := <-p.watcher.Events:
			if !ok {
				return
			}

			logs.D.Println("event:", event)

			if event.Has(fsnotify.Write) {
				if p.cancel != nil {
					p.cancel()
					p.cancel = nil
				}
			}
		case err, ok := <-p.watcher.Errors:
			if !ok {
				return
			}

			if p.cancel != nil {
				logs.E.Println("error:", err)
				p.cancel()
				p.cancel = nil
			}
		}
	}
}

func (p *Service) Close() {
	p.watcher.Close()
}
