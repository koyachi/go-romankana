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
