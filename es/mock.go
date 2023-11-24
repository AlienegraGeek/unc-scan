package es

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"strconv"
	"time"
	"uscan/spec"
)

func mockData(ctx context.Context, client *elastic.Client) {
	// 确保索引存在
	createIndexIfNotExists(ctx, client, "weibo")

	// 模拟数据并批量插入
	bulkRequest := client.Bulk()
	for i := 0; i < 100000; i++ { // 生成 10000 条数据
		blk := spec.Weibo{
			User:     "user" + strconv.Itoa(i),
			Message:  "这是一条微博",
			Retweets: i,
			Image:    "",
			Created:  time.Time{},
			Tags:     nil,
			Location: "",
			Suggest:  nil,
		}

		req := elastic.NewBulkIndexRequest().Index("weibo").Doc(blk)
		bulkRequest = bulkRequest.Add(req)
	}

	// 执行批量插入
	_, err := bulkRequest.Do(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("批量插入完成")
}

func createIndexIfNotExists(ctx context.Context, client *elastic.Client, indexName string) {
	exists, err := client.IndexExists(indexName).Do(ctx)
	if err != nil {
		panic(err)
	}
	if !exists {
		_, err := client.CreateIndex(indexName).Do(ctx)
		if err != nil {
			panic(err)
		}
	}
}
