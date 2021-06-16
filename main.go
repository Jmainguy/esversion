package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"golang.org/x/term"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func parseResponse(response aggResponse) {
	for _, bucket := range response.Aggregations.SixX_hostver.Buckets {
		fmt.Println(bucket.Key.Hostname)
	}
}

func main() {
	usernamePtr := flag.String("username", "", "username to auth with elastic")
	hostPtr := flag.String("host", "", "elastic host")
	flag.Parse()
	fmt.Print("Enter Elasticsearch password: ")
	bytePassword, err := term.ReadPassword(0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println()

	cfg := elasticsearch.Config{
		Addresses: []string{
			*hostPtr,
		},
		Username: *usernamePtr,
		Password: string(bytePassword),
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
				MinVersion:         tls.VersionTLS11,
			},
		},
	}

	client, _ := elasticsearch.NewClient(cfg)
	ctx := context.Background()

	query := aggQuery{}
	query.Aggs.SixX_hostver.Composite.Size = 10000
	query.Aggs.SevenX_hostver.Composite.Size = 10000
	var sevenSourceHostname Source
	var sevenSourceVersion Source
	var sixSourceHostname Source
	var sixSourceVersion Source
	sixSourceVersionTerm := &Term{}
	sixSourceHostnameTerm := &Term{}
	sevenSourceVersionTerm := &Term{}
	sevenSourceHostnameTerm := &Term{}
	sevenSourceVersionTerm.Terms.Field = "agent.version"
	sevenSourceHostnameTerm.Terms.Field = "agent.hostname"
	sevenSourceHostname.Hostname = sevenSourceHostnameTerm
	sevenSourceVersion.Version = sevenSourceVersionTerm
	sixSourceVersionTerm.Terms.Field = "beat.version"
	sixSourceHostnameTerm.Terms.Field = "beat.hostname"
	sixSourceHostname.Hostname = sixSourceHostnameTerm
	sixSourceVersion.Version = sixSourceVersionTerm
	query.Aggs.SixX_hostver.Composite.Sources = append(query.Aggs.SixX_hostver.Composite.Sources, sixSourceHostname)
	query.Aggs.SixX_hostver.Composite.Sources = append(query.Aggs.SixX_hostver.Composite.Sources, sixSourceVersion)
	query.Aggs.SevenX_hostver.Composite.Sources = append(query.Aggs.SevenX_hostver.Composite.Sources, sevenSourceHostname)
	query.Aggs.SevenX_hostver.Composite.Sources = append(query.Aggs.SevenX_hostver.Composite.Sources, sevenSourceVersion)
	query.Query.Bool.Filter.Range.Timestamp.Gte = "2021-06-15T22:01:51.220Z"
	query.Query.Bool.Filter.Range.Timestamp.Lte = "2021-06-15T22:16:51.220Z"

	jojo, _ := json.Marshal(query)

	res, err := client.Search(
		client.Search.WithContext(ctx),
		client.Search.WithIndex("metricbeat-*"),
		client.Search.WithBody(bytes.NewReader(jojo)),
		client.Search.WithTrackTotalHits(true),
	)

	fmt.Println(string(jojo))

	// Check for any errors returned by API call to Elasticsearch
	if err != nil {
		log.Fatalf("Elasticsearch Search() API ERROR:", err)

	} else {

		response := aggResponse{}
		json.NewDecoder(res.Body).Decode(&response)
		parseResponse(response)
		bodyBytes, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(bodyBytes))
		fmt.Println(res.StatusCode)

		defer res.Body.Close()

	}
}
