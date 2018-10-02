package wiki

import (
	"strings"
	"testing"
)

const wikiExample = `<mediawiki xmlns="http://www.mediawiki.org/xml/export-0.8/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.mediawiki.org/xml/export-0.8/ http://www.mediawiki.org/xml/export-0.8.xsd" version="0.8" xml:lang="en">
  <siteinfo>
    <sitename>Wikipedia</sitename>
    <base>http://en.wikipedia.org/wiki/Main_Page</base>
    <generator>MediaWiki 1.24wmf2</generator>
    <case>first-letter</case>
    <namespaces>
      <namespace key="-2" case="first-letter">メディア</namespace>
      <namespace key="-1" case="first-letter">特別</namespace>
      <namespace key="0" case="first-letter" />
      <namespace key="1" case="first-letter">ノート</namespace>
      <namespace key="2" case="first-letter">利用者</namespace>
      <namespace key="3" case="first-letter">利用者‐会話</namespace>
      <namespace key="4" case="first-letter">Wikipedia</namespace>
      <namespace key="5" case="first-letter">Wikipedia‐ノート</namespace>
      <namespace key="6" case="first-letter">ファイル</namespace>
      <namespace key="7" case="first-letter">ファイル‐ノート</namespace>
      <namespace key="8" case="first-letter">MediaWiki</namespace>
      <namespace key="9" case="first-letter">MediaWiki‐ノート</namespace>
      <namespace key="10" case="first-letter">Template</namespace>
      <namespace key="11" case="first-letter">Template‐ノート</namespace>
      <namespace key="12" case="first-letter">Help</namespace>
      <namespace key="13" case="first-letter">Help‐ノート</namespace>
      <namespace key="14" case="first-letter">Category</namespace>
      <namespace key="15" case="first-letter">Category‐ノート</namespace>
      <namespace key="100" case="first-letter">Portal</namespace>
      <namespace key="101" case="first-letter">Portal‐ノート</namespace>
      <namespace key="102" case="first-letter">プロジェクト</namespace>
      <namespace key="103" case="first-letter">プロジェクト‐ノート</namespace>
      <namespace key="828" case="first-letter">モジュール</namespace>
      <namespace key="829" case="first-letter">モジュール‐ノート</namespace>
      <namespace key="2300" case="first-letter">Gadget</namespace>
      <namespace key="2301" case="first-letter">Gadget talk</namespace>
      <namespace key="2302" case="case-sensitive">Gadget definition</namespace>
      <namespace key="2303" case="case-sensitive">Gadget definition talk</namespace>
    </namespaces>
  </siteinfo>
  <page>
    <title>テスト</title>
    <ns>0</ns>
    <id>5</id>
    <revision>
      <id>69584826</id>
      <timestamp>2018-09-01T02:03:04Z</timestamp>
      <text xml:space="preserve">これはテストテキストです</text>
    </revision>
  </page>
  <page>
    <title>テスト</title>
    <ns>0</ns>
    <id>5</id>
    <revision>
      <id>69584826</id>
      <timestamp>2018-09-01T02:03:04Z</timestamp>
      <text xml:space="preserve">これはテストテキストです</text>
    </revision>
  </page>
</mediawiki>`

func TestParseCheck(t *testing.T) {
	parser, err := NewXMLParser(strings.NewReader(wikiExample))
	if err != nil {
		t.Fatal(err)
	}
	pageCount := 0

	for err == nil {
		page, err := parser.Next()
		if err != nil {
			break
		}
		pageCount += 1
		if title := page.Title; title != "テスト" {
			t.Fatalf(`page title must be "テスト", but got %q`, title)
		}
	}
	if pageCount != 2 {
		t.Fatalf("pageCount must be 2, but got %d", pageCount)
	}
}
