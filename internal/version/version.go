package version

// Version is the CLI version string. Override at build time:
//   go build -ldflags "-X github.com/umuttalha/deploy/internal/version.Version=v0.1.0"
var Version = "dev"
