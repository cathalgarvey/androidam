# Android-AM Wrapper for Go
by Cathal Garvey, Copyright 2016, Released under GNU AGPLv3 or later.

## The Poor Person's Android API for Go
[![GoDoc][godoc-badge]][godoc]
The android `am` command can be used to send android intents and trigger all
kinds of useful behaviours, with the only real drawback being that it's largely
undocumented and the intents system is a java-esque mire of inconsistent
namespacing.

All this notwithstanding, while Go lacks native access to Android APIs, being
able to shell out to `am` from Go with type safety checks would be a significant
boon to many projects.

In particular, I wanted to create a Go wrapper for Termux's API tools, and Termux
itself wraps and uses `am` to communicate from its terminal to its API app.

This appears to work for the two subcommands I care about: `broadcast` and `start`.

Pull requests to fix or add functionality from actual Android devs who understand
`am` and intents are very welcome!

[godoc]: https://godoc.org/github.com/cathalgarvey/androidam "GoDoc"
[godoc-badge]: https://godoc.org/github.com/cathalgarvey/androidam?status.svg "GoDoc Badge"
