// ioInfo_test.go
package zwds

import (
	"fmt"
	"testing"
)

func check(ui [13][]string) {
	for i := 0; i < 13; i++ {
		println(len(ui[i]))
	}
}

func TestAny(t *testing.T) {
	var hi HeroInfo
	var lni LiuNianInfo
	hi.bDcopyInfo(infoOne)
	lni.bDcopyInfo(infoTwo)
	lni.bDcopyLiuHua(liuhuaOne)

	var uix UIZHUX
	xg := make(XGINDEX)
	uix.ZxPaiPan(&hi, xg)
	printUI(uix)
	uilx := lni.ZlxPaiPan(hi, uix, xg)
	for k, v := range uilx {
		fmt.Println(k)
		prtuilx(v)
	}
}
func printUI(ui UIZHUX) {
	for i := 0; i < 13; i++ {
		fmt.Printf("%0.2d: ", i)
		for _, v := range ui[i] {
			fmt.Printf("%s ", v)
		}
		fmt.Println(len(ui[i]))
	}
	println()
}

func prtuilx(uilx UILIUX) {
	for i := 0; i < 13; i++ {
		fmt.Printf("%0.2d: ", i)
		for _, v := range uilx[i] {
			fmt.Printf("%s ", v)
		}
		fmt.Println()
	}
	println()
}
