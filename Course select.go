package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func httpDo(url, cookie string) error {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Cookie", cookie)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(body))
	return nil
}

func main() {
	var url, cookie string
	std := bufio.NewScanner(os.Stdin)
	fmt.Print("请输入url：")
	std.Scan()
	url = std.Text()
	fmt.Print("请输入cookie：")
	std.Scan()
	cookie = std.Text()
	for {
		err := httpDo(url, cookie)
		if err != nil {
			fmt.Println("错误，终止运行!请检查url和cookie")
			time.Sleep(3 * time.Second)
			break
		}
		time.Sleep(3 * time.Second)
	}
}
