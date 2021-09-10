package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

type Info struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Cid    int    `json:"cid"`
	Pid    int    `json:"pid"`
	Weight int    `json:"weight"`
}

type Result struct {
	Data []Info `json:"data"`
}

func main() {
	bytes, err := ioutil.ReadFile("/Users/tim/dev/zhiqu/Demo/demoExportRegins/regins.json")
	if err != nil {
		fmt.Println("读取json文件失败", err)
		return
	}
	u := &Result{}
	err = json.Unmarshal(bytes, u)
	if err != nil {
		fmt.Println("解析数据失败", err)
		return
	}

	mj, _ := json.Marshal(u.Data)

	data := []Info{}
	json.Unmarshal(mj, &data)

	pid := 1017
	list := make(map[string][]map[string]string)

	// 省份map
	prosMap := make(map[string]string)
	proSlice := make([]map[string]string, 0)
	for _, v := range data {
		if pid == v.Pid {
			prosMap[strconv.Itoa(v.Id)] = v.Name

			pro := make(map[string]string)
			pro[strconv.Itoa(v.Id)] = v.Name
			proSlice = append(proSlice, pro)
		}

	}
	if len(proSlice) > 0 {
		list[strconv.Itoa(pid)] = proSlice
	}

	for k, _ := range prosMap {
		// 市
		city, cityIds := parseJsonInfo(k, data)
		if len(city) > 0 {
			list[k] = city
		}

		for _, vc := range cityIds {
			direct, _ := parseJsonInfo(vc, data)
			if len(direct) > 0 {
				list[vc] = direct
			}

		}
	}

	pjm, _ := json.Marshal(list)
	fmt.Println(string(pjm))

	if err := ioutil.WriteFile("/Users/tim/dev/zhiqu/Demo/demoExportRegins/china.json", pjm, 077); err != nil {
		panic(err)
	}

	fmt.Println("china.json write success")

}

func parseJsonInfo(pid string, list []Info) ([]map[string]string, []string) {
	data := make([]map[string]string, 0)
	ids := make([]string, 0)
	for _, v := range list {
		p, _ := strconv.Atoi(pid)
		if p == v.Pid {
			data = append(data, map[string]string{
				strconv.Itoa(v.Id): v.Name,
			})

			ids = append(ids, strconv.Itoa(v.Id))
		}
	}

	return data, ids
}
