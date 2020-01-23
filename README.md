# go-resas-api

[![godoc](https://godoc.org/github.com/yyoshiki41/go-resas-api?status.svg)](https://godoc.org/github.com/yyoshiki41/go-resas-api)

[RESAS](https://resas.go.jp/#/13/13101) API client for Go

## RESAS (地域経済分析システム) API について

https://opendata.resas-portal.go.jp/

## Usage

```go
cli, err := resas.NewClient(os.Getenv("RESAS_API_KEY"))
if err != nil {
	log.Fatal(err)
}
// https://opendata.resas-portal.go.jp/docs/api/v1/cities.html
citie, err := cli.GetCitiesByPrefCode(context.Background(), 37)
if err != nil {
	log.Fatal(err)
}
fmt.Printf("%#v\n", cities)
```

## API Key

RESAS API の利用には API が必要です。

https://opendata.resas-portal.go.jp/docs/api/v1/detail/index.html

## API

✅ 都道府県一覧 api/v1/prefectures

✅ 市区町村一覧 api/v1/cities

- [ ] 旧市区町村一覧 api/v1/oldCities
- [ ] 産業大分類 api/v1/industries/broad
- [ ] 産業中分類 api/v1/industries/middle
- [ ] 産業小分類 api/v1/industries/narrow
- [ ] 職業大分類 api/v1/jobs/broad
- [ ] 職業中分類 api/v1/jobs/middle
- [ ] 特許.技術分野 api/v1/patents/broad
- [ ] 特許.技術テーマ api/v1/patents/middle
- [ ] 税関 api/v1/customs
- [ ] 輸出入.取引国\_地域 api/v1/regions/broad
- [ ] 輸出入.取引国\_国 api/v1/regions/middle
- [ ] 農業部門 api/v1/regions/agricultureDepartments
- [ ] 特許権者の所在地 api/v1/patents/locations
- [ ] 輸出入.品目\_大分類 api/v1/tradeInfoItemTypes/broad
- [ ] 輸出入.品目\_中分類 api/v1/tradeInfoItemTypes/middle
- [ ] 輸出入.品目\_小分類 api/v1/tradeInfoItemTypes/narrow

### API 一覧

https://opendata.resas-portal.go.jp/docs/api/v1/index.html
