package watch

import (
	"context"
	"errors"
	"io/fs"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/samber/lo"
	"github.com/xuender/kit/logs"
)

type Service struct {
	watcher *fsnotify.Watcher
	cancel  context.CancelFunc
	pid     int
	lock    *sync.Cond
}

func NewService() *Service {
	return &Service{
		watcher: lo.Must1(fsnotify.NewWatcher()),
		lock:    sync.NewCond(&sync.Mutex{}),
	}
}

func (p *Service) Add(path string) {
	lo.Must0(filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return p.watcher.Add(path)
		}

		return nil
	}))
}

func (p *Service) Run(cmd string, args []string) {
	go p.watch()
	go p.exec(cmd, args)

	osc := make(chan os.Signal, 1)

	signal.Notify(osc, syscall.SIGTERM, syscall.SIGINT)

	num := <-osc

	p.Close()
	logs.D.Println("监听到退出信号:", num)
}

func (p *Service) exec(command string, args []string) {
	for {
		ctx, can := context.WithCancel(context.Background())
		cmd := exec.CommandContext(ctx, command, args...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
		p.cancel = can

		lo.Must0(cmd.Start())
		p.pid = cmd.Process.Pid

		if err := cmd.Wait(); err != nil && (err.Error() == "signal: killed" || errors.Is(err, context.Canceled)) {
			lo.Must0(syscall.Kill(-p.pid, syscall.SIGKILL))
		} else if err != nil {
			logs.Log(err)
			p.lock.L.Lock()
			p.lock.Wait()
			p.lock.L.Unlock()
		} else {
			p.pid = 0
		}

		time.Sleep(time.Second)
	}
}

func (p *Service) unlock() {
	if p.cancel != nil {
		p.cancel()
		p.cancel = nil
	}

	p.lock.Signal()
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
				p.unlock()
			}
		case err, ok := <-p.watcher.Errors:
			if !ok {
				return
			}

			logs.E.Println("error:", err)
			p.unlock()
		}
	}
}

func (p *Service) Close() {
	p.watcher.Close()

	if p.pid > 0 {
		lo.Must0(syscall.Kill(-p.pid, syscall.SIGKILL))
	}
}
