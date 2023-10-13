TARGETS := ci build package test validate
.PHONY: $(TARGETS) $(TEST_TARGETS) clean

.dapper:
	@echo Downloading dapper
	@curl -sL https://releases.rancher.com/dapper/latest/dapper-`uname -s`-`uname -m` > .dapper.tmp
	@@chmod +x .dapper.tmp
	@./.dapper.tmp -v
	@mv .dapper.tmp .dapper

$(TARGETS): .dapper
	@./.dapper $@

clean:
	@./scripts/clean.sh
