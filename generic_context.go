package main

const DEFAULT_CONTEXT_VALUE string = "true"

type GenericContext struct {
	newCtx map[string]string
	oldCtx map[string]string
	addCtx map[string]string
	rmCtx []string
}

func NewGenericContext(newCtx, oldCtx map[string]string) *GenericContext {
	return &GenericContext{
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

func (ac*GenericContext) MergeNew(key string) (string, bool) {
	newS, newP := ac.newCtx[key]
	oldS, oldP := ac.oldCtx[key]

	if newP {
		ac.addCtx[key] = newS
		return newS, true

	} else {
		return oldS, oldP
	}
}

func (ac*GenericContext) Add(key, value string) *GenericContext {
	ac.addCtx[key] = value
	return ac
}

func (ac*GenericContext) AddKey(key string) *GenericContext {
	ac.addCtx[key] = DEFAULT_CONTEXT_VALUE
	return ac
}

func (ac*GenericContext) Remove(key string) *GenericContext {
	ac.rmCtx = append(ac.rmCtx, key)
	return ac
}

func (ac*GenericContext) RemoveAllNow() {
	for k := range ac.oldCtx {
		ac.Remove(k)
	}

	ac.oldCtx = make(map[string]string)
}

func (ac *GenericContext) WillContain (key string) bool {
	_, newContains := ac.newCtx[key]

	if newContains { //added
		return true
	}

	_, oldContains := ac.oldCtx[key]

	if oldContains && !ac.inRemoveList(key) {
		return true
	}

	return false
}

func (ac*GenericContext) inRemoveList(key string) bool {
	for _, b := range ac.rmCtx {
		if b == key {
			return true
		}
	}
	return false
}

func (ac*GenericContext) WriteTo(ar *ActionResponse) *ActionResponse {
	ar.AddContext = ac.addCtx
	ar.RemoveContext = ac.rmCtx
	return ar
}

func (ac *GenericContext) Response() *ActionResponse {
	return ac.WriteTo(&ActionResponse{})
}