run-test:
	go test ./... -cover
publish-pkg:
	chmod +x publish-go-pkg.sh && ./publish-go-pkg.sh -v $(version) -r $(repository_name)