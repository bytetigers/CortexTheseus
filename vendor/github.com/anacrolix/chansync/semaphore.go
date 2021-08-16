package chansync

// Channel semaphore, as is popular for controlling access to limited resources. Should not be used
// zero-initialized.
type Semaphore struct {
	ch chan struct{}
}

// Returns an initialized semaphore with n slots.
func NewSemaphore(n int) Semaphore {
	return Semaphore{ch: make(chan struct{}, n)}
}

// Returns a channel for acquiring a slot.
func (me Semaphore) Acquire() Acquire {
	return me.ch
}

// Returns a channel for releasing a slot.
func (me Semaphore) Release() Release {
	return me.ch
}
