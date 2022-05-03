.PHONY: generate
generate:
	docker run \
      --volume "$(pwd):/workspace" \
      --workdir /workspace \
      bufbuild/buf generate
