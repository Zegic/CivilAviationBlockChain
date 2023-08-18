docker stop $(docker ps -a | awk '($2 ~ /hyperledger*/) {print $1}')
docker rm $(docker ps -a | awk '($2 ~ /hyperledger*/) {print $1}')
#docker rm $(docker ps -a -q)
#docker rmi $(docker images -q)
docker volume rm $(docker volume ls -qf dangling=true)

sudo docker network prune
sudo docker volume prune


rm -rf ./channel-artifacts
rm -rf ./crypto-config