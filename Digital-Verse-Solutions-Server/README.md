## Presentation

Materials - https://drive.google.com/drive/folders/1YMkA0egela66tr3zzgOw7-jKlx4zDF_d

### Problem annotation
Synthetic media is a realistic transformation of audio and video using artificial intelligence. Currently, there are several applications based on this technology, but it’s developing rapidly, attracting more and more public attention.

- On one hand, this technology is a holy grail for advertisers and filmmakers which can give them endless opportunities to use any faces of any celebrities in their projects. With the help of our platform, the celebrity will be able to pronounce the text of the commercial in all the world's known languages. The advertiser will have a chance to create a separate commercial for ten thousand products, having only one digitized version of the celebrity's face. Film producers won't have to pay multimillion-dollar royalties to celebrities, it will be enough to buy their face.

- On the other hand, without proper regulation this technology is a sophisticated threat for businesses and individuals. The illegal use of faces is gaining momentum. Debates about the originality of the synthetic videos and lawsuits attract lots of attention, thereby encouraging the creation of content with celebrities without their consent. Technology is evolving fast and it’s only a matter of time before synthetic videos will be no longer distinguishable from the original.

We believe that with the help of blockchain this problem can be solved, NFT have a potential to create metaverse of the digital avatars of the users and regulate the relationship between video producers and celebrities, while copyright and data protection laws still cannot.

### Solution
Our solution is a blockchain-based NFT marketplace for the digital faces.

### [Platform architecture](https://drive.google.com/file/d/1R9-Pf1u-HHQ4s0dC99Ys2QAIVhVW7u0k/view?usp=sharing)

The user (for example celebrity) will be able to digitize his face by uploading a video of himself to the platform. Digital avatar of the user will be presented on the platform for the potential customers. After receiving the offer to buy the face, the user will be able to choose whether to sell it for this particular video/commercial. In case of consent, an NFT token gets created.

After the token is minted, potential video creators can purchase this NFT, thereby acquiring the right to use the celebrity's face to create one synthetic video.

With our solution, video creators will be able to purchase a digitized face and produce a video with this face on the same platform. And all this without the need for real filming and with digital confirmation on the blockchain. The NFT will eliminate the need to obtain IPR rights, which will be assigned to the video content producers with the purchase of the token.

## Demo access

### Demo video

https://www.youtube.com/watch?v=oqKDBIsdgkY

### Web interface

Url: https://app.digitalverse.ai/

Email: demo@digitalverse.com

Password: qwerty

### Telegram bot

@digitalverse_bot

### Support

@panichm (telegram account)

Recommends for good results: upload small videos (5-20 seconds long), keep your face away from the camera for better quality, use videos with standard lighting.

## Setup:

create .env file (contact my telegram for API keys - @panichm)

docker-compose build

docker=compose up

server runs on port -> localhost:4000

## Deployed smart contracts

- ./solidity/ - path

### Okex

- https://testnet.coinex.net/address/0x1D2A0bA44f522dE56Bf8fa83C312B8BdACfC20eB - NFT
- https://testnet.coinex.net/address/0xE89Cd3B9e0999c840BA68b9069e368A73975700d - NFT Market

Our contracts has all needed basic functions to work with NFT tokens:

- **MintTokenToAddress(address owner, string memory metadataURI)**

    Allows you to mint a NFT token.
- **TransferFrom(address from, address to, uint256 tokenId)**

    Allows you to transfer any NFT token on other address.
- **Implementation of access control**

- **NFT Market**

    Allows you to list and sell created NFT tokens. Functions:

    createMarketItem(address nftContract,uint256 tokenId,uint256 price)

    createMarketSale(address nftContract,uint256 itemId)

    fetchMarketItems() public view returns (MarketItem[] memory)

    fetchMyNFTs() public view returns (MarketItem[] memory)

## API

- **{server}/create/{name}**

    Allows you to mint a NFT token

- **{server}/upload_file_to_ipfs**

    Uploading a media file to IPFS storage

## Open Issues

-

## Roadmap

- Integrate copyrights mechanism
- Integrate NFT sales/transfer options in UI

## Help Tools

### Filecoin/IPFS

- https://nft.storage/#docs

## Blockchains

### OKExChain

- https://okexchain-docs.readthedocs.io/en/latest/developers/quick-start.html

## Other

- https://github.com/oligot/go-mod-upgrade
- go mod tidy
- https://mholt.github.io/curl-to-go/

## License
Apache License, Version 2.0
