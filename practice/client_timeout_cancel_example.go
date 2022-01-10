package practice

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	number := rand.Intn(2)
	fmt.Println("server get req:", number)
	if number == 0 {
		time.Sleep(time.Second * 2)
		fmt.Fprintf(w, "slow response")
		return
	}
	fmt.Fprintf(w, "quick response")
}
func startServer() {
	http.HandleFunc("/", indexHandler)
	err := http.ListenAndServe(":8888", nil)
	fmt.Printf("server err:%v \n", err)
	if err != nil {
		fmt.Println("error:", err)
		panic(err)
	}
	fmt.Println("server close")
}

type respData struct {
	resp *http.Response
	err  error
}

func doCall(ctx context.Context) {
	// 请求频繁可定义全局的client对象并启用长链接
	// 请求不频繁使用短链接
	transport := http.Transport{DisableKeepAlives: true}
	client := http.Client{Transport: &transport}
	respChan := make(chan *respData)
	req, err := http.NewRequest("GET", "http://127.0.0.1:8888/", nil)
	if err != nil {
		fmt.Printf("new request error:%v \n", err)
		return
	}
	req = req.WithContext(ctx)
	var wg sync.WaitGroup
	wg.Add(1)
	defer wg.Wait()

	// 发出请求
	go func() {
		resp, err := client.Do(req)
		fmt.Printf("client.do resp:%v, err:%v\n", resp, err)
		rd := &respData{
			resp: resp,
			err:  err,
		}
		respChan <- rd
		wg.Done()
	}()

	select {
	case <-ctx.Done():
		fmt.Println("call api timeout")
	case result := <-respChan:
		fmt.Println("call api success")
		if result.err != nil {
			fmt.Printf("call api faild err:%v \n", result.err)
			return
		}
		defer func(Body io.ReadCloser) {
			err = Body.Close()
			if err != nil {

			}
		}(result.resp.Body)
		all, err := ioutil.ReadAll(result.resp.Body)
		if err != nil {
			return
		}
		fmt.Printf("resp:%v", string(all))
	}
}

func ClientTimeoutCancel() {
	go startServer()
	time.Sleep(time.Second * 2)
	cancel, cancelFunc := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancelFunc()
	doCall(cancel)
}
