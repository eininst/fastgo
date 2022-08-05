package grace

import (
	"bytes"
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

type Grace struct {
	Entry    string
	Server   *http.Server
	Listener net.Listener
	Graceful *bool
	Timeout  time.Duration
}

type Option struct {
	timeout time.Duration
}

func cmd(s string) (string, error) {
	//"go build -o run run_api_grace.go"
	cmd := exec.Command("/bin/bash", "-c", s)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		return "", err
	}
	err = cmd.Wait()

	r := out.String()
	if len(r) > 0 {
		log.Println(r)
	}
	return r, err
}

func (grace *Grace) Reload() error {
	_, er := cmd(fmt.Sprintf("cd /cmd && go build -o run %v", grace.Entry))
	if er != nil {
		log.Fatal(er)
	}

	_, er = cmd("mv -f /cmd/run /app/run")
	if er != nil {
		log.Fatal(er)
	}

	tl, ok := grace.Listener.(*net.TCPListener)
	if !ok {
		return errors.New("listener is not tcp listener")
	}

	f, err := tl.File()
	if err != nil {
		return err
	}

	args := []string{"-graceful=true"}
	cmd := exec.Command(os.Args[0], args...)
	log.Println(cmd.String())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// put socket FD at the first entry
	cmd.ExtraFiles = []*os.File{f}
	return cmd.Start()
}
func (grace *Grace) Listening() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGUSR2)
	for {
		sig := <-ch
		log.Printf("signal: %v", sig)
		ctx, cancel := context.WithTimeout(context.TODO(), grace.Timeout)
		defer cancel()
		switch sig {
		case syscall.SIGTERM:
			signal.Stop(ch)
			if err := grace.Server.Shutdown(ctx); err != nil {
				log.Fatal("Server Shutdown:", err)
			}
			log.Printf("graceful shutdown")
			return
		case syscall.SIGUSR2:
			log.Println("reload")
			err := grace.Reload()
			if err != nil {
				log.Fatalf("graceful restart error: %v", err)
			}
			if err := grace.Server.Shutdown(ctx); err != nil {
				log.Fatal("Server Shutdown:", err)
			}
			log.Printf("graceful reload success")
			if os.Getpid() != 1 {
				return
			}
		}

	}

}

func Run(entry string, srv *http.Server, timeout time.Duration) {
	graceful := flag.Bool("graceful", false, "listen on fd open 3 (internal use only)")
	flag.Parse()

	_, _ = cmd(fmt.Sprintf("echo '%v' > /tmp/pid", os.Getpid()))

	var listener net.Listener
	var err error
	if *graceful {
		log.Print("Listening to existing file descriptor 3.")
		// cmd.ExtraFiles: If non-nil, entry i becomes file descriptor 3+i.
		// when we put socket FD at the first entry, it will always be 3(0+3)
		f := os.NewFile(3, "")
		listener, err = net.FileListener(f)
	} else {
		log.Print("Listening on a new file descriptor.")
		listener, err = net.Listen("tcp", srv.Addr)
	}
	if err != nil {
		log.Fatalf("listener error: %v", err)
	}
	go func() {
		// server.Shutdown() stops Serve() immediately, thus server.Serve() should not be in main goroutine
		err = srv.Serve(listener)
		log.Printf("server.Serve err: %v\n", err)
	}()

	grace := &Grace{
		Entry:    entry,
		Server:   srv,
		Listener: listener,
		Timeout:  timeout,
	}
	grace.Listening()
	log.Println("Listening end")
}
