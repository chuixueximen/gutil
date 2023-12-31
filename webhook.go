package gutil

import (
	"encoding/json"
	"fmt"
	xhttp "github.com/chuixueximen/gutil/http"
)

// SendMsgToDinTalk 发送消息到钉钉
func SendMsgToDinTalk(title, text, url string, phones ...string) error {
	if len(phones) > 0 {
		for _, p := range phones {
			text = fmt.Sprintf("@%s%s", p, text)
		}
	}
	data := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]string{
			"title": title,
			"text":  text,
		},
		"at": map[string][]string{
			"atMobiles": phones,
		},
	}
	// 数据编码
	dataJson, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("SendMsgToDinTalk: data marshal err-> %v", err)
	}
	// 发送
	_, err = xhttp.PostJSON(url, dataJson)
	if err != nil {
		return fmt.Errorf("SendMsgToDinTalk: Post json err-> %v", err)
	}
	return nil
}
