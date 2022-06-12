ROOT_DIR = $(dir $(realpath $(firstword $(MAKEFILE_LIST))))

.PHONY: clean
clean:
	rm -rf build

.PHONY: build
build: clean
	docker exec tmp_project_name_go make build && \
	docker exec tmp_project_name_node npm install && \
	docker exec tmp_project_name_node npm run build && \
	cp -r $(ROOT_DIR)golang/build $(ROOT_DIR)/ && \
	cp -r $(ROOT_DIR)node/dist/tmp_project_name $(ROOT_DIR)/build/resources/frontend
