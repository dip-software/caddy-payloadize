# Makefile

OPENAPI_SPEC=../tenant-mapper/gen/http/openapi3.yaml
OUTPUT_DIR=./client/tenantmapper
PACKAGE_NAME=tenantmapper

.PHONY: gen

gen:
	@echo "Generating Go client from OpenAPI spec..."
	@openapi-generator generate -i $(OPENAPI_SPEC) -g go -o $(OUTPUT_DIR) \
		-p packageName=$(PACKAGE_NAME) \
		--additional-properties=withGoMod=false \
		--skip-validate-spec
