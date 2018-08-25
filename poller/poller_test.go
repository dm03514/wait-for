package poller

import (
	"fmt"
	"testing"
	"time"
)

func TestPoller_WaitReady_FailFastError(t *testing.T) {
	p := Poller{
		Timeout:  time.Millisecond * 1,
		Interval: time.Millisecond * 1,
		FailFast: true,
		CheckReady: func() (ready bool, err error) {
			return false, fmt.Errorf("fail fast error")
		},
	}
	err := p.WaitReady()
	if err == nil && err.Error() != "fail fast error" {
		t.Errorf("expected error, received %s", err)
	}
}

func TestPoller_WaitReady_CheckReadySuccess(t *testing.T) {
	p := Poller{
		Timeout:  time.Millisecond * 1,
		Interval: time.Millisecond * 1,
		FailFast: true,
		CheckReady: func() (ready bool, err error) {
			return true, nil
		},
	}
	if err := p.WaitReady(); err != nil {
		t.Errorf("expected no error: received: %s", err)
	}
}

func TestPoller_WaitReady_CheckReadyTimeout(t *testing.T) {
	p := Poller{
		Timeout:  time.Millisecond * 1,
		Interval: time.Millisecond * 1,
		FailFast: true,
		CheckReady: func() (ready bool, err error) {
			return false, nil
		},
	}

	err := p.WaitReady()
	if err == nil && err.Error() != "timeout reached" {
		t.Errorf("expected error, received %s", err)
	}
}
