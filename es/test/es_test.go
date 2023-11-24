package test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"testing"
	"time"
	"uscan/spec"
)

func TestDeleteIndex(t *testing.T) {
	// 创建client
	client, err := elastic.NewClient(
		//elastic.SetURL("http://127.0.0.1:9200", "http://127.0.0.1:9201"),
		elastic.SetURL("http://127.0.0.1:9200"),
		// 禁用嗅探器用于兼容内网ip
		elastic.SetSniff(false),
		elastic.SetBasicAuth("user", "nvUt974rcNeg==*k0W3W"))
	if err != nil {
		// Handle error
		fmt.Printf("连接失败: %v\n", err)
	} else {
		fmt.Println("连接成功")
	}
	// 执行ES请求需要提供一个上下文对象
	ctx := context.Background()
	_, err = client.DeleteIndex("weibo").Do(ctx)
	if err != nil {
		fmt.Printf("删除失败: %v\n", err)
	}
	fmt.Printf("删除成功")
}

func TestQuery(t *testing.T) {
	client, err := elastic.NewClient(
		//elastic.SetURL("http://127.0.0.1:9200", "http://127.0.0.1:9201"),
		elastic.SetURL("http://127.0.0.1:9200"),
		// 禁用嗅探器用于兼容内网ip
		elastic.SetSniff(false),
		elastic.SetBasicAuth("user", "nvUt974rcNeg==*k0W3W"))
	if err != nil {
		// Handle error
		fmt.Printf("连接失败: %v\n", err)
	} else {
		fmt.Println("连接成功")
	}
	ctx := context.Background()
	// 定义要查询的用户
	userToQuery := "user99999" // 假设我们要查询 user123 的推文

	// 构建一个 term 查询
	termQuery := elastic.NewTermQuery("user", userToQuery)
	// 开始计时
	startTime := time.Now()
	// 执行搜索
	searchResult, err := client.Search().
		Index("weibo").
		Query(termQuery).
		Sort("created", true). // 根据创建时间排序
		From(0).Size(10).      // 分页参数
		Pretty(true).
		Do(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// 打印搜索结果
	fmt.Printf("查询到 %d 条推文\n", searchResult.TotalHits())
	for _, hit := range searchResult.Hits.Hits {
		var wb spec.Weibo
		_ = json.Unmarshal(hit.Source, &wb)
		fmt.Printf("用户: %s, 推文: %s, 转发数: %d\n", wb.User, wb.Message, wb.Retweets)
		// 可以进一步解析 hit.Source 以获取完整的推文数据
		fmt.Printf("推文 ID: %s\n", hit.Id)
	}
	// 结束计时
	duration := time.Since(startTime)
	fmt.Printf("耗时: %v\n", duration)

}
