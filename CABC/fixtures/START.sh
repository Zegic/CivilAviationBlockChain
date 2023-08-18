docker stop $(docker ps -a | awk '($2 ~ /hyperledger*/) {print $1}')
docker rm $(docker ps -a | awk '($2 ~ /hyperledger*/) {print $1}')
#docker rm $(docker ps -a -q)
#docker rmi $(docker images -q)
docker volume rm $(docker volume ls -qf dangling=true)

sudo docker network prune
sudo docker volume prune


rm -rf ./channel-artifacts
rm -rf ./crypto-config


cryptogen generate --config=crypto-config.yaml


configtxgen -profile TwoOrgsOrdererGenesis -outputBlock ./channel-artifacts/genesis.block -channelID system-channel

configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID mychannel

configtxgen -outputAnchorPeersUpdate ./channel-artifacts/BaiyiMSPanchors.tx -profile TwoOrgsChannel -channelID mychannel -asOrg BaiyiMSP

configtxgen -outputAnchorPeersUpdate ./channel-artifacts/AirChinaMSPanchors.tx -profile TwoOrgsChannel -channelID mychannel -asOrg AirChinaMSP

configtxgen -outputAnchorPeersUpdate ./channel-artifacts/OtherAirlineMSPanchors.tx -profile TwoOrgsChannel -channelID mychannel -asOrg OtherAirlineMSP

#cp -r ./crypto-config/* ../explorer/organizations
#build the explorer
docker-compose up -d
echo -n "o"
sleep 1
echo -n "o"
sleep 1
echo -n "o"
sleep 1
docker ps -a
