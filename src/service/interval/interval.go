package interval

import (
	"context"
	"time"
)

func NewIntervalData[T any](data Data[T]) IntervalData[T] {
	intervalData := intervalData[T]{
		data: data,
	}

	return intervalData
}

type Data[T any] interface {
	Get() (T, error)
}

type IntervalData[T any] interface {
	Start(ctx context.Context, duration time.Duration) (chan T, chan error)
	SetInterval(duration time.Duration)
}

type intervalData[T any] struct {
	intervalTicker *time.Ticker
	data           Data[T]
}

func (s intervalData[T]) Start(ctx context.Context, duration time.Duration) (chan T, chan error) {

	chData := make(chan T, 1)
	chErr := make(chan error, 1)

	//handle case where context is already Done()
	select {
	case <-ctx.Done():
		close(chData)
		close(chErr)
		return chData, chErr
	default:
		//immediately populate some data
		data, err := s.data.Get()
		if err != nil {
			chErr <- err
		} else {
			chData <- data
		}
	}

	s.intervalTicker = time.NewTicker(duration)

	go func() {
		defer s.intervalTicker.Stop()

	LOOP:
		for {
			select {
			case <-ctx.Done():
				break LOOP
			case <-s.intervalTicker.C:
				data, err := s.data.Get()
				if err != nil {
					chErr <- err
				} else {
					chData <- data
				}
			}
		}

		close(chData)
		close(chErr)
	}()

	return chData, chErr
}

func (s intervalData[T]) SetInterval(duration time.Duration) {
	s.intervalTicker.Reset(duration)
}
