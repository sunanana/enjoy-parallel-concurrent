.PHONY: run
run:
	@if [ ! -e ./vendor/bundle/ruby ]; then \
		make build; \
	fi
	@bundle exec ruby parallel.rb

.PHONY: build
build:
	@bundle install

.PHONY: clean
clean:
	@rm -rf ./vendor
