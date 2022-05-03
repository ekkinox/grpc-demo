.PHONY: gen-proto
gen-proto:
	docker run \
      --volume "$(pwd):/workspace" \
      --workdir /workspace \
      bufbuild/buf generate