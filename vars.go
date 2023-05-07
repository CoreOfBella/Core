package main

import (
	"crypto/tls"
	"net/http"
	"regexp"
	"runtime"
	"time"

	"github.com/spf13/afero"
)

// =========================================
// THIS IS ONLY A TEMPLATE FOR NOW
// =========================================
// This is where we put all variables
// that can be accessed by any of the bot's functions.

const (
	Gigabyte = 1 << 30
	Megabyte = 1 << 20
	Kilobyte = 1 << 10

	timeoutTr     = 2 * time.Hour
	memCacheLimit = 300 << 20 // 300 MB
	httpPort      = ":7777"
)

var (
	discordBotToken = "" // fill your discord bot token here

	tlsConf = &tls.Config{
		InsecureSkipVerify: false,
	}

	h1Tr = &http.Transport{
		DisableKeepAlives:      false,
		DisableCompression:     false,
		ForceAttemptHTTP2:      false,
		TLSClientConfig:        tlsConf,
		TLSHandshakeTimeout:    30 * time.Second,
		ResponseHeaderTimeout:  30 * time.Second,
		IdleConnTimeout:        90 * time.Second,
		ExpectContinueTimeout:  1 * time.Second,
		MaxIdleConns:           1000,     // Prevents resource exhaustion
		MaxIdleConnsPerHost:    100,      // Increases performance and prevents resource exhaustion
		MaxConnsPerHost:        0,        // 0 for no limit
		MaxResponseHeaderBytes: 64 << 10, // 64k
		WriteBufferSize:        64 << 10, // 64k
		ReadBufferSize:         64 << 10, // 64k
	}

	normalclient = &http.Client{
		Timeout:   15 * time.Second,
		Transport: h1Tr,
	}

	openaihttpclient = &http.Client{
		Timeout:   3 * time.Minute,
		Transport: h1Tr,
	}

	universalLogs      []string
	universalLogsLimit = 30

	statusInt   = 0
	statusSlice = []string{"idle", "online", "dnd"}

	stickerList []string
	uaChrome    = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36"

	osFS        = afero.NewOsFs()
	memFS       = afero.NewMemMapFs()
	httpCache   = afero.NewHttpFs(osFS)
	httpMem     = afero.NewHttpFs(memFS)
	startTime   = time.Now()
	mem         runtime.MemStats
	totalMem    string
	HeapAlloc   string
	SysMem      string
	Frees       string
	NumGCMem    string
	timeElapsed string
	latestLog   string
	winLogs     string
	tempDirLoc  string

	// vars for openai
	// this is old and will be updated according the the latest openai's docs
	openaiapiKey       = "" // your api key here
	maxReqPerMin       = 10
	currReqPerMin      = 0
	openaiBestOf       = 1
	openaiMinRange     = float64(0.1)
	openaiMaxRange     = float64(0.9)
	openaiPresPen      = float32(0.0)
	openaiLogprobs     = 0
	openaiMaxTokens    = 200 // 100 tokens ~= 75 words
	openaiMaxInputChar = 1800
	openaiTotalTokens  = 0
	openaiModel        = ""
	notifyCreator      = false
	openAIAccess       = []string{
		"631418827841863712", // !Ararasseo~#0218
		"864622332446638142", // Entertainless#7033
	}
	re = regexp.MustCompile("[0-9]+")
)
