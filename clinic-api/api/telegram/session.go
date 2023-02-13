package telegram

import (
	"clinic-api/service"
	tgBot "github.com/Syfaro/telegram-bot-api"
	"time"
)

var (
	endSessionCh     = make(chan int64)
	switchSessionsCh = make(chan session)
)

const (
	sessionTime = time.Minute * 10
)

type sessions struct {
	active []session
	*service.Service
}

func newSessions(service *service.Service) *sessions {
	s := sessions{
		active:  []session{},
		Service: service,
	}

	go func() {
		for {
			select {
			case id := <-endSessionCh:
				for i := range s.active {
					if id == s.active[i].id() {
						s.active[i] = s.active[len(s.active)-1]
						s.active = s.active[:len(s.active)-1]
						i--
					}
				}
			case session := <-switchSessionsCh:
				for i := range s.active {
					if session.id() == s.active[i].id() {
						s.active[i] = session
					}
				}
			}
		}
	}()

	return &s
}

func (s *sessions) check(id int64) session {
	defer func() {
		recover()
	}()
	for _, s := range s.active {
		if s.id() == id {
			s.timerReset(sessionTime)
			return s
		}
	}
	return nil
}

func switchSessions(s session) {
	switchSessionsCh <- s
}

func endSession(id int64) {
	endSessionCh <- id
}

type session interface {
	exec(u *tgBot.Update)
	id() int64
	timerSet(d time.Duration)
	timerReset(d time.Duration)
}

type base struct {
	ID    int64
	lvl   int
	timer *time.Timer
	*service.Service
	api *tgBot.BotAPI
}

func (b *base) id() int64 {
	return b.ID
}

func (b *base) timerReset(d time.Duration) {
	b.timer.Reset(d)
}

func (b *base) timerSet(d time.Duration) {
	b.timer = time.AfterFunc(d, func() {
		endSession(b.ID)
	})
}
