### 证书生成步骤
```shell
# 生成根证书
$ cfssl gencert -initca ca-csr.json | cfssljson -bare ca -
# 签发server证书
$ cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=ca-config.json -profile=server server-csr.json | cfssljson -bare server
# 签发client 证书
$ cfssl gencert -ca=ca.pem -ca-key=ca-key.pem -config=ca-config.json -profile=client client-csr.json | cfssljson -bare client
```
