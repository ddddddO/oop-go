// https://qiita.com/kazuya-n/items/fbee07ef778e166cb6dd
package api

import (
	"fmt"

	c "github.com/ddddddO/oop-go/curator"
)

// APIの子クラス(Baseの孫クラス)
type SpotifyCurator struct {
	c.ApiCurator // 継承
}

// コンストラクタ
func NewSpotifyCurator(target, clientID, clientSecret string) *SpotifyCurator {
	return &SpotifyCurator{
		c.ApiCurator{
			BaseCurator: c.BaseCurator{
				Target: target,
			},

			ClientID:     clientID,
			ClientSecret: clientSecret,
		},
	}
}

// オーバーライド
func (sc *SpotifyCurator) Reserch() error {
	fmt.Println("SpotifyCurator.Reserch")
	return nil
}
