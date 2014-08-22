package romankana

import (
	"testing"
)

type TestData struct {
	input    string
	expected string
}

func TestKanaRoman(t *testing.T) {
	testData := []TestData{
		TestData{
			"",
			"",
		},
		TestData{
			"ano iihatoovo no sukitootta kaze",
			"ano iihatoovo no sukitootta kaze",
		},
		TestData{
			"あのいーはとーゔぉのすきとおったかぜ",
			"anoiihatoovonosukitoottakaze",
		},
		TestData{
			"アノイーハトーヴォノスキトオッタカゼ",
			"anoiihatoovonosukitoottakaze",
		},
		TestData{
			"コウエンジ ヒモンヤ",
			"kouenji himonya",
		},
		TestData{
			"マークザッカーバーグ アーノルドシュワルツネッガー",
			"maakuzakkaabaagu aanorudoshuwarutsuneggaa",
		},
	}
	for _, td := range testData {
		actual := KanaRoman(td.input)
		if actual != td.expected {
			t.Errorf("got %v\nwant %v", actual, td.expected)
		}
	}
}

func TestFindKanaFromStr(t *testing.T) {
	testData := []TestData{
		TestData{
			"",
			"",
		},
		TestData{
			"あ",
			"あ",
		},
		TestData{
			"nnya",
			"ンヤ",
		},
		TestData{
			"nga",
			"ンガ",
		},
		TestData{
			"nya",
			"ニャ",
		},
		TestData{
			"mba",
			"ンバ",
		},
		TestData{
			"ma",
			"マ",
		},
		TestData{
			"mdo",
			"mド",
		},
		TestData{
			"tto",
			"ット",
		},
		TestData{
			"aada",
			"アアダ",
		},
		TestData{
			"aaaaaaaa",
			"アアアアアアアア",
		},
	}
	for _, td := range testData {
		actual := FindKanaFromStr(td.input)
		if actual != td.expected {
			t.Errorf("got %v\nwant %v", actual, td.expected)
		}
	}
}

func _TestRomanKana(t *testing.T) {
	testData := []TestData{
		TestData{
			"あのいーはとーゔぉのすきとおったかぜ",
			"あのいーはとーゔぉのすきとおったかぜ",
		},
		TestData{
			"アノイーハトーヴォノスキトオッタカゼ",
			"アノイーハトーヴォノスキトオッタカゼ",
		},
		TestData{
			"ano iihatoovo no sukitootta kaze",
			"アノ イイハトオヴォ ノ スキトオッタ カゼ",
		},
		TestData{
			"kouennji himonnya",
			"コウエンジ ヒモンヤ",
		},
		TestData{
			"kouenji himonya",
			"コウエンジ ヒモニャ",
		},
		TestData{
			"nishinippori hambaagu",
			"ニシニッポリ ハンバアグ",
		},
		TestData{
			"nisinippori hanbaagu",
			"ニシニッポリ ハンバアグ",
		},
	}
	for _, td := range testData {
		actual := RomanKana(td.input)
		if actual != td.expected {
			t.Errorf("got %v\nwant %v", actual, td.expected)
		}
	}
}
