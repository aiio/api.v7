package main

import (
	"fmt"
	"os"

	"github.com/aiio/qiniu/auth"
	"github.com/aiio/qiniu/storage"
)

var (
	accessKey = os.Getenv("QINIU_ACCESS_KEY")
	secretKey = os.Getenv("QINIU_SECRET_KEY")
	bucket    = os.Getenv("QINIU_TEST_BUCKET")
)

func main() {
	mac := auth.New(accessKey, secretKey)

	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: false,
	}
	// 指定空间所在的区域，如果不指定将自动探测
	// 如果没有特殊需求，默认不需要指定
	//cfg.Zone=&storage.ZoneHuabei
	bucketManager := storage.NewBucketManager(mac, &cfg)

	//每个batch的操作数量不可以超过1000个，如果总数量超过1000，需要分批发送
	keys := []string{
		"github1.png",
		"github2.png",
		"github3.png",
		"github4.png",
		"github5.png",
	}
	statOps := make([]string, 0, len(keys))
	for _, key := range keys {
		statOps = append(statOps, storage.URIStat(bucket, key))
	}
	rets, err := bucketManager.Batch(statOps)
	if err != nil {
		// 遇到错误
		if _, ok := err.(*storage.ErrorInfo); ok {
			for _, ret := range rets {
				// 200 为成功
				fmt.Printf("%d\n", ret.Code)
				if ret.Code != 200 {
					fmt.Printf("%s\n", ret.Data.Error)
				} else {
					fmt.Printf("%v\n", ret.Data)
				}
			}
		} else {
			fmt.Printf("batch error, %s", err)
		}
	} else {
		// 完全成功
		for _, ret := range rets {
			// 200 为成功
			fmt.Printf("%d\n", ret.Code)
			fmt.Printf("%v\n", ret.Data)
		}
	}
}
