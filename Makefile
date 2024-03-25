.SILENT:
.DEFAULT_GOAL := demo-run

.PHONY: demo-run
demo-run:
	go run ./cmd/demo
