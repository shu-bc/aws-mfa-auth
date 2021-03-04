# aws-mfa-auth
AWS CLIを利用する上で、デフォルトのプロファイルを利用して、MFA認証を行い、認証トークンを環境変数にセットするためのコマンドを出力する便利なツール

## インストール
brew
```
brew tap shu-bc/aws-mfa-auth https://github.com/shu-bc/aws-mfa-auth
brew install aws-mfa-auth
```

## 使い方
```
$aws-mfa-auth ****** # ６桁の数字
export AWS_ACCESS_KEY_ID=
export AWS_SECRET_ACCESS_KEY=
export AWS_SESSION_TOKEN=

------------------------------------
use these commands to unset envs
unset AWS_ACCESS_KEY_ID
unset AWS_SECRET_ACCESS_KEY
unset AWS_SESSION_TOKEN
```

exportのコマンドをターミナルにコピペで実行する
unsetのコマンドは環境変数をクリアするためのもので、以下のコマンドでも取得できる

環境変数をクリアするためのコマンドを出力
```
$aws-mfa-auth -u
```
注意： 環境変数をクリアしないと、認証情報の期限切れになった時に、何もできなくなる
