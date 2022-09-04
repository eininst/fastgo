package grace

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/eininst/flog"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp/reuseport"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

const FIBER_CHILD_LOCK_FILE = "/tmp/fiber_child.lock"

const (
	envReload    = "FIBER_ENV_RELOAD"
	envReloadVal = "1"
)

func IsReload() bool {
	return os.Getenv(envReload) == envReloadVal
}

var DefaultTimeout = time.Second * 10

var pids []int

func Listen(app *fiber.App, addr string, timeout ...time.Duration) {
	t := DefaultTimeout
	if len(timeout) > 0 {
		t = timeout[0]
	}
	go func() {
		if app.Config().Prefork {
			pids = []int{}
			app.Hooks().OnFork(func(i int) error {
				pids = append(pids, i)
				return nil
			})
			_ = app.Listen(addr)
		} else {
			if IsReload() {
				ln, _ := reuseport.Listen(app.Config().Network, addr)
				_ = app.Listener(ln)
			} else {
				_ = app.Listen(addr)
			}
		}
	}()
	listenSig(app, t)
}

func ListenTLS(app *fiber.App, addr string, certFile, keyFile string, timeout ...time.Duration) {
	t := DefaultTimeout
	if len(timeout) > 0 {
		t = timeout[0]
	}
	go func() {
		if app.Config().Prefork {
			pids = []int{}
			app.Hooks().OnFork(func(i int) error {
				pids = append(pids, i)
				return nil
			})
			_ = app.ListenTLS(addr, certFile, keyFile)
		} else {
			if IsReload() {
				cert, _ := tls.LoadX509KeyPair(certFile, keyFile)
				ln, _ := reuseport.Listen(app.Config().Network, addr)
				ln = tls.NewListener(ln, &tls.Config{
					MinVersion: tls.VersionTLS12,
					Certificates: []tls.Certificate{
						cert,
					},
				})
				_ = app.Listener(ln)
			} else {
				_ = app.ListenTLS(addr, certFile, keyFile)
			}
		}
	}()
	listenSig(app, t)
}

func listenSig(app *fiber.App, timeout time.Duration) {
	if fiber.IsChild() {
		stop := make(chan int, 1)
		go func() {
			c := make(chan os.Signal, 1) // Create channel to signify a signal being sent
			signal.Notify(c, syscall.SIGTERM)
			for {
				sig := <-c
				switch sig {
				case syscall.SIGTERM:
					_ = app.Shutdown()
					fwrite(func(file *os.File) {
						_, _ = file.WriteString(fmt.Sprintf("%v\n", os.Getpid()))
					})
				case syscall.SIGINT:
					stop <- 1
					return
				}
			}
		}()
		<-stop
	} else {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGTERM, syscall.SIGUSR2)
		for {
			sig := <-c
			flog.Info("SIGN:", sig)
			switch sig {
			case syscall.SIGTERM:

				if app.Config().Prefork {
					stopChild(timeout)
					_ = app.Shutdown()
				} else {
					stop(app, timeout)
				}
				flog.Info("Grace Shutdown Success!")
				return
			case syscall.SIGUSR2:
				f, _ := os.Create(FIBER_CHILD_LOCK_FILE)
				_ = f.Close()
				_ = reload()
				if app.Config().Prefork {
					stopChild(timeout)
					_ = app.Shutdown()
					for _, key := range pids {
						_ = syscall.Kill(key, syscall.SIGINT)
					}
				} else {
					stop(app, timeout)
				}
				log.Println("Grace Reload Success")
				if IsReload() {
					return
				}
			}
		}
	}
}
func stop(app *fiber.App, timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.TODO(), timeout)
	defer cancel()
	chErr := make(chan error, 1)
	go func() {
		chErr <- app.Shutdown()
	}()

	for {
		select {
		case <-chErr:
			return
		case <-ctx.Done():
			return
		}
	}

}
func stopChild(timeout time.Duration) {
	pidMap := make(map[int]int)
	for _, key := range pids {
		_ = syscall.Kill(key, syscall.SIGTERM)
		pidMap[key] = key
	}
	ctx, cancel := context.WithTimeout(context.TODO(), timeout)
	defer cancel()

	for {
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			return
		default:
			content, err := os.ReadFile(FIBER_CHILD_LOCK_FILE)
			if err != nil {
				break
			}
			sarr := strings.Split(string(content), "\n")

			sax := map[int]int{}
			for _, id := range sarr {
				if v, err := strconv.Atoi(id); err == nil {
					sax[v] = v
				}
			}

			okCount := 0
			for id := range sax {
				if _, ok := pidMap[id]; ok {
					okCount += 1
				}
			}

			if okCount == len(pids) {
				cancel()
				return
			}
		}
	}
}

func reload() error {
	cmd := exec.Command(os.Args[0], os.Args[1:]...) // #nosec G204
	log.Println(cmd.String())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = append(os.Environ(),
		fmt.Sprintf("%s=%s", envReload, envReloadVal),
	)
	return cmd.Start()
}

func fwrite(handler func(file *os.File)) {
	f, err := os.OpenFile(FIBER_CHILD_LOCK_FILE, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	defer func() { _ = f.Close() }()
	if err != nil {
		log.Println("create lock file failed", err)
		return
	}
	if err := syscall.Flock(int(f.Fd()), syscall.LOCK_SH|syscall.LOCK_NB); err != nil {
		log.Println("add share lock in no block failed", err)
		return
	}
	handler(f)
	if err := syscall.Flock(int(f.Fd()), syscall.LOCK_UN); err != nil {
		log.Println("unlock share lock failed", err)
	}
	return
}
