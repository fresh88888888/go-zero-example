package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/zeromicro/go-zero/core/breaker"
	"github.com/zeromicro/go-zero/core/lang"
	"github.com/zeromicro/go-zero/core/logx"
	"gopkg.in/cheggaaa/pb.v1"
)

const (
	duration        = time.Minute * 5
	breakRange      = 20
	workRange       = 20
	requestInterval = time.Microsecond
	// mulitply to make it visible in plot
	stateFator = float64(time.Second/requestInterval) / 2
)

type (
	server struct {
		stat int32
	}

	metric struct {
		calls int64
	}
)

func (m *metric) addCall() {
	atomic.AddInt64(&m.calls, 1)
}

func (m *metric) reset() int64 {
	return atomic.SwapInt64(&m.calls, 0)
}

func NewServer() *server {
	return &server{}
}

func (s *server) serve(m *metric) bool {
	m.addCall()
	return atomic.LoadInt32(&s.stat) == 1
}

func (s *server) start() {
	go func() {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		var state int32
		for {
			var v int32
			if state == 0 {
				v = r.Int31n(breakRange)
			} else {
				v = r.Int31n(workRange)
			}
			time.Sleep(time.Second * time.Duration(v+1))
			state ^= 1
			atomic.StoreInt32(&s.stat, state)
		}
	}()
}

func runBreaker(s *server, br breaker.Breaker, duration time.Duration, m *metric) {
	ticker := time.NewTicker(requestInterval)
	defer ticker.Stop()
	done := make(chan lang.PlaceholderType)

	go func() {
		time.Sleep(duration)
		close(done)
	}()

	for {
		select {
		case <-ticker.C:
			_ = br.Do(func() error {
				if s.serve(m) {
					return nil
				} else {
					return breaker.ErrServiceUnavailable
				}
			})
		case <-done:
			return
		}
	}
}
func main() {
	serv := NewServer()
	serv.start()

	gb := breaker.NewBreaker()
	fp, err := os.Create("result.csv")
	logx.Must(err)

	defer fp.Close()
	fmt.Fprintln(fp, "seconds,state,googleCalls,netflixCalls")

	var gm, nm metric
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		var seconds int
		for range ticker.C {
			seconds++
			gcalls := gm.reset()
			ncalls := nm.reset()
			fmt.Fprintf(fp, "%d,%.2f,%d,%d\n", seconds, float64(atomic.LoadInt32(&serv.stat))*stateFator, gcalls, ncalls)
		}
	}()

	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	go func() {
		runBreaker(serv, gb, duration, &gm)
		waitGroup.Done()
	}()

	go func() {
		bar := pb.New(int(duration / time.Second)).Start()
		runBreaker(serv, gb, duration, &gm)
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for range ticker.C {
			bar.Increment()
		}
		bar.Finish()
	}()

	waitGroup.Wait()
}
