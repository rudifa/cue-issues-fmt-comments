module github.com/rudifa/cuedo

go 1.21

replace cuelang.org/go => ../cue // must use the local patched version of cue

// replace cuelang.org/go => github.com/rudifa/cue v0.0.0-20240108165701-3a9556d56f39 // must be exact - go mod tidy advises

require (
	cuelang.org/go v0.5.0
	github.com/davecgh/go-spew v1.1.1
	github.com/rudifa/goutil v0.4.9
	github.com/spf13/cobra v1.8.0
)

require (
	cuelabs.dev/go/oci/ociregistry v0.0.0-20240404174027-a39bec0462d2 // indirect
	github.com/cockroachdb/apd/v3 v3.2.1 // indirect
	github.com/emicklei/proto v1.10.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.1.0 // indirect
	github.com/protocolbuffers/txtpbfmt v0.0.0-20230328191034-3462fbc510c0 // indirect
	github.com/rogpeppe/go-internal v1.12.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/tetratelabs/wazero v1.6.0 // indirect
	golang.org/x/mod v0.16.0 // indirect
	golang.org/x/net v0.22.0 // indirect
	golang.org/x/oauth2 v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/tools v0.19.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
