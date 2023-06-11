# go-ga-toppages

Generate markdown style top pages report based on Google Analytics(GA4) result for Static Site Generator by Go.

Python version is available: [hurutoriya/ga-top-pages-py](https://github.com/hurutoriya/ga-top-pages-py)

## Motivation

Static Site Generator (SSG) is not support data aggregation feature but we want to provide this feature even if we use SSG!

## Environment

This script assume

- Using Google Analytics (GA4)
- Installed Go to run this tools.
- [Hugo](https://gohugo.io/) style' SSG structure

## Show top page result

This package use `Google Analytics Data API v1`.
Please follow official docs to use this API.

[API Quickstart  \|  Google Analytics Data API  \|  Google for Developers](https://developers.google.com/analytics/devguides/reporting/data/v1/quickstart-client-libraries#go)

Please replace GOOGLE_APPLICATION_CREDENTIALS and every options based on your expected result.
Make sure you can modify Dimensions and Metrics based on your purpose.

```shell
GOOGLE_APPLICATION_CREDENTIALS="XXX.json" go run main.go -property_id 12345 -pages_root_url https://shunyaueta.com -site_content_path ../hurutoriya.github.io/content  -top_n 15

## 直近一年間の人気記事 Top15

1. `3676` views: [Pythonで変数を挿入してSQLクエリを柔軟に構築する](https://shunyaueta.com/posts/2021-04-29/)
1. `1375` views: [Twitter の検索システム、Earilybirdの論文を読む](https://shunyaueta.com/posts/2023-04-29-0030/)
1. `1166` views: [Streamlit でアップロードしたファイルのパスを取得して、特定の処理をする](https://shunyaueta.com/posts/2021-07-08/)
1. `915` views: [タスク管理ツールを Todoist から TickTick へ試しに乗り換えてみた](https://shunyaueta.com/posts/2022-03-13/)
1. `893` views: [How to get the uploaded file path and processing its file in  Streamlit](https://shunyaueta.com/posts/2021-07-09/)
1. `893` views: [愛用しているツールを更新: Joplin→Obsidian & TickTick → Todoist](https://shunyaueta.com/posts/2022-06-03-2133/)
1. `647` views: [kubernetes デプロイ時に `MatchExpressions:[]v1.LabelSelectorRequirement(nil)}: field is immutable` エラーが出た際に対処方法](https://shunyaueta.com/posts/2021-12-02/)
1. `527` views: [Twitter の検索システムを学ぶ - 概要編](https://shunyaueta.com/posts/2023-04-17-2252/)
1. `526` views: [Elasticsearch 8.4 から利用可能な従来の検索機能と近似近傍探索を組み合わせたハイブリッド検索を試す](https://shunyaueta.com/posts/2022-10-29-2337/)
1. `509` views: [Jupyter Notebook上にTensorboard を わずか2行で表示させる](https://shunyaueta.com/posts/2019-09-25/)
1. `467` views: [Standard SQLで 列と列の組み合わせの数を集計したい](https://shunyaueta.com/posts/2021-02-09/)
1. `410` views: [Elasticsearchの近似近傍探索を使って、ドラえもんのひみつ道具検索エンジンを作ってみた](https://shunyaueta.com/posts/2022-10-23-2344/)
1. `393` views: [OSSのアノテーションツール Label Studio を使って、快適にアノテーションする](https://shunyaueta.com/posts/2022-01-09/)
1. `387` views: [社内でデータ分析結果を可視化・共有する際に Google Colab が便利](https://shunyaueta.com/posts/2022-05-10-2200/)
```

Example: Following result is how this script's result render in markdown.

---

## 直近一年間の人気記事 Top15

1. `3676` views: [Python で変数を挿入して SQL クエリを柔軟に構築する](https://shunyaueta.com/posts/2021-04-29/)
1. `1375` views: [Twitter の検索システム、Earilybird の論文を読む](https://shunyaueta.com/posts/2023-04-29-0030/)
1. `1166` views: [Streamlit でアップロードしたファイルのパスを取得して、特定の処理をする](https://shunyaueta.com/posts/2021-07-08/)
1. `915` views: [タスク管理ツールを Todoist から TickTick へ試しに乗り換えてみた](https://shunyaueta.com/posts/2022-03-13/)
1. `893` views: [How to get the uploaded file path and processing its file in Streamlit](https://shunyaueta.com/posts/2021-07-09/)
1. `893` views: [愛用しているツールを更新: Joplin→Obsidian & TickTick → Todoist](https://shunyaueta.com/posts/2022-06-03-2133/)
1. `647` views: [kubernetes デプロイ時に `MatchExpressions:[]v1.LabelSelectorRequirement(nil)}: field is immutable` エラーが出た際に対処方法](https://shunyaueta.com/posts/2021-12-02/)
1. `527` views: [Twitter の検索システムを学ぶ - 概要編](https://shunyaueta.com/posts/2023-04-17-2252/)
1. `526` views: [Elasticsearch 8.4 から利用可能な従来の検索機能と近似近傍探索を組み合わせたハイブリッド検索を試す](https://shunyaueta.com/posts/2022-10-29-2337/)
1. `509` views: [Jupyter Notebook 上に Tensorboard を わずか 2 行で表示させる](https://shunyaueta.com/posts/2019-09-25/)
1. `467` views: [Standard SQL で 列と列の組み合わせの数を集計したい](https://shunyaueta.com/posts/2021-02-09/)
1. `410` views: [Elasticsearch の近似近傍探索を使って、ドラえもんのひみつ道具検索エンジンを作ってみた](https://shunyaueta.com/posts/2022-10-23-2344/)
1. `393` views: [OSS のアノテーションツール Label Studio を使って、快適にアノテーションする](https://shunyaueta.com/posts/2022-01-09/)
1. `387` views: [社内でデータ分析結果を可視化・共有する際に Google Colab が便利](https://shunyaueta.com/posts/2022-05-10-2200/)

---

## Copy result to clip board

You can pipe the pbcopy

```shell
GOOGLE_APPLICATION_CREDENTIALS="XXX.json" go run main.go -property_id 12345 -pages_root_url https://shunyaueta.com -site_content_path ../hurutoriya.github.io/content  -top_n 15 | pbcopy
```
