package main

import (
	"fmt"
	"os"
	"github.com/slack-go/slack"
)

func main(){
	os.Setenv("SLACK_BOT_TOKEN" , "xoxb-4017635101060-4012502375733-9BnKExAosGjlqzePqOsw4vVW")
	os.Setenv("CHANNEL_ID" , "C040F271603")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"picture.jpg"}

	for i:= 0 ; i<len(fileArr) ; i++ {
		params := slack.FileUploadParameters{
			Channels : channelArr ,
			File: fileArr[i],
		}
		file ,err := api.UploadFile(params)
		if err != nil{
			fmt.Printf("%s \n" , err)
			return
		}
		fmt.Printf("Name : %s , URL : %s\n", file.Name , file.URL)
	}
}