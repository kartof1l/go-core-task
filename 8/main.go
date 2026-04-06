package main

// вся суть в канале -  при добавлении счетчик увеличивается уходя в канал, при завершении читается из канала
type WaitGroup struct {
	sem chan struct{}
}

func NewWaitGroup() *WaitGroup {
	return &WaitGroup{
		sem: make(chan struct{}),
	}
}

func (wg *WaitGroup) Add(delta int) {
	for i := 0; i < delta; i++ {
		wg.sem <- struct{}{}//чтобы не захламлять 
	}
}

func (wg *WaitGroup) Done() {
	<-wg.sem
}

func (wg *WaitGroup) Wait() {//ждет пока канал не опустошится
	for len(wg.sem) > 0 {
	}
}