#!/bin/bash

# # Get the commit SHA and date for cue
# cue_sha=$(git --git-dir=path/to/cue/.git rev-parse HEAD)
# cue_date=$(git --git-dir=path/to/cue/.git show -s --format=%ci $cue_sha)

# # Get the commit SHA and date for cuedo
# cuedo_sha=$(git --git-dir=path/to/cuedo/.git rev-parse HEAD)
# cuedo_date=$(git --git-dir=path/to/cuedo/.git show -s --format=%ci $cuedo_sha)

# # Build the application
# go build -ldflags "-X main.cueCommitSHA=$cue_sha -X main.cueCommitDate='$cue_date' -X main.cuedoCommitSHA=$cuedo_sha -X main.cuedoCommitDate='$cuedo_date'" ./...

# local builds
# Get the commit SHA and date for cue
cue_sha=$(git --git-dir=$(go list -m -f '{{.Dir}}' cuelang.org/go)/.git rev-parse HEAD)
cue_date=$(git --git-dir=$(go list -m -f '{{.Dir}}' cuelang.org/go)/.git show -s --format=%cI $cue_sha)

# Get the commit SHA and date for cuedo
cuedo_sha=$(git rev-parse HEAD)
cuedo_date=$(git show -s --format=%cI $cuedo_sha)

echo "cue_sha: $cue_sha"
echo "cue_date: $cue_date"
echo "cuedo_sha: $cuedo_sha"
echo "cuedo_date: $cuedo_date"

set -x

# Build the application
go build -v -ldflags "-X github.com/rudifa/cuedo/cmd.cueCommitSHA=$cue_sha -X github.com/rudifa/cuedo/cmd.cueCommitDate=$cue_date -X github.com/rudifa/cuedo/cmd.cuedoCommitSHA=$cuedo_sha -X github.com/rudifa/cuedo/cmd.cuedoCommitDate=$cuedo_date" # ./...
