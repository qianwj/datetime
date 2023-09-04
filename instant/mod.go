package instant

import (
	"github.com/qianwj/datetime"
	"time"
)

var (
	EPOCH = &Instant{seconds: 0, nanos: 0}
)

/**
 * The minimum supported epoch second.
 */
const minSecond int64 = -31557014167219200

/**
 * The maximum supported epoch second.
 */
const maxSecond int64 = 31556889864403199

type Instant struct {
	seconds int64
	nanos   int
}

func (i *Instant) GetEpochSecond() int64 {
	return i.seconds
}

func (i *Instant) GetNano() int {
	return i.nanos
}

func Now() *Instant {
	now := time.Now()
	return &Instant{seconds: now.Unix(), nanos: now.Nanosecond()}
}

func OfEpochSecond(epochSecond int64) (*Instant, error) {
	return create(epochSecond, 0)
}

func OfEpochMilli(epochMilli int64) (*Instant, error) {
	secs := floorDiv(epochMilli, 1000)
	mos := floorMod(epochMilli, 1000)
	return create(secs, mos*1000_000)
}

func create(seconds int64, nanoOfSecond int) (*Instant, error) {
	if (seconds | int64(nanoOfSecond)) == 0 {
		return EPOCH, nil
	}
	if seconds < minSecond || seconds > maxSecond {
		return nil, datetime.NewDateTimeError("Instant exceeds minimum or maximum instant")
	}
	return &Instant{seconds: seconds, nanos: nanoOfSecond}, nil
}

func floorDiv(x, y int64) int64 {
	r := x / y
	// if the signs are different and modulo not zero, round down
	if (x^y) < 0 && (r*y != x) {
		r--
	}
	return r
}

func floorMod(x int64, y int) int {
	// Result cannot overflow the range of int.
	return int(x - floorDiv(x, int64(y))*int64(y))
}
