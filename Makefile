.PHONY: all format check todo

all: format check todo

format:
	@sh scripts/format.sh

check:
	@sh scripts/check.sh

todo:
	@sh scripts/todo.sh
