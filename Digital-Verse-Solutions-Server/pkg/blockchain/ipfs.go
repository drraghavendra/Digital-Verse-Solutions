package blockchain

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/rs/zerolog/log"
)

func CreateFileUrl(cid string) string {
	return fmt.Sprintf("https://ipfs.io/ipfs/%s", cid)
}

// UploadJsonToIpfs https://github.com/nftstorage/go-client/blob/main/docs/NFTStorageAPI.md#store
func UploadJsonToIpfs(nftStorageKey string, name, description, image string) (cid string, err error) {
	nftJson := NftJson{
		Name:        name,
		Description: description,
		Image:       image,
	}
	b, err := json.Marshal(nftJson)
	if err != nil {
		log.Error().Err(err).Msg("Error converting to json")
		return
	}
	cid, err = PostToNftStorage(nftStorageKey, bytes.NewReader(b))
	if err != nil {
		log.Error().Err(err).Msg("Error uploading json to Ipfs")
		return
	}
	return
}

func PostToNftStorage(nftStorageKey string, body io.Reader) (cid string, err error) {
	req, err := http.NewRequest("POST", "https://api.nft.storage/upload", body)
	if err != nil {
		return
	}
	req.Header.Set("Authorization", "Bearer "+nftStorageKey)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error().Err(err).Msg("http.DefaultClient.Do")
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Error().Err(err).Msg("ioutil.ReadAll")
			return "", err
		}
		var response NeoStorageResponse
		json.Unmarshal(bodyBytes, &response)
		return response.Value.Cid, nil
	}
	return
}
