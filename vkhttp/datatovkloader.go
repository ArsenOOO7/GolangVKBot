package vkhttp

/*import (
	"../userdata"
	"../util"
	"encoding/json"
	"fmt"
)

func LoadImage(filepath string, peerId int64) *userdata.Attachment  {
	request := VKRequest{
		Method: "photos.getMessagesUploadServer",
		Params: append(ParseConfig(&util.DefaultConfig), VKParam{
			Key:   "peer_id",
			Value: fmt.Sprintf("%v", peerId),
		}),
	}

	response := SendRequest(request)
	type UploadDataResponse struct {
		AlbumId int `json:"album_id"`
		UploadUrl string `json:"upload_url"`
		UserId int64 `json:"user_id"`
	}
	updata := &UploadDataResponse{}
	_ = json.Unmarshal(response.Response, &updata)

	fmt.Println(updata)

	return nil
}*/
