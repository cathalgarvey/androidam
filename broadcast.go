package androidam

import "os/exec"

// Broadcast execs `am broadcast` with the provided arguments.
func Broadcast(opts *Opts, bcArgs *BroadcastArgs, intent *IntentArgs, args ...string) *exec.Cmd {
	var argv []string
	argv = append(argv, "broadcast")
	argv = append(argv, bcArgs.ToArgv()...)
	argv = append(argv, intent.ToArgv()...)
	argv = append(argv, args...)
	if opts != nil && opts.AMPath != "" {
		return exec.Command(opts.AMPath, argv...)
	}
	return exec.Command("am", argv...)
}

// BroadcastArgs holds common arguments for the Broadcast subcommand.
type BroadcastArgs struct {
	User string // --user <USER_ID> | all | current
}

// ToArgv returns commands suitable for appending to the argv string when exec'ing am.
func (bca *BroadcastArgs) ToArgv() []string {
	if bca == nil {
		return nil
	}
	if bca.User != "" {
		return []string{"--user", bca.User}
	}
	return nil
}
