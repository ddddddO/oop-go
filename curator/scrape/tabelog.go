package scrape

import (
	"fmt"

	c "github.com/ddddddO/oop-go/curator"
)

// Scrapeの子クラス(Baseの孫クラス)
type TabelogCurator struct {
	c.ScrapeCurator // 継承
}

// コンストラクタ
func NewTabelogCurator(target string) *TabelogCurator {
	return &TabelogCurator{
		c.ScrapeCurator{
			c.BaseCurator{
				Target: target,
			},
		},
	}
}

// オーバーライド
func (tc *TabelogCurator) Reserch() error {
	fmt.Println("TabelogCurator.Reserch")
	return nil
}
