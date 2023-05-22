module github.com/bankApi

require (
	github.com/bank v0.0.1
)

replace github.com/bank => ../bankCore

go 1.20
