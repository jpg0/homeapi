package main

const DEFAULT_VALUE string = "true"

type ActionContext struct {
	newCtx map[string]string
	oldCtx map[string]string
	addCtx map[string]string
	rmCtx []string
}

func NewActionContext(newCtx, oldCtx map[string]string) *ActionContext {
	return &ActionContext{
		newCtx:newCtx,
		oldCtx:oldCtx,
		addCtx:make(map[string]string),
		rmCtx:make([]string, 0),
	}
}

//func (ac* ActionContext) GetNewOrExisting(key string) (string, bool) {
//	s, p := ac.newCtx[key]
//
//	if p {
//		return s, p
//	} else {
//		return ac.oldCtx[key]
//	}
//}

func (ac* ActionContext) MergeNew(key string) (string, bool) {
	newS, newP := ac.newCtx[key]
	oldS, oldP := ac.oldCtx[key]

	if newP {
		ac.addCtx[key] = newS
		return newS, true

	} else {
		return oldS, oldP
	}
}

func (ac* ActionContext) Add(key, value string) *ActionContext {
	ac.addCtx[key] = value
	return ac
}

func (ac* ActionContext) AddKey(key string) *ActionContext {
	ac.addCtx[key] = DEFAULT_VALUE
	return ac
}

func (ac* ActionContext) Remove(key string) *ActionContext {
	ac.rmCtx = append(ac.rmCtx, key)
	return ac
}

func (ac* ActionContext) RemoveAllNow() {
	for k := range ac.oldCtx {
		ac.Remove(k)
	}

	ac.oldCtx = make(map[string]string)
}

func (ac* ActionContext) WriteTo(ar *ActionResponse) *ActionResponse {
	ar.AddContext = ac.addCtx
	ar.RemoveContext = ac.rmCtx
	return ar
}

func (ac *ActionContext) Response() *ActionResponse {
	return ac.WriteTo(&ActionResponse{})
}