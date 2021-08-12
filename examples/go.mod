module github.com/lonelyevil/khl/examples

go 1.16

require (
	github.com/bits-and-blooms/bloom/v3 v3.0.1 // indirect
	github.com/lonelyevil/khl v0.0.1
	github.com/lonelyevil/khl/log_adapter/plog v0.0.1
	github.com/phuslu/log v1.0.71
)

replace (
	github.com/lonelyevil/khl => ../.
	github.com/lonelyevil/khl/log_adapter/plog => ../log_adapter/plog
)
