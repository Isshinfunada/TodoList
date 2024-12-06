# TodoList
## ローカル環境セットアップ
``` bash
make start
```

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
  - デプロイ *途中
    - フロントエンド：Vercel
    - バックエンド：Heroku
  - IaC(Infrastructure as Code)
    - Terraform
  - プロジェクト管理
    - Git、Github
  - オプション
    - CI *未搭載
      - Github Action
      - Circle CI
    - 認証 

2. 完成イメージ（プロトタイプ）
Claude
https://claude.site/artifacts/203a61a4-4d80-47bd-b1db-609a5802e4c3
