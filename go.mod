module github.com/rudifa/cuedo

go 1.20

// replace cuelang.org/go => ../cue // must use the local patched version of cue
replace cuelang.org/go => github.com/rudifa/cue v0.0.0-20240108165701-3a9556d56f39 // must be exact - go mod tidy advises

require (
    cuelang.org/go v0.7.0
    github.com/davecgh/go-spew v1.1.1
    github.com/rudifa/goutil v0.4.6
    github.com/spf13/cobra v1.8.0
)

require (
    cuelabs.dev/go/oci/ociregistry v0.0.0-20231217163254-6feb86eb6e06 // indirect
    github.com/cockroachdb/apd/v3 v3.2.1 // indirect
    github.com/emicklei/proto v1.10.0 // indirect
    github.com/google/uuid v1.3.0 // indirect
    github.com/inconshreveable/mousetrap v1.1.0 // indirect
    github.com/mitchellh/go-wordwrap v1.0.1 // indirect
    github.com/mpvl/unique v0.0.0-20150818121801-cbe035fff7de // indirect
    github.com/opencontainers/go-digest v1.0.0 // indirect
    github.com/opencontainers/image-spec v1.1.0-rc4 // indirect
    github.com/protocolbuffers/txtpbfmt v0.0.0-20230328191034-3462fbc510c0 // indirect
    github.com/rogpeppe/go-internal v1.12.0 // indirect
    github.com/spf13/pflag v1.0.5 // indirect
    github.com/tetratelabs/wazero v1.0.2 // indirect
    golang.org/x/mod v0.14.0 // indirect
    golang.org/x/net v0.19.0 // indirect
    golang.org/x/text v0.14.0 // indirect
    golang.org/x/tools v0.16.1 // indirect
    gopkg.in/yaml.v3 v3.0.1 // indirect
)
