// Package upload provides a unified file upload abstraction over multiple storage backends.
//
// Supported drivers: local filesystem, Aliyun OSS, Tencent COS, AWS S3,
// and Qiniu Kodo. Use NewOss() to get a driver instance by name; all
// drivers implement the OSS interface (upload, delete, token generation).
package upload
