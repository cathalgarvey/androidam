package androidam

import "strconv"

// IntentArgs is all the flags that correspond to the <INTENT> portion of am's
// command spec. Keys not set will be omitted from the output when converted
// to a string-slice like an argv spec using the ToArgv method.
type IntentArgs struct {
	URI                         string            // Positional after flags
	Action                      string            // -a <ACTION>
	DataURI                     string            // -d <DATA_URI>
	MimeType                    string            // -t <MIME_TYPE>
	Category                    []string          // [-c <CATEGORY> [-c <CATEGORY>] ...]
	ExtraKeyStrings             map[string]string // [-e|--es <EXTRA_KEY> <EXTRA_STRING_VALUE> ...]
	ExtraKeyFlags               []string          // [--esn <EXTRA_KEY> ...]
	ExtraKeyBools               map[string]bool   // [--ez <EXTRA_KEY> <EXTRA_BOOLEAN_VALUE> ...]
	ExtraKeyInts                map[string]int    // [-e|--ei <EXTRA_KEY> <EXTRA_INT_VALUE> ...]
	Component                   string            // -n <COMPONENT>
	Flags                       string            // -f <FLAGS>
	GrantReadURIPermission      bool              // --grant-read-uri-permission
	GrantWriteURIPermission     bool              // --grant-write-uri-permission
	DebugLogResolution          bool              // --debug-log-resolution
	ActivityBroughtToFront      bool              // --activity-brought-to-front
	ActivityClearTop            bool              // --activity-clear-top
	ActivityClearWhenTaskReset  bool              // --activity-clear-when-task-reset
	ActivityExcludeFromRecents  bool              // --activity-exclude-from-recents
	ActivityLaunchedFromHistory bool              // --activity-launched-from-history
	ActivityMultipleTask        bool              // --activity-multiple-task
	ActivityNoAnimation         bool              // --activity-no-animation
	ActivityNoHistory           bool              // --activity-no-history
	ActivityNoUserAction        bool              // --activity-no-user-action
	ActivityPreviousIsTop       bool              // --activity-previous-is-top
	ActivityReorderToFront      bool              // --activity-reorder-to-front
	ActivityResetTaskIfNeeded   bool              // --activity-reset-task-if-needed
	ActivitySingleTop           bool              // --activity-single-top
	ReceiverRegisteredOnly      bool              // --receiver-registered-only
	ReceiverReplacePending      bool              // --receiver-replace-pending
}

// ToArgv converts an IntentArgs into a collection of key-value argstrings in order.
func (i *IntentArgs) ToArgv() (argv []string) {
	apparg := func(s string) {
		argv = append(argv, s)
	}
	appIfUnempty := func(flag, value string) {
		if value != "" {
			apparg(flag)
			apparg(value)
		}
	}
	appIfTrue := func(flag string, val bool) {
		if val {
			apparg(flag)
		}
	}
	appAllSlice := func(flag string, sl []string) {
		for _, v := range sl {
			apparg(flag)
			apparg(v)
		}
	}
	appIfUnempty("-a", i.Action)
	appIfUnempty("-d", i.DataURI)
	appIfUnempty("-t", i.MimeType)
	appIfUnempty("-n", i.Component)
	appIfUnempty("-f", i.Flags)
	appAllSlice("-c", i.Category)
	appAllSlice("--esn", i.ExtraKeyFlags)
	for k, v := range i.ExtraKeyBools {
		apparg("--ez")
		apparg(k)
		if v {
			apparg("true")
		} else {
			apparg("false")
		}
	}
	for k, v := range i.ExtraKeyStrings {
		apparg("--es")
		apparg(k)
		apparg(v)
	}
	for k, v := range i.ExtraKeyInts {
		apparg("--ei")
		apparg(k)
		apparg(strconv.Itoa(v))
	}
	appIfTrue("--grant-read-uri-permission", i.GrantReadURIPermission)
	appIfTrue("--grant-write-uri-permission", i.GrantWriteURIPermission)
	appIfTrue("--debug-log-resolution", i.DebugLogResolution)
	appIfTrue("--activity-brought-to-front", i.ActivityBroughtToFront)
	appIfTrue("--activity-clear-top", i.ActivityClearTop)
	appIfTrue("--activity-clear-when-task-reset", i.ActivityClearWhenTaskReset)
	appIfTrue("--activity-exclude-from-recents", i.ActivityExcludeFromRecents)
	appIfTrue("--activity-launched-from-history", i.ActivityLaunchedFromHistory)
	appIfTrue("--activity-multiple-task", i.ActivityMultipleTask)
	appIfTrue("--activity-no-animation", i.ActivityNoAnimation)
	appIfTrue("--activity-no-history", i.ActivityNoHistory)
	appIfTrue("--activity-no-user-action", i.ActivityNoUserAction)
	appIfTrue("--activity-previous-is-top", i.ActivityPreviousIsTop)
	appIfTrue("--activity-reorder-to-front", i.ActivityReorderToFront)
	appIfTrue("--activity-reset-task-if-needed", i.ActivityResetTaskIfNeeded)
	appIfTrue("--activity-single-top", i.ActivitySingleTop)
	appIfTrue("--receiver-registered-only", i.ReceiverRegisteredOnly)
	appIfTrue("--receiver-replace-pending", i.ReceiverReplacePending)
	apparg(i.URI)
	return argv
}
