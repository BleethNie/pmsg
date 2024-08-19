package bot

import (
	"github.com/BleethNie/pmsg/cmd/variable"
	dingTalk "github.com/BleethNie/pmsg/internal/im/dingtalk/bot"
	feiShu "github.com/BleethNie/pmsg/internal/im/feishu/bot"
	WorkWeiXin "github.com/BleethNie/pmsg/internal/im/weixin/work/bot"
)

func FeiShuSendText(token string, msg string) error {
	arg := feiShu.CmdSendParams{
		UserAgent:   variable.UserAgent,
		AccessToken: token,
		Secret:      variable.Secret,
		MsgType:     "text",
		Data:        msg,
	}

	return feiShu.CmdSend(&arg)
}

func DingTalkSendText(token string, msg string) error {
	arg := dingTalk.CmdSendParams{
		UserAgent:   variable.UserAgent,
		AccessToken: token,
		Secret:      variable.Secret,
		MsgType:     "text",
		Data:        msg,
	}

	return dingTalk.CmdSend(&arg)
}

func WorkWeiXinSendText(token string, msg string) error {
	arg := WorkWeiXin.CmdSendParams{
		UserAgent: variable.UserAgent,
		Key:       token,
		MsgType:   "text",
		Data:      msg,
	}
	return WorkWeiXin.CmdSend(&arg)
}
