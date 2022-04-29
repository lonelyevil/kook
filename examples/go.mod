module github.com/lonelyevil/khl/examples

go 1.16

require (
	github.com/lonelyevil/khl v0.0.25
	github.com/lonelyevil/khl/log_adapter/plog v0.0.25
	github.com/phuslu/log v1.0.77
)

replace (
	github.com/lonelyevil/khl => ../.
	github.com/lonelyevil/khl/log_adapter/plog => ../log_adapter/plog
)
