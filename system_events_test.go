package khl

import (
	"encoding/json"
	"testing"
)

func TestSession_GuildViewUnmarshal(t *testing.T) {
	testSet := [][]byte{
		// added at 2021.07.13, 场合：转移服务器主
		[]byte(`
{"channel_type":"GROUP","type":255,"target_id":"234567890","author_id":"1","content":"[\u7cfb\u7edf\u6d88\u606f]"
,"extra":{"type":"updated_guild","body":{"id":"234567890","name":"virsaaaaaaa","user_id":"234567890","icon":"",
"notify_type":1,"region":"shanghai","enable_open":0,"open_id":0,"default_channel_id":"234567890",
"welcome_channel_id":"234567890","banner":"","banner_status":0,"custom_id":"","boost_num":0,"buffer_boost_num":0,
"level":0}},"msg_id":"badafc5c-f315-4d12-ba0c-b13af3a22839","msg_timestamp":1626154986962,"nonce":"","from_type":1}
`),
	}
	for _, item := range testSet {
		ed := &EventData{}
		err := json.Unmarshal(item, ed)
		if err != nil {
			t.Error(err)
		}
		eds := &EventDataSystem{}
		err = json.Unmarshal(ed.Extra, eds)
		if err != nil {
			t.Error(err)
		}
		gu := &GuildUpdate{}
		err = json.Unmarshal(eds.Body, gu)
		if err != nil {
			t.Error(err)
		}
	}
}
