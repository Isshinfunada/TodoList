# TodoList
## ローカル環境セットアップ
``` bash
make start
```

まだフロントしかデプロイしていないので機能が動きません。ローカルでは動きます。

https://todo-list-c6s9.vercel.app/


Contributer: Isshin, Shinon

![image](https://github.com/user-attachments/assets/9fc2bf3e-80aa-4f34-a1c6-add3ab0b9138)



1. 背景
プロジェクトの背景、目的、期待される効果、成果を説明します
- 背景：技術力向上
- やること：以下機能を搭載したTODOリストを作成する
  - 機能
    - CRUD操作
    - ステータス管理
    - 会員登録
- 技術構成
  - フロントエンド：Vue（Nuxt.jsというフレームワーク）
  - バックエンド：Go（GoのEchoというフレームワーク）
    - DBとの接続：sqlc
  - DB：PostgreSQL
  - コンテナ：Docker（Orbstack）
  - デプロイ
    - フロントエンド：Vercel
    - バックエンド：Google Cloud Platform
  - IaC(Infrastructure as Code)
    - Terraform：GCPをコードで設定できる
  - プロジェクト管理
    - Git、Github
  - オプション
    - CI
      - Github Action
      - Circle CI
    - 認証
      -Firebase Authentication

2. 完成イメージ
Claude
https://claude.site/artifacts/203a61a4-4d80-47bd-b1db-609a5802e4c3
