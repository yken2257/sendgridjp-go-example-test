# sendgridjp-go-example-test [![sendgridjp-go-example-test](https://circleci.com/gh/yken2257/sendgridjp-go-example-test.svg)](https://app.circleci.com/pipelines/github/yken2257/sendgridjp-go-example-test)
[SendGridJPのGoサンプルコード](https://github.com/SendGridJP/sendgridjp-go-example)の動作確認のためのリポジトリです。

## 概要
CircleCI上でGo1.18とSendGrid公式ライブラリ(最新版)をインストールし、サンプルコードの動作検証をします。
具体的には、サンプルコードの最後でHTTPレスポンスコード202が返ってくればテスト成功とみなします。

- main.go: [サンプルコード](https://github.com/SendGridJP/sendgridjp-go-example/blob/master/src/main/main.go)をモジュール化したもの
- send_test.go: 上記をtestingでテストするコード
- .circleci/config.yml: CircleCI設定（環境設定、環境変数設定、テストののち、用いたバージョンを表示します。毎月2日の午前9時に定期実行されます。）

（手動でテストする場合の手順）

```bash
# Install dependencies
go mod init sendgridjp-go-example-test
go mod tidy
# .envファイルを編集
echo "API_KEY=$SENDGRID_API_KEY" >> .env
echo "TOS=alice@sink.sendgrid.net,bob@sink.sendgrid.net,carol@sink.sendgrid.net" >> .env
echo "FROM=you@example.com" >> .env
# Show Version
go version
cat go.sum | grep sendgrid-go
# Test
go test
```
