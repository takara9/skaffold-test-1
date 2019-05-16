# skaffold-test-1

IKS (IBM Cloud Kubernetes Service) で Skaffold を利用できることを確認するためのサンプルコートです。

これは、Go言語で書いたウェブサーバーをコンテナとしてビルドして、レジストリに登録しておき、Kubernetesにデプロイする簡単なコードです。
Go言語のソースコードを変更する都度、Skaffoldが監視していて、Kubernetesへデプロイしてくれます。


## 前提条件

以下の前提条件を満たしている必要があります。

* パソコンにDockerコマンドがインストールされている
* コンテナのレジストリにアカウントがあり、docker pushでイメージを登録できる
* IKSのクラスタが作成済みであり、kubectl コマンドで操作できる


## 準備

このレポジトリを、ご自分のGitHubにフォークしてください。そして、次のように編集します。

マニフェストのimage: の<YOUR REG ID> をご自分のIDに変更します。

~~~file:k8s-webserver.yaml
    spec:
      containers:
      - name: webserver
        image: <YOUR REG ID>/skaffold-example
        ports:
        - containerPort: 8080
~~~

もう一つ、skaffoldの設定ファイルも修正します。

~~~
apiVersion: skaffold/v1beta10
kind: Config
build:
  artifacts:
  - image: <YOUR REG ID>/skaffold-example
deploy:
  kubectl:
    manifests:
    - k8s-webserver.yaml
~~~


## 操作方法

ターミナルを二つ開き一方で、skaffold dev を実行して、もう一方で　ファイルを編集します。
main.go や index.html を修正すると、ローカルでコンテナのビルド、プッシュが実行され、
次にデプロイが実行されます。 ファイルが編集される都度、実行されます。


レジストリにログインして、docker push ができる状態にしておきます。

~~~
$ docker login
~~~

<your repo name> をご自分のリポジトリのアカウントIDに置き換えて実行します。

~~~
$ skaffold dev --default-repo <your reg account id>
~~~



## 動作確認

自動的にKubernetes クラスタへデプロイされるので、以下のコマンドで、アクセスして確認します。
ワーカーノードのIPアドレスが解らないときは、`ibmcloud ks workers <CLUSTER NAME>`でパブリックIPアドレスを取得してください。

~~~
$ curl http://<YOUR NODE IP>:31080/
~~~


プログラムを変更して、保存するごとに、ビルドからデプロイまでが自動実行されます。



