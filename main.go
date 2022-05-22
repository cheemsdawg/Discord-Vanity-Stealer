package main

import (
	"fmt"
  "net/http"
  "bytes"
  "runtime"
  "time"
)

var (
  token string = "" // your token, Note: add 'Bot ' prefix if stealing with bot
  keepguild string = "" // server id of the server you stealing from
  removeguild string = "" // server id of the server where you keeping the vanity
)


func main () {
  runtime.GOMAXPROCS(4)
  go func () {   
    client := &http.Client{}
    var json  = []byte(`{"code":"abcxyz"}`) // abcxyz is the vanity you want in the server you stealing from
    ses, err := http.NewRequest("PATCH", fmt.Sprintf("https://discord.com/api/v9/guilds/%s/vanity-url", removeguild), bytes.NewBuffer(json))
    ses.Header.Set("authorization", token)
    ses.Header.Set("Content-Type", "application/json")
    if err != nil {
      fmt.Println(err)
    }
    doer, err := client.Do(ses)
    if doer.StatusCode == 200 {
      fmt.Println(fmt.Sprintf("Successful remove vanity attempt | StatusCode: %d", doer.StatusCode))
    } else {
      fmt.Println(fmt.Sprintf("Unsuccessful remove vanity attempt | StatusCode: %d", doer.StatusCode))
    }

    }() 
  
  time.Sleep(time.Duration(150) * time.Millisecond)
  
  go func () {
    client := &http.Client{}
    var jsonn  = []byte(`{"code":"abcxyz"}`) // abcxyz is the vanity you stealing
    ses2, err := http.NewRequest("PATCH", fmt.Sprintf("https://discord.com/api/v9/guilds/%s/vanity-url", keepguild), bytes.NewBuffer(jsonn))
    ses2.Header.Set("authorization", token)
    ses2.Header.Set("Content-Type", "application/json")
    if err != nil {
      fmt.Println(err)
    }
    doer, err := client.Do(ses2)
    if doer.StatusCode == 200 {
      fmt.Println(fmt.Sprintf("Successful put vanity attempt | StatusCode: %d", doer.StatusCode))
    } else {
      fmt.Println(fmt.Sprintf("Unsuccessful put vanity attempt | StatusCode: %d", doer.StatusCode))
    }

  }()
  time.Sleep(12 * time.Second)
}
