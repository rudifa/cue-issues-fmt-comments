# Define Go command and flags
GO = go
GOFLAGS = -ldflags="-s -w"

# Define the target executable
TARGET = cuedo

# Default target: build the executable
all: $(TARGET)

# Rule to build the target executable
$(TARGET): main.go
	$(GO) build $(GOFLAGS) -o $(TARGET) main.go

# Clean target: remove the target executable
clean:
	rm -f $(TARGET)

# Run target: build and run the target executable
run: $(TARGET)
	./$(TARGET)

# Test target: run Go tests for the project
test:
	$(GO) test ./...

# Test target: run Go tests for the project with verbose output
vtest:
	$(GO) test -v ./...


.PHONY: build install

GIT_REV_PARSE=git rev-parse HEAD
GIT_SHOW_DATE=git show -s --format=%cI

SET_VARS=date_now=$$(date -Iseconds); \
	user=$$(whoami); \
	cuedo_sha=$$($(GIT_REV_PARSE)); \
	cuedo_date=$$($(GIT_SHOW_DATE) $$cuedo_sha); \
	replace_path=$$(awk '/replace cuelang.org\/go => / {print $$4}' go.mod); \
	if [ -n "$$replace_path" ]; then \
		cd $$replace_path; \
		cue_sha=$$($(GIT_REV_PARSE)); \
		cue_date=$$($(GIT_SHOW_DATE) $$cue_sha); \
		cd -; \
	else \
		cue_sha=$$(go list -m -json cuelang.org/go | jq -r .Version); \
		cue_date=$$(go list -m -json cuelang.org/go | jq -r .Time); \
	fi;

# WARNING the else branc above is not tested yet

LD_FLAGS="-X github.com/rudifa/cuedo/cmd.cueCommitSHA=$$cue_sha \
-X github.com/rudifa/cuedo/cmd.cueCommitDate=$$cue_date \
-X github.com/rudifa/cuedo/cmd.cuedoCommitSHA=$$cuedo_sha \
-X github.com/rudifa/cuedo/cmd.cuedoCommitDate=$$cuedo_date \
-X github.com/rudifa/cuedo/cmd.DateNow=$$date_now \
-X github.com/rudifa/cuedo/cmd.User=$$user"

build: $(TARGET)
	$(SET_VARS) \
	go build -v -ldflags $(LD_FLAGS)

install: $(TARGET)
	$(SET_VARS) \
	go install -v -ldflags $(LD_FLAGS)
