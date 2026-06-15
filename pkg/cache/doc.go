// Package cache provides in-memory caching with a built-in captcha store.
//
// The local cache is backed by gkit's local_cache and supports TTL-based
// expiration. The Base64CaptchaStore implements the captcha store interface
// for verification code storage.
package cache
