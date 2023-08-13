# webapp-learn
カンバン方式のタスク管理webアプリケーションです.
Go言語でのバックエンド開発を中心に, Nuxt3を利用したフロントエンド制作とともに学習しました。

https://github.com/pisces2336/webapp-learn/assets/98950347/90109b44-f205-48c0-946e-64c512db4d4b

## 起動方法
- `docker compose up`をターミナル上で実行してください.
- 数分待ち, ターミナルに下のように表示が出たことを確認した後, `localhost:3000/kanbans`にアクセスしてください.
    - なお, データベースの立ち上がりまではapiやwebのコンテナでエラーを吐き続けますが問題ありません.
```
Nuxi 3.6.1
wNuxt 3.6.1 with Nitro 2.5.2

  > Local:    http://localhost:3000/
  > Network:  http://172.21.0.4:3000/


✔ Nuxt Devtools is enabled v0.6.4 (experimental)

 WARN  Slow module @nuxt/devtools took 21021.9ms to setup.


ℹ Vite client warmed up in 4989ms
[nitro] ✔ Nitro built in 2980 ms
```

## 使用技術
- フロントエンド
    - nuxt 3.6.1
    - bootstrap 5.3.0
- バックエンド (api)
    - golang 1.20.5
    - Echo v4
    - gorm v1
- データベース
    - mysql 8.0.33
- その他
    - Docker, Docker-compose
    - github

## 機能一覧
- カンバン作成 (タイトル, 本文)
    - 新規作成したカンバンは未着手カテゴリに振り分けられます.
    - Fetch APIを用いてバックエンドと通信し, データベースへ保存を行います.
- カンバン削除
    - 各カンバンに表示された赤色のボタンを押下することで削除が行えます.
    - Fetch APIを用いてバックエンドと通信し, データベースから該当のカンバンのデータを削除します.
- カテゴリ間移動
    - 未着手 ⇔ 作業中 ⇔ 完了 の移動を, 各カンバンに表示された矢印ボタンで行えます.
    - Fetch APIを用いてバックエンドと通信し, データベースから該当のカンバンのデータを更新します.

## 工夫, 苦労した点
- docker-composeの利用
    - ローカル環境をなるべく汚すことなく開発を行うことが出来た上, githubからcloneした際の起動方法も簡単にすることが出来ました.
- bootstrapの導入
    - バックエンドの学習に力を入れたかったため, bootstrapを導入することで画面デザインを高速に行うことが出来るようにしました.
- フロントエンド, バックエンド間の通信
    - docker-composeを使用していることで名前解決が特殊な形で行われており, サーバーサイドとブラウザでの差異の解決に苦労しました (docker-composeファイルでのextra_hostsの指定で解決).
