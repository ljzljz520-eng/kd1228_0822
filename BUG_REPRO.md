# BUG_REPRO

The following failures were observed while validating the initial project state.
Each section records what failed, how to reproduce it, and the complete command output.
They are preserved intentionally; only failing build gates are omitted from the generated Dockerfile.

## Failure 1: Go test (.)

- Observed problem: `Go test (.)` failed in the initial project state.
- Working directory: `.`
- Command: `cd /app && GOTOOLCHAIN=local GOPROXY=off GOSUMDB=off go test -count=1 ./...`
- Exit status: `1`

```text
--- FAIL: TestJellyfishSpeedChangesStayLocal (0.01s)
    regression_test.go:27: speed for jelly-c changed from 0.80 to 3.70
FAIL
FAIL	jellyfield	0.062s
ok  	jellyfield/analysis	0.003s
ok  	jellyfield/app	0.016s
ok  	jellyfield/catalog	0.001s
?   	jellyfield/cmd/jellyfield	[no test files]
ok  	jellyfield/control	0.002s
ok  	jellyfield/exporter	0.002s
?   	jellyfield/geometry	[no test files]
ok  	jellyfield/gesture	0.002s
?   	jellyfield/input	[no test files]
ok  	jellyfield/layout	0.002s
ok  	jellyfield/model	0.002s
ok  	jellyfield/persistence	0.010s
?   	jellyfield/presentation	[no test files]
ok  	jellyfield/protocol	0.002s
ok  	jellyfield/replay	0.002s
ok  	jellyfield/simulation	0.001s
ok  	jellyfield/timeline	0.001s
FAIL
```

## Architecture reproduction

### linux/amd64
- Go toolchain version: exit `0`
- Node.js version: exit `0`
- Go build (.): exit `0`
- Go test (.): exit `1`
- Go run smoke (cmd/jellyfield): exit `0`
- Frontend build (web): exit `0`
### linux/arm64
- Go toolchain version: exit `0`
- Node.js version: exit `0`
- Go build (.): exit `0`
- Go test (.): exit `1`
- Go run smoke (cmd/jellyfield): exit `0`
- Frontend build (web): exit `0`
