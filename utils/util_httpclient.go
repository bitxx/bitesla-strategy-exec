package utils

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

const (
	GetMethod    = "GET"
	PostMethod   = "POST"
	SendtypeFrom = "from"
	SendtypeJson = "json"
	ProxyUrl     = "localhost:1086"
	//ProxyUrl = "host.docker.internal:1086"
)

//默认为SendtypeFrom
type HttpSend struct {
	Link     string
	SendType string
	Header   map[string]string
	Body     interface{}
	sync.RWMutex
}

func NewHttpSend(link string, typ string) *HttpSend {
	return &HttpSend{
		Link:     link,
		SendType: typ,
	}
}

func (h *HttpSend) SetBody(body interface{}) {
	h.Lock()
	defer h.Unlock()
	h.Body = body
}

func (h *HttpSend) SetHeader(header map[string]string) {
	h.Lock()
	defer h.Unlock()
	h.Header = header
}

func (h *HttpSend) SetSendType(sendType string) {
	h.Lock()
	defer h.Unlock()
	h.SendType = sendType
}

func (h *HttpSend) Get(needProxy bool) ([]byte, error) {
	return h.send(GetMethod, needProxy)
}

func (h *HttpSend) Post(needProxy bool) ([]byte, error) {
	return h.send(PostMethod, needProxy)
}

func GetUrlBuild(link string, data map[string]string) string {
	u, _ := url.Parse(link)
	q := u.Query()
	for k, v := range data {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String()
}

func (h *HttpSend) send(method string, needProxy bool) ([]byte, error) {
	var (
		req      *http.Request
		resp     *http.Response
		client   http.Client
		sendData string
		err      error
	)

	if h.Body != nil {
		if strings.ToLower(h.SendType) == SendtypeJson {
			sendBody, jsonErr := json.Marshal(h.Body)
			if jsonErr != nil {
				return nil, jsonErr
			}
			sendData = string(sendBody)
		} else {
			formBody := h.Body.(map[string]string)
			sendBody := http.Request{}
			sendBody.ParseForm()
			for k, v := range formBody {
				sendBody.Form.Add(k, v)
			}
			sendData = sendBody.Form.Encode()
		}
	}

	if needProxy {
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			Proxy: func(req *http.Request) (*url.URL, error) {
				return &url.URL{
					Scheme: "socks5",
					Host:   ProxyUrl}, nil
			},
			Dial: (&net.Dialer{
				Timeout: 10 * time.Second,
			}).Dial,
		}
	} else {
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			Dial: (&net.Dialer{
				Timeout: 10 * time.Second,
			}).Dial,
		}
	}

	client.Timeout = time.Duration(10 * time.Second)

	req, err = http.NewRequest(method, h.Link, strings.NewReader(sendData))
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	//设置默认header
	if len(h.Header) == 0 {
		//json
		if strings.ToLower(h.SendType) == SendtypeJson {
			h.Header = map[string]string{
				"Content-Type": "application/json; charset=utf-8",
			}
		} else { //form
			h.Header = map[string]string{
				"Content-Type": "application/x-www-form-urlencoded",
			}
		}
	}

	for k, v := range h.Header {
		if strings.ToLower(k) == "host" {
			req.Host = v
		} else {
			req.Header.Add(k, v)
		}
	}

	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("error http code :%d", resp.StatusCode))
	}

	return ioutil.ReadAll(resp.Body)
}
