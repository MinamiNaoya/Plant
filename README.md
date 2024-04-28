# 植物の管理をするWebアプリケーション
1. 植物の情報（水やり頻度、肥料、最低気温（耐寒）、最高気温（耐熱）、生息地、購入日、夏型か冬型）といった情報をsqlite3で登録する。
2. 気象庁のAPIから取得した気温や降水量などをjsonで取得
3. 植物ごとの管理アルゴリズムをつくって（事前に？もしくはAIなど？）それにそってその植物を家の中にいれるかなどを管理する。

## 必要な条件
sqlite3のパッケージがC言語で開発されているためGCCをインストールする必要がある。
https://github.com/mattn/go-sqlite3/blob/master/README.md#windows
