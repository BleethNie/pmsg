package bot

import (
	"fmt"
	"github.com/lenye/pmsg/cmd/variable"
	dingTalk "github.com/lenye/pmsg/internal/im/dingtalk/bot"
	feiShu "github.com/lenye/pmsg/internal/im/feishu/bot"
	WorkWeiXin "github.com/lenye/pmsg/internal/im/weixin/work/bot"
)

func FeiShuSendText(token string, msg string) bool {
	arg := feiShu.CmdSendParams{
		UserAgent:   variable.UserAgent,
		AccessToken: token,
		Secret:      variable.Secret,
		MsgType:     "text",
		Data:        msg,
	}

	if err := feiShu.CmdSend(&arg); err != nil {
		fmt.Println(err)
	}
	return true
}

func DingTalkSendText(token string, msg string) bool {
	arg := dingTalk.CmdSendParams{
		UserAgent:   variable.UserAgent,
		AccessToken: token,
		Secret:      variable.Secret,
		MsgType:     "text",
		Data:        msg,
	}

	if err := dingTalk.CmdSend(&arg); err != nil {
		fmt.Println(err)
	}
	return true
}

func WorkWeiXinSendText(token string, msg string) bool {
	arg := WorkWeiXin.CmdSendParams{
		UserAgent: variable.UserAgent,
		Key:       token,
		MsgType:   "text",
		Data:      msg,
	}
	if err := WorkWeiXin.CmdSend(&arg); err != nil {
		fmt.Println(err)
	}
	return true
}
