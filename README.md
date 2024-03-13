`pmsg`是一个发送消息的小工具。


## 说明

- 项目克隆自 https://github.com/lenye/pmsg
- 为满足自己项目的需要,做了一些修改

## 发送机器人消息
```go
package main

import (
	"github.com/BleethNie/pmsg/pkg/bot"
)


func main() {
	bot.DingTalkSendText("1a22ec329c4c720d55709f87765b3066eddd770e263ac93c8c6cadeef144d54e", "测试  通过golang发送钉钉消息")
	bot.FeiShuSendText("1a22ec329c4c720d55709f87765b3066eddd770e263ac93c8c6cadeef144d54e", "测试  通过golang发送飞书消息")
	bot.WorkWeiXinSendText("1a22ec329c4c720d55709f87765b3066eddd770e263ac93c8c6cadeef144d54e", "测试  通过golang发送企微消息")
}

```

## License

`pmsg` is released under the [Apache 2.0 license](https://github.com/lenye/pmsg/blob/main/LICENSE). 