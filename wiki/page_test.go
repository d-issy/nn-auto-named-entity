package wiki

import (
	"reflect"
	"testing"
)

func TestParseLinks(t *testing.T) {
	input := `[[ウィキリンク]]
[[パイプ付きリンク|表示名]]
[[セクションリンク#セクション名]]
[[マルチリンク#セクション名|表示名]]
[[重複チェック]]
[[重複チェック]]
[[Category:無視]]
[[:ファイル:無視]]`
	expect := []string{"ウィキリンク", "パイプ付きリンク", "セクションリンク",
		"マルチリンク", "重複チェック"}
	result := parseLinks(input)
	if !reflect.DeepEqual(expect, result) {
		t.Fatalf("Expected %q, but got %q", expect, result)
	}
}

func TestParseCategories(t *testing.T) {
	input := `[[Category: 英字カテゴリ]]
[[カテゴリ: 日本語カテゴリ]]
[[Category: ソートキーカテゴリ|ソートキ]]`
	expect := []string{"英字カテゴリ", "日本語カテゴリ", "ソートキーカテゴリ"}
	result := parseCategories(input)
	if !reflect.DeepEqual(expect, result) {
		t.Fatalf("Expected %q, but got %q", expect, result)
	}
}

func TestParseTemplates(t *testing.T) {
	input := `{{標準テンプレート|引数1=1|引数2=2}}
{{長いテンプレート
|引数1 = {{1}}
|引数2 = {{2{{3{{4}}}}}}
}}
{{#無視}}
{{重複}}
{{重複}}`
	expect := []string{"標準テンプレート", "長いテンプレート", "重複"}
	result := parseTemplates(input)
	if !reflect.DeepEqual(expect, result) {
		t.Fatalf("Expected %q, but got %q", expect, result)
	}
}
