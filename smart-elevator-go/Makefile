SERVICE_NAME=elevatorapp
PROJECTS=$(SERVICE_NAME)/...

all:
	@cd src/ && go install $(PROJECTS)

test:
	@cd src/ && go test -coverprofile=coverage.out $(PROJECTS)
	@go tool cover -html=src/coverage.out

# generate XML report in Cobertura format
test.xml:
	@cd src/ && gocov test $(PROJECTS) | gocov-xml > coverage.xml

clean:
	@rm -rf ./bin/$(SERVICE_NAME)

clean-all:

run:
	@./bin/testapp

stop-all:
	@pkill -f "./bin/testapp" &


run-testapp-notifications:
	@./bin/testapp
stop-testapp-notifications:
	@pkill -f "./bin/testapp"