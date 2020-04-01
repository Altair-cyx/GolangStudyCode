package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

// etcd client put/get demo
// use etcd/clientv3

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		// handle error!
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()
	// put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	str := `[{"path":"E:/tmp/nginx.log","topic":"web_log"},{"path":"E:/tmp/redis.log","topic":"redis_log"},{"path":"E:/tmp/tomcat.log","topic":"web_log"},{"path":"E:/tmp/apache.log","topic":"web_log"}]`
	_, err = cli.Put(ctx, "/logagent/192.168.1.3/collect_config", str)
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}
}
