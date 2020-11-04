## MEV-geth

This is a fork of go-ethereum, [the original README is here](README.original.md).

Flashbots is a research and development organization formed to mitigate the negative externalities the miner-extractable value (MEV) crisis poses to smart-contract-based blockchains by creating an open-entry, transparent and fair marketplace for MEV extraction, starting with Ethereum.

To fix this problem, we have designed and implemented a PoC for permissionless MEV extraction as a fork of geth. It is a sealed-bid block space auction mechanism that aims to obviate the use of frontrunning and backrunning techniques.

### Design goals
* Permissionless: allow anyone to participate in MEV extraction
* Pre-trade privacy: keep MEV extraction requests hidden until mined
* MEV Coverage: cover all known classes of measurable MEV extraction strategies
* Compatibility: maintain full compatibility with the network
* Security: avoid impacting current node operation
* Efficiency: minimize latency


### How it works
MEV-geth implements a new RPC method to accept bundles of ethereum transactions over https.

Once received, the bundles are validated and placed in a local bundle pool. Bundles in the bundle pool are not be gossiped to the rest of the network.

MEV-geth then selects the bundle that offers the highest payment to the miner and includes it at the beginning of a block template. The remainder of the block template is then filled with transactions from the mempool.
In parallel, the client produces a ‘normal’ block template based on the regular transaction pool.

Finally, MEV-geth compares the revenue of each template (normal vs MEV-geth) and starts mining on the most profitable one. This last step ensures that at worst, miners running MEV-geth end up with the status quo.

### How to use as a searcher
Anyone interested in extracting MEV can do so by sending bundles of transactions to MEV-geth clients. We call this role being a searcher. A searcher’s job is to monitor the ethereum state and transaction pool for MEV opportunities and produce flashbots bundles that extract that MEV. A flashbots bundle is composed of an array of valid ethereum transactions, a blockheight, and a timestamp range over which the bundle is valid.

```jsonc
{
    "signedTransactions": ['...'],
    "blocknumber": "0x386526",
    "minTimestamp": 12345, // optional
    "maxTimestamp": 12345 // optional
}
```

MEV-geth miners select the most valuable bundle and place it at the beginning of the block template at the given blockheight. Miners determine the value of a bundle by adding how much ETH was sent to the coinbase to the amount of ETH spent on gas.

To submit a bundle, the searcher sends the bundle directly to the miner using the rpc method eth_sendBundle. In the near term, a public registry of trusted miners will be maintained by the flashbots core dev team.