package dingtalk

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	robot "github.com/alibabacloud-go/dingtalk/robot_1_0"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/heqiaomu/gtools/gerror"
	"github.com/heqiaomu/gtools/ghttp"
	"github.com/heqiaomu/gtools/glog"
	"github.com/kevin2027/easy-dingtalk/dingtalk"
	"github.com/kevin2027/easy-dingtalk/utils"
	"github.com/pkg/errors"
	"net/http"
	"net/url"
	"time"
)

const (
	epSend = "/v1.0/robot/groupMessages/send"
)

type Client struct {
	Address  string
	token    *token
	client   *robot.Client
	dingtalk dingtalk.Dingtalk
	httpCli  ghttp.Client
}

type token struct {
	token  string
	expire time.Duration
	start  time.Time
}

func (t *token) isExpire() *token {
	if t.start.Add(t.expire).Before(time.Now()) {
		return t
	}
	// 这里需要自己去生成 新的 token
	token, _, err := DefaultClient().Oauth2().GetAccessToken()
	if err != nil {
		glog.Errorf("New Token failed. %s", err.Error())
		return nil
	}
	t.token = token
	return t

}

func DefaultClient() dingtalk.Dingtalk {
	cli, _, _ := dingtalk.NewDingtalk(utils.DingtalkOptions{
		AppKey:    "dingdnadh8qedgtb6t9x",
		AppSecret: "k75cTqYhGNQZbUF6UWhjzZcEkPVSChf5eaUaQqfvkCHeutE6nzUkata4Kws8pDxS",
		AgentId:   16769005,
	})
	return cli
}

func NewClient(address, appkey, appSecret string, agentID int64) (*Client, error) {
	config := &openapi.Config{}
	config.Protocol = tea.String("https")
	config.RegionId = tea.String("central")

	dingtalkCli, _, _err := dingtalk.NewDingtalk(utils.DingtalkOptions{
		AppKey:    appkey,
		AppSecret: appSecret,
		AgentId:   agentID,
	})
	if _err != nil {
		return nil, _err
	}
	to, _, _err := dingtalkCli.Oauth2().GetAccessToken()
	if _err != nil {
		return nil, _err
	}
	t := token{
		token:  to,
		expire: time.Second * 7200,
		start:  time.Now(),
	}

	cli := &robot.Client{}
	cli, _err = robot.NewClient(config)
	if _err != nil {
		return nil, _err
	}

	client, _err := ghttp.NewClient(ghttp.Config{Address: address})
	if _err != nil {
		return nil, _err
	}
	return &Client{
		httpCli:  client,
		client:   cli,
		dingtalk: dingtalkCli,
		token:    &t,
	}, _err
}

// SendMessage SendWithWebhookWebhook https://oapi.dingtalk.com/robot/send?access_token=64ee7f12d49f6451811240dc2950e92e6bd00c74fde66bb84a0947e52e9c8034
func (cli *Client) SendMessage(ctx context.Context, data []byte) error {

	u := cli.httpCli.URL("", nil)
	req, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json;charset=utf-8")

	resp, body, err := cli.httpCli.Do(ctx, req)
	if err != nil {
		return err
	}
	if resp.StatusCode/100 != 2 {
		return gerror.ErrorHttpRequest
	}
	glog.Infof("body=%s", string(body))
	var errCode Error
	if err = json.Unmarshal(body, &errCode); err != nil {
		return err
	}
	if errCode.Errcode != 0 {
		glog.Infof("errCode=%d, errMsg=%s", errCode.Errcode, errCode.Errmsg)
		return errors.New("执行失败" + errCode.Errmsg)
	}
	return nil

}

type Error struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func getSign() string {
	timestamp := time.Now().UnixMilli()
	secret := "SEC56db986f5d6eefacdb53b019cabe1461b760567179aaba5e57635d9acdf1ee2f"
	stringToSign := fmt.Sprintf("%d\n%s", timestamp, secret)
	hash := hmac.New(sha256.New, []byte(secret))
	hash.Write([]byte(stringToSign))
	signData := hash.Sum(nil)
	sign := url.QueryEscape(base64.StdEncoding.EncodeToString(signData))
	return sign
}
