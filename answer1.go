package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
    ctx := context.Background()
    client := redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
        // Addr:     "35.201.197.213:9322",
    })

    // 每 10 秒鐘呼叫一次 API
    for range time.Tick(10 * time.Second) {
        // 呼叫指定的 API
        resp, err := http.Get("http://www-dev.vir888.com/entrance/langi18n/list.json")
        if err != nil {
            fmt.Println("Error:", err)
            continue
        }

        // 讀取回應內容
        body, err := ioutil.ReadAll(resp.Body)
        resp.Body.Close()
        if err != nil {
            fmt.Printf("Error: %w", err)
            continue
        }

        // 將回應內容寫入 Redis 中
        err = client.Set(ctx, "joey_chin", string(body), 0).Err()
        if err != nil {
            fmt.Printf("Error: %w", err)
            continue
        }
    }
}
