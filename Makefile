
.PHONY: test
test:
	go test -v ./... -cover



rf:
	git add .
	git commit -m "save"
	git push origin master

push:
	git push origin master