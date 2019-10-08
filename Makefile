include Makefile.variables

## prefix before other make targets to run in your local dev environment
local: | quiet
	@$(eval DOCKRUN= )
	@mkdir -p tmp
	@touch tmp/dev_image_id
quiet: # this is silly but shuts up 'Nothing to be done for `local`'
	@:

prepare: tmp/dev_image_id
tmp/dev_image_id:
	@mkdir -p tmp
	@docker rmi -f ${DEV_IMAGE} > /dev/null 2>&1 || true
	@docker build -t ${DEV_IMAGE} -f Dockerfile.dev .
	@docker inspect -f "{{ .ID }}" ${DEV_IMAGE} > tmp/dev_image_id

.PHONY: format
format: prepare
	${DOCKRUN} bash /opt/format.sh

.PHONY: check
check: prepare
	${DOCKRUN} bash /opt/check.sh

.PHONY: todo
todo: prepare
	${DOCKRUN} bash /opt/todo.sh