package usecases

import "time"

type UseCase struct{}

func (uc UseCase) Save() {
	time.Sleep(time.Millisecond * 200)
}
