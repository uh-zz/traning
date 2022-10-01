package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const MONSTER_BATTLE_API = "https://ob6la3c120.execute-api.ap-northeast-1.amazonaws.com/Prod/battle/"

type Match struct {
	Winner string `json:"winner"`
	Loser  string `json:"loser"`
}

func main() {
	flag.Parse()

	monsters := flag.Args()

	monsterPair := createPairs(monsters)

	allMatches := fetchAllMatches(monsterPair)

	ranking := makeRanking(allMatches)

	fmt.Println(strings.Join(ranking, ","))
}

// モンスターの対戦組み合わせのリスト
func createPairs(monsters []string) []string {
	var pairs []string
	// 組み合わせ 5C2
	for i := 0; i < len(monsters); i++ {
		for j := i + 1; j < len(monsters); j++ {
			pairs = append(pairs, monsters[i]+"+"+monsters[j])
		}
	}
	return pairs
}

// すべての対戦結果を取得
func fetchAllMatches(pairs []string) []Match {

	var matches []Match

	for _, pair := range pairs {
		resp, err := http.Get(MONSTER_BATTLE_API + pair)
		if err != nil {
			fmt.Printf("MONSTER_BATTLE_API failed: %s", err.Error())
		}
		defer resp.Body.Close()

		var match Match
		bytesArray, _ := io.ReadAll(resp.Body)
		json.Unmarshal(bytesArray, &match)

		matches = append(matches, match)
	}
	return matches
}

// 対戦結果からランキングを作成
func makeRanking(matches []Match) []string {
	var ranking []string
	for _, match := range matches {
		ranking = sort(match, ranking)
	}
	return ranking
}

func sort(match Match, ranking []string) []string {
	if len(ranking) == 0 {
		ranking = append(ranking, match.Winner)
		ranking = append(ranking, match.Loser)
		return ranking
	}

	for i := 0; i < len(ranking); i++ {
		if ranking[i] == match.Loser {
			ranking = append(ranking[:i+1], ranking[i:]...)
			ranking[i] = match.Winner

			ranking = unique(ranking)
			break
		}
	}
	return ranking
}

func unique(ranking []string) []string {
	var list []string
	for _, monster := range ranking {
		if !contains(list, monster) {
			list = append(list, monster)
		}
	}
	return list
}

func contains(list []string, monster string) bool {
	for _, m := range list {
		if m == monster {
			return true
		}
	}
	return false
}
