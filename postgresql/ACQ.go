package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"net"
	"net/http"
  "time"

	"github.com/go-resty/resty/v2"
)

type ACQs struct {
	ACQs []ACQ `json:"acq"`
}

type ACQ struct {
	RxTime       float64   `json:"rxTime"`
	ExperimentId int       `json:"experimentId"`
	SignalId     int       `json:"signalId"`
	Doppler      float32   `json:"doppler"`
	CodePhase    float32   `json:"codePhase"`
	AcfCorr      []float32 `json:"acfCorr`
	NoiseFloor   float32   `json:"noiseFloor"`
	AcqMode      int16     `json:"acqMode"`
}

const (
	maxIdleConns              = 100
	maxConnsPerHost           = 100
	maxIdleConnsPerHost       = 100
	clientTimeout             = 10 * time.Second
	dialContextTimeout        = 10 * time.Second
	clientTLSHandshakeTimeout = 10 * time.Second
	clientRetryWaitTime       = 300 * time.Millisecond
	retryCount                = 3
)

func NewHttpClient() *resty.Client {
	transport := &http.Transport{
		DialContext:         (&net.Dialer{Timeout: dialContextTimeout}).DialContext,
		MaxIdleConns:        maxIdleConns,
		MaxConnsPerHost:     maxConnsPerHost,
		MaxIdleConnsPerHost: maxIdleConnsPerHost,
		TLSHandshakeTimeout: clientTLSHandshakeTimeout,
	}

	client := resty.New().
		SetTimeout(clientTimeout).
		SetRetryCount(retryCount).
		SetRetryWaitTime(clientRetryWaitTime).
		SetTransport(transport)

	return client
}
func main() {
	jsonFile, err := os.Open("ACQ.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")

  byteValue, _ := io.ReadAll(jsonFile)
  var acqs ACQs

  json.Unmarshal(byteValue, &acqs)

  // for i := 0; i < len(acqs.ACQs); i++ {
  //   
  // }

  // fmt.Println(acqs.ACQs)
  
  for index, item := range acqs.ACQs {
    fmt.Println(index, item.AcqMode)
  }

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
}
