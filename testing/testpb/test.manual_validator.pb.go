package testpb

import (
	"math"
)

func (x *PingRequest) Validate(bool) error { _ = "STUB: not implemented"; return nil }

func (x *PingErrorRequest) Validate() error { _ = "STUB: not implemented"; return nil }

func (x *PingListRequest) Validate(bool) error { _ = "STUB: not implemented"; return nil }

func (x *PingStreamRequest) Validate(bool) error { _ = "STUB: not implemented"; return nil }

func (x *PingResponse) Validate() error { _ = "STUB: not implemented"; return nil }

func (x *PingResponse) ValidateAll() error { _ = "STUB: not implemented"; return nil }

var (
	GoodPing       = &PingRequest{Value: "something", SleepTimeMs: 9999}
	GoodPingError  = &PingErrorRequest{Value: "something", SleepTimeMs: 9999}
	GoodPingList   = &PingListRequest{Value: "something", SleepTimeMs: 9999}
	GoodPingStream = &PingStreamRequest{Value: "something", SleepTimeMs: 9999}

	BadPing       = &PingRequest{Value: "something", SleepTimeMs: 10001}
	BadPingError  = &PingErrorRequest{Value: "something", SleepTimeMs: 10001}
	BadPingList   = &PingListRequest{Value: "something", SleepTimeMs: 10001}
	BadPingStream = &PingStreamRequest{Value: "something", SleepTimeMs: 10001}

	GoodPingResponse = &PingResponse{Counter: 100}
	BadPingResponse  = &PingResponse{Counter: math.MaxInt16 + 1}
)
