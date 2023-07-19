package watch

import (
	"context"
	"errors"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"syscall"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/samber/lo"
	"github.com/xuender/kit/logs"
	"github.com/xuender/kit/oss"
	"golang.org/x/sync/errgroup"
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
	ctx, cancel := context.WithCancel(context.Background())
	group := errgroup.Group{}

	group.Go(func() error {
		return p.watch(ctx)
	})
	group.Go(func() error {
		return p.exec(ctx, cmd, args)
	})
	group.Go(oss.CancelFunc(cancel))

	lo.Must0(group.Wait())
}

func (p *Service) exec(parent context.Context, command string, args []string) error {
	for {
		select {
		case <-parent.Done():
			return nil
		default:
			ctx, can := context.WithCancel(parent)
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
}

func (p *Service) unlock() {
	if p.cancel == nil {
		p.lock.Signal()
	} else {
		p.cancel()
		p.cancel = nil
	}
}

func (p *Service) watch(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case event, ok := <-p.watcher.Events:
			if !ok {
				return nil
			}

			logs.D.Println("event:", event)

			if event.Has(fsnotify.Write) {
				p.unlock()
			}
		case err, ok := <-p.watcher.Errors:
			if !ok {
				return err
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
