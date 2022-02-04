package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Khan/genqlient/graphql"
)

func main() {
	httpClient := http.Client{}
	graphqlClient := graphql.NewClient("https://app.testing.husarnet.com/graphql", &httpClient)

	_ = `# @genqlient
	query getNetworks {
		allNetworks {
			name
		}
	}
	`

	var networks, err = getNetworks(context.Background(), graphqlClient)
	if err != nil {
		return
	}
	for _, network := range networks.AllNetworks {
		fmt.Println("Network name: ", network.Name)
	}

}

//go:generate go run github.com/Khan/genqlient genqlient.yaml
