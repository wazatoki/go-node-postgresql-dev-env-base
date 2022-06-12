# go-node-postgresql-dev-env-base

## tmp_project_nameを修正する。

```
find . -type f -print0 | xargs -0 sed -i -e "s/tmp_project_name/test_sample/" 
```

## go プロジェクトの初期化

```
mkdir project_dir
cd project_dir
go mod tidy

```

## frontend プロジェクトの作成


