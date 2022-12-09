package xoj

import (
	"JudgeHost/xoj/dao"
	"sync"
)

const queueMaxLength = 20

type WaitQueue struct {
	waitArray []*dao.JudgeStatus
	start     int
	end       int
	mu        sync.Mutex
}

func NewWaitQueue() *WaitQueue {
	return &WaitQueue{waitArray: make([]*dao.JudgeStatus, queueMaxLength)}
}

func (w *WaitQueue) push(p *dao.JudgeStatus) {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.end < queueMaxLength {
		w.waitArray[w.end] = p
		w.end = w.end + 1
	}
}

func (w *WaitQueue) front() (res *dao.JudgeStatus) {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.empty() {
		return nil
	}
	res = w.waitArray[w.start]
	w.waitArray[w.start] = nil
	w.start = w.start + 1
	if w.empty() && w.end == queueMaxLength {
		w.start = 0
		w.end = 0
	}
	return
}

func (w *WaitQueue) empty() bool {
	return w.start == w.end
}
