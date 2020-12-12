package curator

import (
	"fmt"
)

// インタフェース
// 収集・研究・保存・表示などを行うなにか
type Curator interface {
	Collect() error
	Reserch() error
	Save() error
	Show() error
}

// ポリモーフィズム
func Run(curators []Curator) error {
	for _, curator := range curators {
		fmt.Println("----------------------------")

		if apiCurator, ok := curator.(Authenticator); ok {
			if err := apiCurator.Auth(); err != nil {
				return err
			}
		}

		if err := curator.Collect(); err != nil {
			return err
		}

		if err := curator.Reserch(); err != nil {
			return err
		}

		if err := curator.Save(); err != nil {
			return err
		}

		if err := curator.Show(); err != nil {
			return err
		}
	}

	return nil
}

// 親クラス
type BaseCurator struct {
	Target string        // 情報を集める対象
	data   []interface{} // 収集した情報
}

func (bc *BaseCurator) Collect() error {
	fmt.Println(bc.Target)
	fmt.Println("BaseCurator.Collect")
	bc.data = append(bc.data, "BaseCollect")
	return nil
}

func (bc *BaseCurator) Reserch() error {
	fmt.Println("BaseCurator.Reserch")
	return nil
}

func (bc *BaseCurator) Save() error {
	fmt.Println("BaseCurator.Save")
	return nil
}

func (bc *BaseCurator) Show() error {
	fmt.Println("BaseCurator.Show")
	return nil
}
