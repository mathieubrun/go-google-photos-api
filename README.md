# go-google-photos-api

I got frustrated not having official google photos api, so I code generated mine, based on the last version of the discovery document : https://photoslibrary.googleapis.com/$discovery/rest?version=v1

Client code based on googleapis last version, before it got removed : https://github.com/googleapis/google-api-go-client/blob/10267775243d8c189ce96dc29556e9673459e6f3/photoslibrary/v1/photoslibrary-gen.go

## using official generator

After finishing writing this generator I found out google already made one.

```` sh
go install google.golang.org/api/google-api-go-generator
~/go/bin/google-api-go-generator -api_json_file photoslibrary/v1/photoslibrary-api.json -gendir .
````

## usage

```` go

import (
	"fmt"
	"github.com/mathieubrun/go-google-photos-api/v1"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"os"
)

func main() {

    credentials, err := ioutil.ReadFile("credentials.json")
    if err != nil {
        fmt.Fprintf(os.Stderr, "Go to https://developers.google.com/photos/library/guides/get-started")
        fmt.Fprintf(os.Stderr, "- click 'ENABLE THE GOOGLE PHOTOS LIBRARY API'")
        fmt.Fprintf(os.Stderr, "- choose 'Other'")
        fmt.Fprintf(os.Stderr, "- save the file")
        return nil, fmt.Errorf("Error reading credentials file: %v", err)
    }

    config, err := google.ConfigFromJSON(credentials, photoslibrary.PhotoslibraryReadonlyScope)
    if err != nil {
        return nil, fmt.Errorf("Error fetching googleapis configuration: %v", err)
    }

    tokenFileName := "token.json"
    token, err := readToken(tokenFileName)
    if err != nil {
        token, err = getToken(config)
        if err != nil {
            return nil, fmt.Errorf("Error getting token: %v", err)
        }

        err = saveToken(tokenFileName, token)
        if err != nil {
            return nil, fmt.Errorf("Error saving token file: %v", err)
        }
    }
    client := config.Client(context.Background(), token)
    service, err := photoslibrary.New(client)
    if err != nil {
        return nil, fmt.Errorf("Error creating photoslibrary client: %v", err)
    }
}
````
