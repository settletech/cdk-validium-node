GOBIN=/home/nachofq/Documents/Settle/plasma/repos/plasma-research/submodules/cdk-validium-node/dist 
CGO_ENABLED=0 
GOOS=linux 
GOARCH=amd64 
LDFLAGS="all=-X 'github.com/0xPolygonHermez/zkevm-node.Version=v0.6.5+cdk-4-g9ff24cff' -X 'github.com/0xPolygonHermez/zkevm-node.GitRev=9ff24cff' -X 'github.com/0xPolygonHermez/zkevm-node.GitBranch=HEAD' -X 'github.com/0xPolygonHermez/zkevm-node.BuildDate=Wed, 27 Nov 2024 12:26:38 -0300'"

COMMAND="run"
OPTIONS_SEQUENCER="--network custom --custom-network-file ../../../../src/kubernetes/deployData/contracts/output/genesis.json --cfg ../../../../src/kubernetes/deployData/infra/config/node.config.toml --components sequencer --http.api eth,net,debug,zkevm,txpool,web3"
OPTIONS_SYNCHRONIZER="--network custom --custom-network-file ../../../../src/kubernetes/deployData/contracts/output/genesis.json --cfg ../../../../src/kubernetes/deployData/infra/config/node.config.toml --components synchronizer"
OPTIONS_SEQUENCE_SENDER="--network custom --custom-network-file ../../../../src/kubernetes/deployData/contracts/output/genesis.json --cfg ../../../../src/kubernetes/deployData/infra/config/node.config.toml --components sequence-sender"
PROGRAM="../../cmd"

#go run $PROGRAM $COMMAND $OPTIONS # RUN
dlv debug $PROGRAM --headless --listen=:2345 --api-version=2 -- $COMMAND $OPTIONS_SYNCHRONIZER # DEBUG


# ARCHIVE
#go run -ldflags "$LDFLAGS" $PROGRAM $COMMAND $OPTIONS # RUN