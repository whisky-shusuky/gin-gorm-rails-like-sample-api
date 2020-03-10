# 概要
- `gin`,`gorm` を使って`golang`で`Rails`風な構成でサンプルAPIを作ったもの。

# 使用方法
- `sudo docker-compose up`でnginx,gin,mysqlが動くコンテナが立ち上がります。
  - 初回はmysqlの構築のためにginの起動に失敗するので再度立ち上げて`ctrl + c`で抜けてください。
- テーブルなどの構築は自動で行われますが勝手にレコードは入力されません。
  - `sql/mysql/seed.sql`にseedを用意しましたのでmysqlコンテナに入ってseedの中身を流し込めばサンプルレコードが作成されます。
```bash
$ sudo docker exec -it mysql bash
# mysql -uroot

Welcome to the MySQL monitor.  Commands end with ; or \g.
...

mysql> use sample
Database changed
mysql> INSERT INTO `shops` ( shop_name, shop_description) VALUES
    -> ('whisky shop','Shop of liquors.'),
    -> ('shusuky no mise', 'Sake suki oyaji no mise.');
Query OK, 2 rows affected (0.01 sec)
Records: 2  Duplicates: 0  Warnings: 0

mysql> INSERT INTO `books` ( book_name, book_description, sales) VALUES
    -> ('whisky diary','Secret diary of whisky!', null),
    -> ('shusuky diary', 'Secret diary of whisky!', 20),
    -> ('book of whisky','A book about world whisky!', 200),
    -> ('book of shusuky','A book about world nonbei people!', 20000);
Query OK, 4 rows affected (0.00 sec)
Records: 4  Duplicates: 0  Warnings: 0
```

- レコードが作成されればローカルからGETが可能になります。
```bash
$ curl localhost:10080/api/v1/shops | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   170  100   170    0     0  10000      0 --:--:-- --:--:-- --:--:-- 10000
{
  "shops": [
    {
      "id": 1,
      "shop_name": "whisky shop",
      "shop_description": "Shop of liquors."
    },
    {
      "id": 2,
      "shop_name": "shusuky no mise",
      "shop_description": "Sake suki oyaji no mise."
    }
  ]
}


$ curl localhost:10080/api/v1/books | jq
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   492  100   492    0     0  54666      0 --:--:-- --:--:-- --:--:-- 54666
{
  "books": [
    {
      "id": 1,
      "book_name": "whisky diary",
      "book_description": "Secret diary of whisky!",
      "sales": 0,
      "rank": null
    },
    {
      "id": 2,
      "book_name": "shusuky diary",
      "book_description": "Secret diary of whisky!",
      "sales": 20,
      "rank": "ちょいちょい"
    },
    {
      "id": 3,
      "book_name": "book of whisky",
      "book_description": "A book about world whisky!",
      "sales": 200,
      "rank": "そこそこ"
    },
    {
      "id": 4,
      "book_name": "book of shusuky",
      "book_description": "A book about world nonbei people!",
      "sales": 20000,
      "rank": "ベストセラー"
    }
  ]
}

```
- 同様にして以下のようにして`shops`と`books`のCRUDが試せます。
```bash
# GET
$ curl localhost:10080/api/v1/shops
$ curl localhost:10080/api/v1/shops/1
$ curl localhost:10080/api/v1/shops/1
$ curl localhost:10080/api/v1/books/1

# POST
$ curl http://localhost:10080/api/v1/shops -X POST -H "Content-Type: application/json" -d '{"ShopName": "test shop name","ShopDescription": "test shop description"}'
$ curl http://localhost:10080/api/v1/books -X POST -H "Content-Type: application/json" -d '{"BookName": "test book name","BookDescription": "test book description", "Sales": 10}'

# PUT
$ curl http://localhost:10080/api/v1/shops/4 -X PUT -H "Content-Type: application/json" -d '{"ShopName": "test shop name change","ShopDescription": "test shop description change"}'
$ curl http://localhost:10080/api/v1/books/5 -X PUT -H "Content-Type: application/json" -d '{"BookName": "test book name","BookDescription": "test book description", "Sales": 1000000}'

# DELETE
$ curl localhost:10080/api/v1/shops/4 -X DELETE
$ curl localhost:10080/api/v1/books/6 -X DELETE

```
- 

# 参照
- 
  - [日本語の記事](https://qiita.com/Asuforce/items/0bde8cabb30ac094fcb4)。
  - [深センの方が作ったサンプルAPI](https://github.com/eddycjy/go-gin-example)