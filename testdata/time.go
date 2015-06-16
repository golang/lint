// Test of time suffixes.

// Package foo ...
package foo

import (
	"flag"
	"time"
)

var rpcTimeoutMsec = flag.Duration("rpc_timeout", 100*time.Millisecond, "some flag") // MATCH /Msec.*\*time.Duration/

var timeoutSecs = 5 * time.Second // MATCH /Secs.*time.Duration/

var delay = flag.Duration("delay", 5, "wait impatiently") // MATCH /time.*expression.*"5".*constants/

func f(...time.Duration) {}

var (
	_ = time.Duration(15)                     // MATCH /time.*expression.*"15".*constants/
	_ = time.Duration(15) + time.Duration(30) // Not caught, but rare in real code; see comments in lintDurationUnits
	_ = []time.Duration{
		1,  // MATCH /time.*expression.*"1".*constants/
		10, // MATCH /time.*expression.*"10".*constants/
		1 * time.Nanosecond,
		10 * 10 * 10 * 10,
		time.Hour,
		0, // 0 is OK
	}
	abc = f(17) // MATCH /time.*expression.*"17".*constants/
)

const x = 10

func g() {
	time.Sleep(
		12, // MATCH /time.*expression.*"12".*constants/
	)
	time.Sleep(x) // Not caught; see comments in lintDurationUnits
	_ = f(1,      // MATCH /time.*expression.*"1".*constants/
		1e12) // MATCH /time.*expression.*"1e12".*constants/
}
