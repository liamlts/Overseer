module Overseer

go 1.17

replace example.com/logMonitor => ../Overseer/logMonitor

require (
	example.com/commands v0.0.0-00010101000000-000000000000
	example.com/logMonitor v0.0.0-00010101000000-000000000000
	example.com/sha256checksums v0.0.0-00010101000000-000000000000
	github.com/urfave/cli v1.22.5
)

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.0-20190314233015-f79a8a8ca69d // indirect
	github.com/russross/blackfriday/v2 v2.0.1 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
)

replace example.com/commands => ../Overseer/commands

replace example.com/sha256checksums => ../Overseer/sha256checksums
