package poller

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

// CheckReadyFn is configurable by the caller and will be executed on the
// poll interval
type CheckReadyFn func() (ready bool, err error)

// Poller executes CheckReady function on an interval against a timeout
type Poller struct {
	Timeout    time.Duration
	Interval   time.Duration
	FailFast   bool
	CheckReady CheckReadyFn
}

// WaitReady calls the configured CheckReady function until it returns
// ready or until the timeout is reached.  It optionally will fail fast
// on error.
func (p Poller) WaitReady() error {
	interval := time.Tick(p.Interval)
	timer := time.NewTimer(p.Timeout)

	for {
		log.WithFields(log.Fields{}).Info("polling")

		ready, err := p.CheckReady()

		log.WithFields(log.Fields{
			"ready": ready,
			"err":   err,
		}).Debug("poll_result")

		if p.FailFast && err != nil {
			return err
		}

		if ready {
			return nil
		}

		select {
		case <-interval:
			continue
		case <-timer.C:
			log.WithFields(log.Fields{}).Info("timeout_reached")
			return fmt.Errorf("timeout reached: %s", p.Timeout)
		}
	}
}
