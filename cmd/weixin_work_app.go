package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/lenye/pmsg/pkg/flags"
	"github.com/lenye/pmsg/pkg/weixin/work/message"
)

// weiXinWorkAppCmd 企业微信应用消息
var weiXinWorkAppCmd = &cobra.Command{
	Use:   "app",
	Short: "publish work weixin app message",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		arg := message.CmdWorkSendAppParams{
			UserAgent:              userAgent,
			AccessToken:            accessToken,
			CorpID:                 corpID,
			CorpSecret:             corpSecret,
			ToUser:                 toUser,
			ToParty:                toParty,
			ToTag:                  toTag,
			AgentID:                agentID,
			MsgType:                msgType,
			Safe:                   safe,
			EnableIDTrans:          enableIDTrans,
			EnableDuplicateCheck:   enableDuplicateCheck,
			DuplicateCheckInterval: duplicateCheckInterval,
			Data:                   args[0],
		}
		if err := message.CmdWorkSendApp(&arg); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}

func init() {
	weiXinWorkCmd.AddCommand(weiXinWorkAppCmd)

	weiXinWorkSetAccessTokenFlags(weiXinWorkAppCmd)

	weiXinWorkAppCmd.Flags().StringVarP(&toUser, flags.ToUser, "o", "", "work weixin user id list")
	weiXinWorkAppCmd.Flags().StringVarP(&toParty, flags.ToParty, "p", "", "work weixin party id list")
	weiXinWorkAppCmd.Flags().StringVarP(&toTag, flags.ToTag, "g", "", "work weixin tag id list")

	weiXinWorkAppCmd.Flags().Int64VarP(&agentID, flags.AgentID, "e", 0, "work weixin agent id (required)")
	weiXinWorkAppCmd.MarkFlagRequired(flags.AgentID)

	weiXinWorkAppCmd.Flags().StringVar(&msgType, flags.MsgType, "", "message type (required)")
	weiXinWorkAppCmd.MarkFlagRequired(flags.MsgType)

	weiXinWorkAppCmd.Flags().IntVar(&safe, flags.Safe, 0, "safe")
	weiXinWorkAppCmd.Flags().IntVarP(&enableIDTrans, flags.EnableIDTrans, "r", 0, "enable id translated")
	weiXinWorkAppCmd.Flags().IntVarP(&enableDuplicateCheck, flags.EnableDuplicateCheck, "c", 0, "enable duplicate check")
	weiXinWorkAppCmd.Flags().IntVarP(&duplicateCheckInterval, flags.DuplicateCheckInterval, "d", 1800, "duplicate check interval")
}
