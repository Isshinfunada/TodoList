-- init.sql

-- データベーススキーマの作成
CREATE TABLE todos (
  id SERIAL PRIMARY KEY,
  text VARCHAR(255) NOT NULL
);

-- 初期データの挿入
INSERT INTO todos (text) VALUES ('最初のTODO');
INSERT INTO todos (text) VALUES ('二番目のTODO');
INSERT INTO todos (text) VALUES ('三番目のTODO');
