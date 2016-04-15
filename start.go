package androidam

import (
	"os/exec"
	"strconv"
)

// Start represents the start subcommand
func Start(opts *Opts, stArgs *StartArgs, intent *IntentArgs, args ...string) *exec.Cmd {
	var argv []string
	argv = append(argv, "start")
	argv = append(argv, stArgs.ToArgv()...)
	argv = append(argv, intent.ToArgv()...)
	argv = append(argv, args...)
	if opts != nil && opts.AMPath != "" {
		return exec.Command(opts.AMPath, argv...)
	}
	return exec.Command("am", argv...)
}

// StartArgs structures the arguments for the Start subcommand.
type StartArgs struct {
	Debugging      bool   // -D  // enable debugging
	Wait           bool   // -W  // wait for launch to complete
	ForceStopFirst bool   // -S // force stop target app before starting activity
	OpenGLTrace    bool   // --opengl-trace // enable tracing of opengl functions
	RepeatLaunch   int    // -R <value> // repeat the activity launch <COUNT> times.  Prior to each repeat the top activity will be finished.
	StartProfiler  string // --start-profiler <value>
	StartProfilerP string // -P <value> // Like above but stops when app goes idle
	UserID         string // --user <USER_ID> | current // specify which user to run as, default current
}

// ToArgv converts StartArgs to an argv slice
func (sa *StartArgs) ToArgv() []string {
	if sa == nil {
		return nil
	}
	var argv []string
	if sa.Debugging {
		argv = append(argv, "-D")
	}
	if sa.Wait {
		argv = append(argv, "-W")
	}
	if sa.ForceStopFirst {
		argv = append(argv, "-S")
	}
	if sa.OpenGLTrace {
		argv = append(argv, "--opengl-trace")
	}
	if sa.RepeatLaunch != 0 {
		argv = append(argv, "-R")
		argv = append(argv, strconv.Itoa(sa.RepeatLaunch))
	}
	if sa.StartProfiler != "" {
		argv = append(argv, "--start-profiler")
		argv = append(argv, sa.StartProfiler)
	}
	if sa.StartProfilerP != "" {
		argv = append(argv, "-P")
		argv = append(argv, sa.StartProfilerP)
	}
	if sa.StartProfilerP != "" {
		argv = append(argv, "--user")
		argv = append(argv, sa.UserID)
	}
	return argv
}
