CURDIR=$(shell pwd)
PPROFDIR=${CURDIR}/pprof

.PHONY: traceprof
traceprof:
	curl http://localhost:8080/debug/pprof/trace\?seconds\=20 -o trace.out
	go tool trace -http ":8082" trace.out

.PHONY: cpuprof
cpuprof:
	(PPROF_TMPDIR=${PPROFDIR} go tool pprof -http :8082 -seconds 20 http://127.0.0.1:8080/debug/pprof/profile)
