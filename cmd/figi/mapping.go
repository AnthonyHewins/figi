package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/AnthonyHewins/figi"
)

type mappingCmd struct{}

func (m mappingCmd) name() string { return "mapping" }

func (m mappingCmd) help() string { return "Get the mapping for a symbol" }

func (m mappingCmd) run(a *args) error {
	c := figi.New()

	resp, err := c.Mapping(context.Background(), &figi.MappingRequest{
		IDType:  figi.TICKER,
		IDValue: a.shift(),
	})

	if err == nil {
		if buf, err := json.Marshal(resp); err == nil {
			fmt.Println(string(buf))
		}
	}

	return err
}
