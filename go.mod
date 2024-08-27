module github.com/rudifa/cuedo

go 1.22

toolchain go1.22.2

replace cuelang.org/go => ../cue // must use the local patched version of cue

// replace github.com/rudifa/goutil => ../goutil // must use the local version of goutil

// replace cuelang.org/go => github.com/rudifa/cue v0.0.0-20240108165701-3a9556d56f39 // must be exact - go mod tidy advises

require (
	cuelang.org/go v0.5.0
	github.com/davecgh/go-spew v1.1.1
	github.com/rudifa/goutil v0.4.10
	github.com/spf13/cobra v1.8.1
)

require (
	cuelabs.dev/go/oci/ociregistry v0.0.0-20240807094312-a32ad29eed79 // indirect
	github.com/cockroachdb/apd/v3 v3.2.1 // indirect
	github.com/emicklei/proto v1.13.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.1.0 // indirect
	github.com/pelletier/go-toml/v2 v2.2.2 // indirect
	github.com/protocolbuffers/txtpbfmt v0.0.0-20230328191034-3462fbc510c0 // indirect
	github.com/rogpeppe/go-internal v1.12.1-0.20240709150035-ccf4b4329d21 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/tetratelabs/wazero v1.6.0 // indirect
	golang.org/x/mod v0.20.0 // indirect
	golang.org/x/net v0.28.0 // indirect
	golang.org/x/oauth2 v0.22.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/text v0.17.0 // indirect
	golang.org/x/tools v0.24.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
