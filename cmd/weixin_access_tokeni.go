package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/lenye/pmsg/pkg/http/client"
	"github.com/lenye/pmsg/pkg/weixin/token"
)

// weiXinAccessTokenCmd 获取微信接口调用凭证
var weiXinAccessTokenCmd = &cobra.Command{
	Use:   "token",
	Short: "get weixin access token (mp, miniprogram)",
	Run: func(cmd *cobra.Command, args []string) {
		if err := WeiXinGetAccessToken(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}

func init() {
	weiXinCmd.AddCommand(weiXinAccessTokenCmd)
}

// WeiXinGetAccessToken 获取微信接口调用凭证
func WeiXinGetAccessToken() error {

	if accessToken != "" {
		return fmt.Errorf("flags %q not required", nameAccessToken)
	}

	if appID == "" {
		return fmt.Errorf("flags %q required", nameAppID)
	}

	if userAgent != "" {
		client.UserAgent = userAgent
	}

	accessTokenResp, err := token.GetAccessToken(appID, appSecret)
	if err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("access_token: %+v", accessTokenResp))

	return nil
}
