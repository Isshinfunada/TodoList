# TodoList

1. 背景
プロジェクトの背景、目的、期待される効果、成果を説明します
- 背景：技術力向上
- やること：以下機能を搭載したTODOリストを作成する
  - 機能
    - CRUD操作
    - ステータス管理
    - 通知機能
      - UPDATEされたら会員登録したメールアドレスに通知が届く
    - 会員登録
- 技術構成
  - フロントエンド：Vue（Nuxt.jsというフレームワーク）
  - バックエンド：Go（GoのEchoというフレームワーク）
    - DBとの接続：sqlc
  - DB：PostgreSQL
  - コンテナ：Docker
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

2. 完成イメージ
https://claude.site/artifacts/203a61a4-4d80-47bd-b1db-609a5802e4c3