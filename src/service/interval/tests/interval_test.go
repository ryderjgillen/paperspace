package tests

import (
	"context"
	"fmt"
	"portService/server/interval"
	"testing"
	"time"
)

type mockIntervalData struct {
	data mockData
	err  error
}

type mockData struct {
	prop1 string
}

func (m mockIntervalData) Get() (mockData, error) {
	return m.data, m.err
}

func Test_IntervalData(t *testing.T) {

	//given
	ctx := context.Background()
	var data interval.Data[mockData] = mockIntervalData{
		data: mockData{
			prop1: "val1",
		},
	}
	intervalData := interval.NewIntervalData(data)

	//when
	chData, chErr := intervalData.Start(ctx, 1*time.Second)

	//then
	select {
	case <-ctx.Done():
		return
	case data := <-chData:
		expected := "val1"
		if data.prop1 != "val1" {
			t.Fatalf("expected: %s, got: %s", expected, data.prop1)
		}
	case err := <-chErr:
		t.Fatal(err.Error())
	}
}

func Test_IntervalData_Cancel(t *testing.T) {

	//given
	var data interval.Data[mockData] = mockIntervalData{}
	intervalData := interval.NewIntervalData(data)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	//when
	cancel()
	chData, chErr := intervalData.Start(ctx, 1*time.Second)

	//then
	select {
	case <-ctx.Done():
		return
	case data := <-chData:
		t.Fatalf("expected: ctx.Done(), got: %v", data)
	case err := <-chErr:
		t.Fatal(err.Error())
	}
}

func Test_IntervalData_Error(t *testing.T) {

	//given
	var data interval.Data[mockData] = mockIntervalData{
		err: fmt.Errorf("mock error"),
	}
	intervalData := interval.NewIntervalData(data)
	ctx := context.Background()

	//when
	chData, chErr := intervalData.Start(ctx, 1*time.Second)

	//then
	select {
	case <-ctx.Done():
		t.Fatal("expected: error, got: ctx.Done()")
	case data := <-chData:
		t.Fatalf("expected: error, got: %v", data)
	case _ = <-chErr:
	}
}
