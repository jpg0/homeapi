package main

import (
	"testing"
	"encoding/json"
	"github.com/stretchr/testify/assert"
)

func TestSampleRequest(t *testing.T) {
	req := &APIAIRequest{}
	err := json.Unmarshal([]byte(sampleReq), req)

	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.Equal(t, "ee3e4f66-95e0-4330-b1f5-5c1b38ba6c18", req.ID)
}


var sampleReq = `{
    "id": "ee3e4f66-95e0-4330-b1f5-5c1b38ba6c18",
    "timestamp": "2017-01-12T04:28:10.653Z",
    "result": {
        "source": "agent",
        "resolvedQuery": "TV",
        "speech": "",
        "action": "potential_downloads",
        "actionIncomplete": false,
        "parameters": {
            "showname": "foo",
            "showtype": "tv"
        },
        "contexts": [
            {
                "name": "generic",
                "parameters": {
                    "showtype.original": "TV",
                    "showname": "foo",
                    "telegram_chat_id": "328982357",
                    "showname.original": "",
                    "showtype": "tv"
                },
                "lifespan": 4
            }
        ],
        "metadata": {
            "intentId": "4be44fc6-c060-4fb3-9acc-87729fc84d05",
            "webhookUsed": "true",
            "webhookForSlotFillingUsed": "true",
            "intentName": "query show"
        },
        "fulfillment": {
            "speech": "",
            "messages": [
                {
                    "type": 0,
                    "speech": ""
                },
                {
                    "type": 0,
                    "speech": ""
                }
            ]
        },
        "score": 1.0
    },
    "status": {
        "code": 200,
        "errorType": "success"
    },
    "sessionId": "9358ffe0-d714-11e6-a8ef-a77544811b33",
    "originalRequest": {
        "source": "telegram",
        "data": {
            "update_id": 19482021,
            "message": {
                "date": 1484195290,
                "chat": {
                    "last_name": "Gilbert",
                    "id": 328982357,
                    "type": "private",
                    "first_name": "Jonathan"
                },
                "message_id": 881,
                "from": {
                    "last_name": "Gilbert",
                    "id": 328982357,
                    "first_name": "Jonathan"
                },
                "text": "TV"
            }
        }
    }
}`