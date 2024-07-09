# fabric-training-data-contract

“基于区块链的AI训练集数据共享平台 ”区块链合约部分。

---
#### How to use
```sh
go mod vendor

export FABRIC_CFG_PATH=<your fabric config path>
peer lifecycle chaincode package training_data_contract.tar.gz --path . --lang golang --label training_data
```
或者直接用 `Releases` 中我打包好的 `training_data_contract.tar.gz`，然后依照第七次实验中的方法部署合约。

---
#### API
- `CreateData`: Create or update a training data with the given ID, data and metadata.
- `QueryData`: Query the training data with the given ID.
- `QueryMetadata`: Query the metadata of the training data with the given ID.
- `QueryAllMetadatas`: Query the metadata of all training data.
- `ChangeMetadata`: Change the metadata of the training data with the given ID.
- `ChangeData`: Change the data of the training data with the given ID.

数据模型见 `types.go`。

---
#### Q&A
- Q: 如何实现查找符合条件的数据集？
- A: 用区块链存储数据保证了训练数据的不可篡改，而查询因此变得困难。小规模情况下可以考虑取出所有 `Metadata` 进行查询。大规模情况下，可以考虑维护与区块链同步的元数据数据库用于查询。

- Q: 如何在区块链中保存超级大的数据集？
- A: 做不到，但是既然已经用了区块链，不妨考虑将数据集做种，然后把种子塞到区块链上面。当然，也可以用 `IPFS`，然后塞一个哈希值进去就行了。（参考zlib？）
