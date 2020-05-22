package main

import (
	"fmt"
	"testing"
)

func TestInitialFetch(t *testing.T) {
	url := "https://play.google.com/store/apps/collection/cluster?clp=igM6ChkKEzgyMDQ2OTkzNjYyNDAwMTk3MDQQCBgDEhsKFWNvbS52encuaHNzLm15dmVyaXpvbhABGAMYAQ%3D%3D:S:ANO1ljK_y6A&gsr=Cj2KAzoKGQoTODIwNDY5OTM2NjI0MDAxOTcwNBAIGAMSGwoVY29tLnZ6dy5oc3MubXl2ZXJpem9uEAEYAxgB:S:ANO1ljLQ2zk&gl=US"
	appInfo, err := retrieveAppInfoFromPlayStore(url)
	if err != nil {
		fmt.Println(err)
		t.Fatal(err)
	}
	// fmt.Printf("appinfo: %v\n", appInfo)
	if csvFormatted, err := formatCsv(appInfo); err != nil {
		fmt.Println(err)
		t.Fatal(err)
	} else {
		fmt.Printf("csv:\n%s\n", csvFormatted)
	}
}
