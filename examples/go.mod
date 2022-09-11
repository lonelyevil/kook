module github.com/lonelyevil/kook/examples

go 1.16

require (
	github.com/lonelyevil/kook v0.0.28
	github.com/lonelyevil/kook/log_adapter/plog v0.0.28
	github.com/phuslu/log v1.0.80
)

replace (
	github.com/lonelyevil/kook => ../.
	github.com/lonelyevil/kook/log_adapter/plog => ../log_adapter/plog
)
