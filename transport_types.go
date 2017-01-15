package main

import "time"

type APIAIRequest struct {
	ID              string `json:"id,omitempty"`
	Timestamp       time.Time `json:"timestamp,omitempty"`
	Result          struct {
				Source           string `json:"source,omitempty"`
				ResolvedQuery    string `json:"resolvedQuery,omitempty"`
				Speech           string `json:"speech,omitempty"`
				Action           string `json:"action,omitempty"`
				ActionIncomplete bool `json:"actionIncomplete,omitempty"`
				Parameters       map[string]string `json:"parameters,omitempty"`
				Contexts         []APIAIContext `json:"contexts,omitempty"`
				Metadata         struct {
							 IntentID                  string `json:"intentId,omitempty"`
							 WebhookUsed               string `json:"webhookUsed,omitempty"`
							 WebhookForSlotFillingUsed string `json:"webhookForSlotFillingUsed,omitempty"`
							 IntentName                string `json:"intentName,omitempty"`
						 } `json:"metadata,omitempty"`
				Fulfillment      struct {
							 Speech   string `json:"speech,omitempty"`
							 Messages []struct {
								 Type   int `json:"type,omitempty"`
								 Speech string `json:"speech,omitempty"`
							 } `json:"messages,omitempty"`
						 } `json:"fulfillment,omitempty"`
				Score            float32 `json:"score,omitempty"`
			} `json:"result,omitempty"`
	Status          struct {
				Code      int `json:"code,omitempty"`
				ErrorType string `json:"errorType,omitempty"`
			} `json:"status,omitempty"`
	SessionID       string `json:"sessionId,omitempty"`
	OriginalRequest struct {
				Source string `json:"source,omitempty"`
				Data   struct {
					       UpdateID int `json:"update_id,omitempty"`
					       Message  struct {
								Date      int `json:"date,omitempty"`
								Chat      struct {
										  LastName  string `json:"last_name,omitempty"`
										  ID        int `json:"id,omitempty"`
										  Type      string `json:"type,omitempty"`
										  FirstName string `json:"first_name,omitempty"`
									  } `json:"chat,omitempty"`
								MessageID int `json:"message_id,omitempty"`
								From      struct {
										  LastName  string `json:"last_name,omitempty"`
										  ID        int `json:"id,omitempty"`
										  FirstName string `json:"first_name,omitempty"`
									  } `json:"from,omitempty"`
								Text      string `json:"text,omitempty"`
							} `json:"message,omitempty"`
				       } `json:"data,omitempty"`
			} `json:"originalRequest,omitempty"`
}

type APIAIContext struct {
	Name       string `json:"name,omitempty"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	Lifespan   int `json:"lifespan,omitempty"`
}

func NewAPIAIContext(name string, lifespan int) *APIAIContext {
	return &APIAIContext{
		Name: name,
		Lifespan: lifespan,
		Parameters: make(map[string]interface{}),
	}
}

type APIAIResponse struct {
	Speech      string `json:"speech"`
	DisplayText string `json:"displayText"`
	Data        map[string]interface{} `json:"data"`
	ContextOut  []APIAIContext `json:"contextOut"`
	Source      string `json:"source"`
}

func NewAPIAIResponse(msg string) *APIAIResponse {
	res := &APIAIResponse{
		Data: make(map[string]interface{}),
		ContextOut: make([]APIAIContext, 0),
	}

	res.DisplayText = msg
	res.Data["telegram"] = &TelegramData{Text:msg}

	return res
}

func (res *APIAIResponse) SetMessage(msg string) {
	res.DisplayText = msg
	res.Data["telegram"] = &TelegramData{Text:msg}
}

func (res *APIAIResponse) AddContext(ctx *APIAIContext) {
	res.ContextOut = append(res.ContextOut, *ctx)
}

func (req *APIAIRequest) GetContext(name string) *APIAIContext {
	for i := range req.Result.Contexts {
		ctx := req.Result.Contexts[i]

		if ctx.Name == name {
			return &ctx
		}
	}

	return nil
}

type TelegramData struct {
	Text string `json:"text"`
}