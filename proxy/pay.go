// @author cold bin
// @date 2023/9/1

package proxy

import (
	"context"
)

// Uploader 上传文件的抽象
type Uploader interface {
	UploadSingle(ctx context.Context, key string, value []byte)
	UploadMultiple(ctx context.Context, keys []string, values [][]byte)
}

// UploaderProxy 上传的代理者
//
//	将所有实现都进行代理，我们可以在Up执行前或后做一些事情，例如实现一些hook
type UploaderProxy struct {
	Up Uploader
}

func (u *UploaderProxy) UploadSingle(ctx context.Context, key string, value []byte) {
	// 这里可以做一些事情：
	//  1、校验文件名是否正确
	//  2、记录图床开始上传时间
	//  3、统计上传频率，做限流等
	u.Up.UploadSingle(ctx, key, value)
	// 这里也可以做一些事情：
	//  1、hook结果
	//  2、记录图床结束上传时间
	//  3、记录图床上传时间
}

func (u *UploaderProxy) UploadMultiple(ctx context.Context, keys []string, values [][]byte) {
	u.Up.UploadMultiple(ctx, keys, values)
}

type QiniuOss struct {
	// 注入七牛的依赖
}

func (q *QiniuOss) UploadSingle(ctx context.Context, key string, value []byte) {

}

func (q *QiniuOss) UploadMultiple(ctx context.Context, keys []string, values [][]byte) {

}

type AliOss struct {
	// 注入阿里云oss依赖
}

func (a *AliOss) UploadSingle(ctx context.Context, key string, value []byte) {

}

func (a *AliOss) UploadMultiple(ctx context.Context, keys []string, values [][]byte) {

}
