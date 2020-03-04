# goroutine-channel-learning

* goroutine
* channel

上記がよく分かっていないので勉強用

## 特に確認したいこと

* context
* case

### わかったこと

* `channel`は`goroutine`の間で値を受け渡すことができる
* `wg.Add(1)`と`wg.Done()`は同じ場所で行う（というかwaitgroupは別にシングルトンなわけではないのでそれぞれで持つひつようがある）

## 各パッケージ

* server
  Listenerが常に接続を監視してクライアントが来たらConnを作成するイメージ
  * mainプロセスから閉じたときもうまいことすべてのプロセスを殺せるようになる
  * Connのプロセスが死んだだけではメインは死なない。新しい接続を受け付ける



## 参考にさせていただいたサイト

[Go の channel 処理パターン集](https://hori-ryota.com/blog/golang-channel-pattern/)

[いまさら聞けないselectあれこれ](https://www.slideshare.net/lestrrat/select-66633666)