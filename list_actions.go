package main

import (
	"github.com/tehjojo/go-sabnzbd"
	"github.com/juju/errors"
	"bytes"
	"fmt"
	"reflect"
)

type ListActions Configuration

func list(ac *GenericContext, cfg map[string]string) (*ActionResponse, error) {

	s, err := sabnzbd.New(sabnzbd.Addr(cfg["sabnzbd_address"]), sabnzbd.ApikeyAuth(cfg["sabnzbd_apikey"]))

	if err != nil {
		return nil, errors.Annotate(err, "Failed to create NZB client")
	}

	history, err := s.History(0, 10)

	var buffer bytes.Buffer

	for _, slot := range history.Slots {
		buffer.WriteString(fmt.Sprintf("%v\n", slot.Name))
	}

	ac.Add("period", "7 days")
	ac.Add("data", buffer.String())

	return ac.Response(), nil
}


func InitListActions() {
	Register("list", AsGeneric(list))
}

func AsGeneric(toWrap func(*GenericContext, map[string]string) (*ActionResponse, error)) ActionRunner {
	return func (newCtx, oldCtx map[string]string, cfg map[string]string) (*ActionResponse, error) {
		return toWrap(NewGenericContext(newCtx, oldCtx), cfg)
	}
}

func WithHydration(toWrap interface{}) ActionRunner {
	//toWrap = func(*SpecificContext, map[string]string) (*ActionResponse, error)

	toWrapVal := reflect.ValueOf(toWrap)
	specificType := toWrapVal.Type().In(0)


	return func (newCtx, oldCtx map[string]string, cfg map[string]string) (*ActionResponse, error) {

		specificCtx := reflect.New(specificType)

		Hydrate(oldCtx, specificCtx.Interface())

		params := []reflect.Value{
			specificCtx,
			reflect.ValueOf(newCtx),
			reflect.ValueOf(cfg),
		}

		returns := toWrapVal.Call(params)

		return returns[0].Interface().(*ActionResponse), returns[0].Interface().(error)
	}
}