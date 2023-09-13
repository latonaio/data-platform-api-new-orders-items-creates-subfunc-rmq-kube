# data-platform-api-orders-creates-subfunc-items-rmq-kube
data-platform-api-orders-creates-subfunc-items-rmq-kube は、データ連携基盤において、オーダーAPIサービスの明細登録更新補助機能を担うマイクロサービスです。

## 動作環境
・ OS: LinuxOS  
・ CPU: ARM/AMD/Intel  

## 対象APIサービス
data-platform-api-orders-creates-subfunc-items-rmq-kube の 対象APIサービスは次の通りです。

*  APIサービス URL: https://xxx.xxx.io/api/API_ORDERS_SRV/creates/

## 本レポジトリ が 対応する データ
data-platform-api-orders-creates-subfunc-items-rmq-kube が対応する データ は、次のものです。

* OrdersHeader（オーダー - ヘッダデータ）
* OrdersHeaderPartner（オーダー - ヘッダ取引先データ）
* OrdersHeaderPartnerPlant（オーダー - ヘッダ取引先プラントデータ）
* OrdersHeaderPartnerContact（オーダー - ヘッダ取引先コンタクトデータ）

## Output
data-platform-api-orders-creates-subfunc-items-rmq-kube では、[golang-logging-library-for-data-platform](https://github.com/latonaio/golang-logging-library-for-data-platform) により、Output として、RabbitMQ へのメッセージを JSON 形式で出力します。以下の項目のうち、"cursor" ～ "time"は、golang-logging-library-for-data-platform による 定型フォーマットの出力結果です。

```
```