分布式爬虫，原有的爬虫项目即crawlerPro文件夹，已经不存在独立的意义
而定位更像是一个core包，所有的初始化、调度分发等逻辑都在文件夹外面
crawlerPro文件夹之外的所有文件夹，都应该放在一个distrubute文件夹下
这里为了方便，直接吧crawler copy过来了，需要注意。
1.分布式体现在，将存储和网页解析拆为两个单独的服务，通过rpc通信
2.rpc使用golang自带的，序列化使用json序列化
