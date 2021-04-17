init:
	git init && for file in ./cmd/*; do go build $$file; done;