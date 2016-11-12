# goPacketCapture
Go言語でパケットキャプチャする

## 動作確認済み環境

Vagrant 1.8.1
CentOS Linux release 7.2.1511 (Core)
go version go1.6.3 linux/amd64


## トラブルと対処方法

### find_deviceでデバイスが表示されない

権限が足りてない可能性があるので
`sudo ./find_device`
にて実行する。

# 参考
[Goでパケットキャプチャを実践してみる](http://qiita.com/kkyouhei/items/846e74c6a9653b069e5f)
