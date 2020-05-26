# HP-Cache

快速。 性能在多核CPU上表现更好。
线程安全的。 并发goroutine可以读写单个缓存实例。
fastcache设计用于存储大量 K/V 数据而无需GC开销。
Fastcache在创建期间达到设置的最大大小时会自动驱逐旧条目。
简单的API。
