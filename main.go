package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"golang.org/x/term"
	//"github.com/elastic/go-elasticsearch/v7/estransport"
	"bytes"
	"context"
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func constructQuery(q string, size int) *strings.Reader {

	// Build a query string from string passed to function
	var query = `{"query": {`

	// Concatenate query string with string passed to method call
	query = query + q

	// Use the strconv.Itoa() method to convert int to string
	query = query + `}, "size": ` + strconv.Itoa(size) + `}`
	//fmt.Println("\nquery:", query)

	// Check for JSON errors
	isValid := json.Valid([]byte(query)) // returns bool

	// Default query is "{}" if JSON is invalid
	if isValid == false {
		fmt.Println("constructQuery() ERROR: query string not valid:", query)
		fmt.Println("Using default match_all query")
		query = "{}"
	}

	// Build a new string from JSON query
	var b strings.Builder
	b.WriteString(query)

	// Instantiate a *strings.Reader object from string
	read := strings.NewReader(b.String())

	// Return a *strings.Reader object
	return read
}

func parseResponse(response Response) {
	hosts := make(map[string]beat)
	version := make(map[string]bool)
	for _, hit := range response.Hits.Hits {
		var name string
		if hit.Source.Beat.Name != "" {
			name = hit.Source.Beat.Name
		} else if hit.Source.Agent.Name != "" {
			name = hit.Source.Agent.Name
		}
		hitBeat := hosts[name]
		if hit.Source.Agent.Version != "" {
			hitBeat.AgentVersion = hit.Source.Agent.Version
			version[hit.Source.Agent.Version] = true
		}
		if hit.Source.Beat.Version != "" {
			hitBeat.BeatVersion = hit.Source.Beat.Version
			version[hit.Source.Beat.Version] = true
		}
		hosts[name] = hitBeat
	}
	for k, host := range hosts {
		fmt.Printf("%s BeatVersion: %s AgentVersion: %s\n", k, host.BeatVersion, host.AgentVersion)
	}
	for key, _ := range version {
		fmt.Println(key)
	}
}

func oldSchool() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	f, err := os.Open("payload.json")
	if err != nil {
		// handle err
	}
	defer f.Close()

	req, err := http.NewRequest("GET", "example.com:9200/metricbeat-*/_search", f)
	if err != nil {
		fmt.Println(err)
	}
	req.SetBasicAuth("username", "password")
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()

	response := Response{}
	json.NewDecoder(resp.Body).Decode(&response)
	hosts := make(map[string]beat)
	version := make(map[string]bool)
	for _, hit := range response.Hits.Hits {
		var name string
		if hit.Source.Beat.Name != "" {
			name = hit.Source.Beat.Name
		} else if hit.Source.Agent.Name != "" {
			name = hit.Source.Agent.Name
		}
		hitBeat := hosts[name]
		if hit.Source.Agent.Version != "" {
			hitBeat.AgentVersion = hit.Source.Agent.Version
			version[hit.Source.Agent.Version] = true
		}
		if hit.Source.Beat.Version != "" {
			hitBeat.BeatVersion = hit.Source.Beat.Version
			version[hit.Source.Beat.Version] = true
		}
		hosts[name] = hitBeat
	}
	for k, host := range hosts {
		fmt.Printf("%s BeatVersion: %s AgentVersion: %s\n", k, host.BeatVersion, host.AgentVersion)
	}
	for key, _ := range version {
		fmt.Println(key)
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

	//log.Print(es.Transport.(*estransport.Client).URLs())
	//fmt.Println(client.Info())

	/*
	   	var query = `
	       "bool": {
	         "must": [],
	         "filter": [
	           {
	             "match_all": {}
	           },
	           {
	             "range": {
	               "@timestamp": {
	                 "gte": "2021-06-15T22:01:51.220Z",
	                 "lte": "2021-06-15T22:16:51.220Z",
	                 "format": "strict_date_optional_time"
	               }
	             }
	           }
	         ],
	         "should": [],
	         "must_not": []
	       }
	   `
	*/

	var query = `
    "bool": {
      "must": [],
      "filter": [
        {
          "match_all": {}
        },
        {
          "exists": {
            "field": "agent.version"
          }
        },
        {
          "range": {
            "@timestamp": {
              "gte": "2021-06-16T00:01:52.634Z",
              "lte": "2021-06-16T00:02:52.634Z",
              "format": "strict_date_optional_time"
            }
          }
        }
      ],
      "should": [],
      "must_not": []
    }
`

	read := constructQuery(query, 10000)
	//fmt.Println("read:", read)
	var buf bytes.Buffer

	// Attempt to encode the JSON query and look for errors
	if err := json.NewEncoder(&buf).Encode(read); err != nil {
		log.Fatalf("json.NewEncoder() ERROR:", err)

		// Query is a valid JSON object
	} else {
		//fmt.Println("json.NewEncoder encoded query:", read, "\n")
		res, err := client.Search(
			client.Search.WithContext(ctx),
			client.Search.WithIndex("metricbeat-*"),
			client.Search.WithBody(read),
			client.Search.WithTrackTotalHits(true),
		)

		// Check for any errors returned by API call to Elasticsearch
		if err != nil {
			log.Fatalf("Elasticsearch Search() API ERROR:", err)

		} else {

			response := Response{}
			json.NewDecoder(res.Body).Decode(&response)
			parseResponse(response)

			defer res.Body.Close()

		}
	}
}
