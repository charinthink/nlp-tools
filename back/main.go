package main

import (
	"fmt"

	"github.com/charinthink/thai-sentence/back/adapter/utils"
	parallelconfig "github.com/charinthink/thai-sentence/back/core/parallel"
)

func main() {
	/* Thai segmentation */
	thaiSegmentation, _ := utils.NewSegmentation("../data/dict/th/th_corpus.txt")
	wordSegmentation := thaiSegmentation.Segmentation("สวัสดีครับนายหัว")
	fmt.Println(wordSegmentation) // output: [สวัสดี ครับ นายหัว]

	/* Thai similarity */
	thaiSimilarity, _ := utils.NewSimilarity("../data/model/cc.th.300.vec", nil)
	resultThaiSim := thaiSimilarity.Similarity("ครับ", 10)
	fmt.Println(resultThaiSim) // output: [{ครับ 1} {ครับ.แต่ 0.7405722625688913} {ครับ.ถ้า 0.7347521155140914} {ครับ.แล้ว 0.7143124852394049} {ครับPM 0.706172312369067} {ครับ.ผม 0.7059755665531694} {หน่อย 0.6969460786663937} {เลย 0.6947457112436395} {ครับ.ด้วย 0.6890742951117245} {ครับpm 0.6885408235764859}]

	/* Thai detect bad word */
	thaiDetectBadWord, _ := utils.NewDetectBadWord(
		"../data/model/cc.th.300.vec",
		"../data/dict/th/th_swear.txt",
		thaiSegmentation,
		parallelconfig.Default())
	resultBadWords := thaiDetectBadWord.DectectBadWord("สวัสดีหน้า...ส้นตีน", 0.5)
	fmt.Println(resultBadWords) // output: [ส้นตีน]
}
