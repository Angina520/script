package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
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
	var str string
	fmt.Print("请先确定想选择几门课（只能为正整数）：")
	std := bufio.NewScanner(os.Stdin)
	std.Scan()
	str = std.Text()
	value, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("类型错误")
		time.Sleep(time.Hour)
		return
	}

	url, cookie := make([]string, value), make([]string, value)
	for i := 0; i < value; i++ {
		fmt.Print("请输入url：")
		std.Scan()
		url[i] = std.Text()
		fmt.Print("请输入cookie：")
		std.Scan()
		cookie[i] = std.Text()
	}
	for {
		for i := 0; i < value; i++ {
			err := httpDo(url[i], cookie[i])
			if err != nil {
				fmt.Println("错误，终止运行!请检查url和cookie")
				time.Sleep(time.Hour)
				return
			}
			time.Sleep(time.Second)
		}
		time.Sleep(time.Second)
	}
}
