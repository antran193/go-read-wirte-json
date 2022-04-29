package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type RECORDS struct {
	RECORDS []record `json:"RECORDS"`
}
type record struct {
	Gid string `json:"gid"`
}

func main() {
	readJson()

}

func readJson() {
	jsonFile, err := os.Open("road_ori.json")
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var record RECORDS
	json.Unmarshal(byteValue, &record)
	// listGid := make([]string, len(record.RECORDS))
	listGid := []string{}

	for i := 0; i < len(record.RECORDS); i++ {
		listGid = append(listGid, record.RECORDS[i].Gid)
		// fmt.Println(record.RECORDS[i].Gid)
	}
	checkGid(listGid)

}

func checkGid(listGid []string) {
	fmt.Println("check gid start ...")

	jsonFile, err := os.Open("road_bk.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var record RECORDS
	json.Unmarshal(byteValue, &record)

	listGidNoMatch := []string{}
	for i := 0; i < len(record.RECORDS); i++ {
		s := record.RECORDS[i].Gid
		count := 0
		fmt.Println("value current", record.RECORDS[i].Gid, i)
		for j := 0; j < len(listGid); j++ {
			if listGid[j] == s {
				count++
				break
			}
		}
		if count == 0 {
			fmt.Println("value not have", record.RECORDS[i].Gid, i)
			listGidNoMatch = append(listGidNoMatch, record.RECORDS[i].Gid)
		}
	}
	// fmt.Println(listGidNoMatch)
	strResult := "'" + listGidNoMatch[0] + "'"
	for i := 1; i < len(listGidNoMatch); i++ {
		strResult = strResult + ", '" + listGidNoMatch[i] + "'"
	}

	writeFile(strResult)
}

func writeFile(strResult string) {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(strResult)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
