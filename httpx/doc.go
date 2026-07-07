// Package httpServer provides the core HTTP server framework.
//
// It combines Gin, Cobra, and service management (daemon) into a single
// entry point. Use Execute() to start the server with subcommands for run,
// install, and uninstall as a system service.
//
// Modules register routes via SetRouter() in their init() functions,
// making integration pluggable without modifying core code.
package httpx
