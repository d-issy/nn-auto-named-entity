package wiki

import (
	"reflect"
	"testing"
)

func TestGetProtoBuf(t *testing.T) {
	xmlPage := &XMLPage{
		ID:    1,
		Title: "テスト",
		Revisions: []XMLRevision{
			XMLRevision{Text: `# テスト
{{TestTemplate|A=1|B=2|c=3}}
[[リンク|表示]]]
[[Category: テスト]]`},
		},
	}
	p, err := xmlPage.GetProtoBuf()
	if err != nil {
		t.Fatal(err)
	}
	if p.Id != 1 {
		t.Errorf(`title must be 1, but got %d`, p.Id)
	}
	if p.Title != "テスト" {
		t.Errorf(`title must be "テスト", but got %q`, p.Title)
	}
	if p.Links[0] != "リンク" {
		t.Errorf(`title must be "リンク", but got %q`, p.Links[0])
	}
	if p.Categories[0] != "テスト" {
		t.Errorf(`title must be "リンク", but got %q`, p.Categories[0])
	}
	if p.Templates[0] != "TestTemplate" {
		t.Errorf(`title must be "リンク", but got %q`, p.Templates[0])
	}
}

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
