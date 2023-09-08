package redis

import "time"

type (
	// A RedisConf is a redis config.
	RedisConf struct {
		Host     string // Host: redis 服务地址 ip:port, 如果是 redis cluster 则为 ip1:port1,ip2:port2,ip3:port3...
		Type     string `json:",default=node,options=node|cluster"` // Type: node 单节点 redis，cluster redis 集群
		Pass     string `json:",optional"`                          // Pass: 密码
		Tls      bool   `json:",optional"`                          // Tls: 是否开启tls
		NonBlock bool   `json:",default=true"`
		// PingTimeout is the timeout for ping redis.
		PingTimeout time.Duration `json:",default=1s"`
	}

	// A RedisKeyConf is a redis config with key.
	RedisKeyConf struct {
		RedisConf
		Key string
	}
)
