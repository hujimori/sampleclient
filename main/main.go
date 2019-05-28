package main

import (
	"context"
	"fmt"

	"github.com/hujimori/sampleclient"
)

func main() {
	qiita, err := sampleclient.NewClient("https://qiita.com/api/v2", "TOKEN", nil)
	if err != nil {
		panic(err)
	}

	items, err := qiita.GetUserItems(context.Background(), "yaotti", 2, 3)
	if err != nil {
		panic(err)
	}

	for _, item := range items {
		fmt.Printf("%+v\n", item)
	}
}
