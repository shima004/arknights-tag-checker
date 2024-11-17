# Arknights Tag Checker

This project is a tool for filtering Arknights operators by tags.

## Features

- Select tags to search for operators
- Combine multiple tags for filtering
- Search for operators obtainable through public recruitment

## Installation

Follow these steps to install:

```sh
go get github.com/shimapaca/arknights-tag-checker
```

## Example

### Search for operators by tags

```go
package main

import (
  "arknights-tag-checker/pkg/filter"
  "arknights-tag-checker/pkg/model"
  "arknights-tag-checker/pkg/search"
  "fmt"
)

func main() {
  // Load operator data
  search, err := search.NewSearchOperator("data/operators.json")
  if err != nil {
    panic(err)
  }
  // Filter by tags and rarity
  tags := model.Tags{
    model.TagDPS, model.TagDefense,
  }
  rarities := []int{6, 5}
  // Search for operators
  operators := search.SearchOperator(
    &filter.TagFilter{Tags: tags, IsAnd: true},
    &filter.RarityFilter{Rarities: rarities},
  )
  // Display results
  fmt.Printf("%v : %v\n", tags, operators)
}

>> >> [火力 防御] : [ソーンズ : ★ 6 ブレミシャイン : ★ 6 ホシグマ : ★ 6 ホルン : ★ 6 マドロック : ★ 6 ユーネクテス : ★ 6 滌火ジェシカ : ★ 6 アステシア : ★ 5 Blitz : ★ 5 アスベストス : ★ 5 アッシュロック : ★ 5 ヴァルカン : ★ 5 オーロラ : ★ 5 シャレム : ★ 5 ツェルニー : ★ 5 ファイヤーホイッスル : ★ 5 リスカム : ★ 5]

```

### Search for operators obtainable through public recruitment

```go
package main

import (
  "arknights-tag-checker/pkg/filter"
  "arknights-tag-checker/pkg/model"
  "arknights-tag-checker/pkg/search"
  "fmt"
)

func main() {
  // Load operator data
  search, err := search.NewSearchOperator("data/operators.json")
  if err != nil {
    panic(err)
  }
  // Filter by tags, rarity, and acquisition method
  tags := model.Tags{
    model.TagDPS, model.TagDefense, model.TagHealing, model.TagVanguard, model.TagSlow,
  }
  rarities := []int{5, 4, 3}
  acquisitionMethods := []model.AcquisitionMethod{model.AMPublicRecruitment}
  // Search for all combinations
  mutationTags := tags.CombinationsTags()
  for _, tags := range mutationTags {
    operators := search.SearchOperator(
      &filter.TagFilter{Tags: tags, IsAnd: true},
      &filter.RarityFilter{Rarities: rarities},
      &filter.AcquisitionMethodFilter{AcquisitionMethods: acquisitionMethods, IsAnd: true},
    )

    if len(operators) != 0 {
      fmt.Printf("%v : %v\n", tags, operators)
    }
  }
}

>> [火力] : [キアーベ : ★ 5 リード : ★ 5 アステシア : ★ 5 インドラ : ★ 5 エアースカーペ : ★ 5 スワイヤー : ★ 5 フリント : ★ 5 アスベストス : ★ 5 ヴァルカン : ★ 5 リスカム : ★ 5 アズリウス : ★ 5 アンドレアナ : ★ 5 エイプリル : ★ 5 グレースロート : ★ 5 ファイヤーウォッチ : ★ 5 プラチナ : ★ 5 プロヴァンス : ★ 5 ナイトメア : ★ 5 レイズ : ★ 5 イースチナ : ★ 5 クリフハート : ★ 5 マンティコア : ★ 5 ヴィグナ : ★ 4 スカベンジャー : ★ 4 アレーン : ★ 4 ウタゲ : ★ 4 カッター : ★ 4 ドーベルマン : ★ 4 ビーハンター : ★ 4 フロストリーフ : ★ 4 マトイマル : ★ 4 ムース : ★ 4 アシッドドロップ : ★ 4 アンブリエル : ★ 4 ヴァーミル : ★ 4 ジェシカ : ★ 4 メイ : ★ 4 メテオ : ★ 4 カシャ : ★ 4 ヘイズ : ★ 4 ジェイ : ★ 4 プリュム : ★ 3 ミッドナイト : ★ 3 メランサ : ★ 3 アドナキエル : ★ 3 クルース : ★ 3 スチュワード : ★ 3]
[火力 防御] : [アステシア : ★ 5 アスベストス : ★ 5 ヴァルカン : ★ 5 リスカム : ★ 5]
[火力 治療] : [ナイトメア : ★ 5]
[火力 治療 減速] : [ナイトメア : ★ 5]
[火力 先鋒タイプ] : [キアーベ : ★ 5 リード : ★ 5 ヴィグナ : ★ 4 スカベンジャー : ★ 4 プリュム : ★ 3]
[火力 減速] : [アンドレアナ : ★ 5 ナイトメア : ★ 5 イースチナ : ★ 5 フロストリーフ : ★ 4 アンブリエル : ★ 4 メイ : ★ 4]
[防御] : [アステシア : ★ 5 アスベストス : ★ 5 ヴァルカン : ★ 5 ウン : ★ 5 クロワッサン : ★ 5 ニアール : ★ 5 リスカム : ★ 5 ビーズワクス : ★ 5 クオーラ : ★ 4 グム : ★ 4 バブル : ★ 4 マッターホルン : ★ 4 グラベル : ★ 4 カーディ : ★ 3 スポット : ★ 3 ビーグル : ★ 3]
[防御 治療] : [ウン : ★ 5 ニアール : ★ 5 グム : ★ 4 スポット : ★ 3]
[治療] : [ウン : ★ 5 ニアール : ★ 5 ナイトメア : ★ 5 サイレンス : ★ 5 フィリオプシス : ★ 5 ワルファリン : ★ 5 テンニンカ : ★ 4 グム : ★ 4 ススーロ : ★ 4 セイリュウ : ★ 4 パフューマー : ★ 4 ミルラ : ★ 4 ポデンコ : ★ 4 スポット : ★ 3 アンセル : ★ 3 ハイビスカス : ★ 3]
[治療 先鋒タイプ] : [テンニンカ : ★ 4]
[治療 減速] : [ナイトメア : ★ 5 ポデンコ : ★ 4]
[先鋒タイプ] : [エリジウム : ★ 5 キアーベ : ★ 5 ズィマー : ★ 5 テキサス : ★ 5 リード : ★ 5 ヴィグナ : ★ 4 スカベンジャー : ★ 4 テンニンカ : ★ 4 バニラ : ★ 3 フェン : ★ 3 プリュム : ★ 3]
[減速] : [アンドレアナ : ★ 5 ナイトメア : ★ 5 イースチナ : ★ 5 グラウコス : ★ 5 エフイーター : ★ 5 フロストリーフ : ★ 4 アンブリエル : ★ 4 シラユキ : ★ 4 メイ : ★ 4 グレイ : ★ 4 アーススピリット : ★ 4 ポデンコ : ★ 4 オーキッド : ★ 3]

```

## Contribution

Contributions are welcome. Please report issues before sending a pull request.

## License

This project is licensed under the MIT License.
