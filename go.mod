module Overseer

go 1.17

replace example.com/logMonitor => ../Overseer/logMonitor

require (
	example.com/commands v0.0.0-00010101000000-000000000000
	example.com/logMonitor v0.0.0-00010101000000-000000000000
	example.com/sha256checksums v0.0.0-00010101000000-000000000000
)

replace example.com/commands => ../Overseer/commands

replace example.com/sha256checksums => ../Overseer/sha256checksums
