package main

import (
	"auto-order/data"
	"bufio"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"math/rand"
	"os"
	"time"
)

func main() {
	glog.Debug(context.Background(), gtime.Now().String())
	glog.Info(context.Background(), "哈！又不想自己想中午吃什么了！")
	glog.Info(context.Background(), "那就让我来帮你吧！")
	glog.Info(context.Background(), "自动决定中午吃什么机器人v0.1启动中！")
	time.Sleep(2 * time.Second)
	glog.Info(context.Background(), "自动决定中午吃什么机器人v0.1启动成功！")
	glog.Info(context.Background(), "自动决定中午吃什么机器人v0.1开始工作！")
	glog.Info(context.Background(), "正在决定中......")
	time.Sleep(3 * time.Second)
	for {
		reader := bufio.NewReader(os.Stdin)
		glog.Info(context.Background(), "按下回车键进行显示：")
		_, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("读取输入时发生错误:", err)
		}
		glog.Info(context.Background(), "正在选择......")
		glog.Info(context.Background(), "正在决定中......")
		orderMap := GetRandomOrder()
		glog.Warning(context.Background(), "今天中午吃：", orderMap["key"], " 的 ", orderMap["value"])
		time.Sleep(3 * time.Second)
	}
}

func GetRandomOrder() g.Map {
	// 将 Map 的键存储到切片中
	keys := make([]string, 0, len(data.CanteenOrder))
	for k := range data.CanteenOrder {
		keys = append(keys, k)
	}

	// 设置随机种子
	rand.Seed(time.Now().UnixNano())

	// 生成随机索引
	randomIndex := rand.Intn(len(keys))

	// 获取随机键
	randomKey := keys[randomIndex]

	fmt.Println("随机获取的键:", randomKey)
	// 随机从map中获取一个key
	list := data.CanteenOrder[randomKey]
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())

	// 生成随机索引
	randomIndex = rand.Intn(len(list))

	// 获取随机值
	randomValue := list[randomIndex]

	return g.Map{
		"key":   randomKey,
		"value": randomValue,
	}
}
