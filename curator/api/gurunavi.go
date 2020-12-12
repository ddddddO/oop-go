// https://api.gnavi.co.jp/api/
package api

import (
	"fmt"

	c "github.com/ddddddO/oop-go/curator"
)

// APIの子クラス(Baseの孫クラス)
type GurunaviCurator struct {
	c.ApiCurator // 継承
}

// コンストラクタ
func NewGurunaviCurator(target, clientID, clientSecret string) *GurunaviCurator {
	return &GurunaviCurator{
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
func (gc *GurunaviCurator) Reserch() error {
	fmt.Println("GurunaviCurator.Reserch")
	return nil
}
