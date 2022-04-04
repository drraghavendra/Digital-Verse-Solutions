package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"gitlab.com/digitalverse/blockchain/pkg/blockchain"
)

func getCreateNftRequest(ctx *gin.Context) (createNftRequest CreateNFTRequest, err error) {
	if err = ctx.BindJSON(&createNftRequest); err != nil {
		log.Error().Err(err).Msg("Failed to encode CreateNFTRequest")
		return
	}

	if len(createNftRequest.Name) == 0 {
		createNftRequest.Name = "Digital Verse"
	}
	if len(createNftRequest.Description) == 0 {
		createNftRequest.Description = "Celebrity video"
	}
	if len(createNftRequest.URL) == 0 {
		log.Error().Msg("File url not specified")
		return createNftRequest, fmt.Errorf("file url not specified")
	}
	return createNftRequest, nil
}

func createRouter(downloader blockchain.FileDownloader, registry map[string]blockchain.Network) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/create/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		n, ok := registry[name]
		if !ok {
			log.Error().Msg("network not found")
			ctx.JSON(http.StatusNotFound, gin.H{"err": "network not found"})
			return
		}

		createReq, err := getCreateNftRequest(ctx)
		if err != nil {
			log.Error().Err(err).Msg("Failed to get params from request")
			ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}

		f, err := downloader(createReq.URL)
		if err != nil {
			log.Error().Err(err).Msg("error while download file")
			ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}
		defer f.Close()
		defer os.Remove(f.Name())

		ipfsCid, err := blockchain.PostToNftStorage(bc.NftStorageKey, f)
		if err != nil {
			log.Error().Err(err).Msg("Failed to upload file to IPFS")
			ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}

		// Make json and upload
		metadataURI, err := blockchain.UploadJsonToIpfs(
			bc.NftStorageKey,
			createReq.Name,
			createReq.Description,
			ipfsCid,
		)
		if err != nil {
			log.Error().Err(err).Msg("Error uploading json to Ipfs")
			ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}

		resp, err := n.Mint(metadataURI)
		if err != nil {
			log.Error().Err(err).Msg("Failed to mint NFT")
			ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}

		ctx.JSON(200, resp)
	})

	r.POST("/upload_file_to_ipfs", func(ctx *gin.Context) {
		fileUrl := ctx.PostForm("fileUrl")

		f, err := downloader(fileUrl)
		if err != nil {
			log.Error().Err(err).Msg("error while download file")
			ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}
		defer f.Close()
		defer os.Remove(f.Name())

		cid, err := blockchain.PostToNftStorage(bc.NftStorageKey, f)
		if err != nil {
			log.Error().Err(err).Msg("Failed to upload file to IPFS")
			ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}

		ctx.JSON(200, gin.H{
			"cid":     cid,
			"fileUrl": blockchain.CreateFileUrl(cid),
		})
	})

	return r
}
