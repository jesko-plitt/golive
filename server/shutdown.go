package server

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/ao-concepts/logging"
	"github.com/gofiber/fiber/v2"
)

// ShutdownFunc a function that is executed when the server is shut down
type ShutdownFunc func() error

// ShutdownService handles shutdown functions
type ShutdownService struct {
	shutdownFuncs []ShutdownFunc
	log           logging.Logger
	lock          *sync.Mutex
}

func ProvideShutdownService(log logging.Logger) *ShutdownService {
	return &ShutdownService{
		log:  log,
		lock: &sync.Mutex{},
	}
}

func (s *ShutdownService) OnShutdown(fn ShutdownFunc) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.shutdownFuncs = append(s.shutdownFuncs, fn)
}

func (s *ShutdownService) ListenForShutdown(router *fiber.App, wg *sync.WaitGroup) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)

	go func() {
		_ = <-ch
		s.log.Info("server: shutting down...")
		_ = router.Shutdown()

		s.lock.Lock()
		defer s.lock.Unlock()

		for _, fn := range s.shutdownFuncs {
			if err := fn(); err != nil {
				s.log.ErrError(err)
			}
		}

		s.log.Info("server: shutdown finished")
		wg.Done()
	}()
}
