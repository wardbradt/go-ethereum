rm -rf datadir

GETH=build/bin/geth

$GETH init --datadir datadir genesis.json

# dlv exec $GETH -- --datadir datadir \
#       --verbosity=0 \
#       --ws \
#       --http.api debug,personal,eth,net,web3,txpool,admin,miner \
#       --miner.etherbase=0xd912aecb07e9f4e1ea8e6b4779e7fb6aa1c3e4d8 \
#       --miner.gasprice 0 \
#       --miner.strictprofitswitch 3s \
#       --mine \
#       --miner.threads=8

$GETH --datadir datadir \
      --verbosity=0 \
      --ws \
      --http.api debug,personal,eth,net,web3,txpool,admin,miner \
      --miner.etherbase=0xd912aecb07e9f4e1ea8e6b4779e7fb6aa1c3e4d8 \
      --miner.gasprice 0 \
      --miner.strictprofitswitch 3s \
      --mine \
      --miner.threads=8

