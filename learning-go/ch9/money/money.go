// パッケージmoneyは金銭を管理するための様々なユーティリティを提供する
package money

// Moneyは「金額」と「通貨」の組みを表す
type Money struct {
	Value     int
	Currencty string
}

// Convertはひとつの通貨を別の通貨に変換する
func Convert(from Money, to string) (Money, error) {
	return Money{}, nil
}