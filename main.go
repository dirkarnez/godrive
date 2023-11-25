package main

import (
	"fmt"

	"golang.org/x/net/context"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

func main() {
	var serviceAccountFilePath string = "{{服務帳戶 金鑰 JSON 檔案}}"

	sharedService, err := drive.NewService(context.Background(), option.WithServiceAccountFile(serviceAccountFilePath))
	if err != nil {
		panic(err)
	}

	r, err := sharedService.Files.List().PageSize(10).Fields("nextPageToken, files(id, name)").Do()
	if err != nil {
		panic(err)
	}

	fmt.Println("Files:")
	if len(r.Files) == 0 {
		fmt.Println("No files found.")
	} else {
		for _, i := range r.Files {
			fmt.Printf("%s (%s)\n", i.Name, i.Id)
		}
	}
}
