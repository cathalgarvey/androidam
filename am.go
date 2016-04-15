package androidam

import (
	"os"
	"strconv"
)

// Opts is a set of optional options that alter the behaviour of AM or this wrapper.
type Opts struct {
	// Override path to AM command
	AMPath string
}

// GetCurrentAndroidUserID returns the calculated uid for this user.
// Intended for use when setting the User property of BroadcastArgs.
func GetCurrentAndroidUserID() string {
	return strconv.Itoa(os.Getuid() / 100000)
}

/*
func buildAMStartActivity(D, W bool, intent AMIntent) *AMCommand {
// start an Activity: am start [-D] [-W] <INTENT>
//        -D: enable debugging
//        -W: wait for launch to complete
} //*/
/*
func buildAMStartService(intent AMIntent) {
 // start a Service: am startservice <INTENT>

} //*/
/*
func buildAMBroadcast(intent AMIntent) *AMCommand {
 // send a broadcast Intent: am broadcast <INTENT>

} //*/
/*
func buildAMStartInstrumentation(eargs map[string]string, r, w bool, p, component string) *AMCommand {
//    start an Instrumentation: am instrument [flags] <COMPONENT>
//      -r: print raw results (otherwise decode REPORT_KEY_STREAMRESULT)
//      -e <NAME> <VALUE>: set argument <NAME> to <VALUE>
//      -p <FILE>: write profiling data to <FILE>
//      -w: wait for instrumentation to finish before returning

} //*/
/*
func buildAMStartProfiling(process, file string) *AMCommand {
  // start profiling: am profile <PROCESS> start <FILE>
  // stop profiling: am profile <PROCESS> stop
} //*/
/*
func buildAMStopProfiling(process, file string) *AMCommand {
 // start profiling: am profile <PROCESS> start <FILE>
 // stop profiling: am profile <PROCESS> stop
} //*/
