# battle-bot

DiscordのBattle botです。

Rumble loyalの日本語ver.を想定して作成しています。

## bot名

Battle Royale

## Bot管理コンパネ

以下のURLでコンパネにアクセスできます。

[コンパネURL](https://discord.com/developers/applications/975019338534899733/information)

## 導入方法

以下のURLでサーバーに追加することができます。

[サーバー追加URL](https://discord.com/api/oauth2/authorize?client_id=975019338534899733&permissions=275146427456&scope=bot)

## 権限

以下の権限を許可しています。

- GENERAL PERMISSIONS
    - Manage Roles（ロールの管理）
    - Read Messages/View Channels（メッセージを見る）
- TEXT PERMISSIONS
    - Send Messages（メッセージを送信）
    - Send Messages in Threads（スレッドでメッセージを送信）
    - Embed Links（埋め込みリンク）
    - Read Message History（メッセージ履歴を読む）
    - Add Reaction（リアクションの追加）

## botコマンド

- そのチャンネルのみで配信する場合

```
b
```

- 別チャンネルで途中経過を配信する場合

※該当チャンネルの書き込み権限を確認してください

```
b <チャンネルリンク>
```

- バトルを停止する場合

※起動したチャンネルで以下のコマンドを実行します

```
stopb
```

- 起動中のバトルを表示する場合

```
processb
```

- 新規バトルを中止する場合

※起動中のバトルは中断されません

※フラグを戻すコマンドは用意していないため、コンテナを起動し直す必要があります

```
rejectstartb
```

## リリース前チェックシート

### 正常

#### 1回目

- [ ] 通常起動ができる(b)

#### 2回目

- [ ] No Entryメッセージが送信できる

#### 3回目

- [ ] 開始後にstopbコマンドが実行できる

#### 4回目

- [ ] カウントダウンの途中でstopbコマンドが実行できる
- [ ] エントリーメッセージ送信後にstopbコマンドが実行できる
- [ ] バトルメッセージの途中でstopbコマンドが実行できる
- [ ] processbコマンドでプロセス削除できる
- [ ] rejectstartbコマンドで新規起動を停止できる
- [ ] rejectstartbコマンド実行後にアップデートで解除できる
- [ ] 複数サーバーで同時に起動できる

### 異常

- [ ] エントリーメッセージ削除時にエラーメッセージが送信される
- [ ] 重複起動時にエラーメッセージを送信できる
- [ ] 不正な引数入力時にエラーメッセージが表示できる
- [ ] 別チャンネルで重複起動した時にエラーメッセージが表示できる