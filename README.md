# thai-sentence

### **** Before running the ThaiSimilarity function you need to download model from the following link and place it in the `data/model` directory. ***
https://dl.fbaipublicfiles.com/fasttext/vectors-crawl/cc.th.300.vec.gz
>
**MacOS/Linux**
>>
`curl -o model.bin.gz https://dl.fbaipublicfiles.com/fasttext/vectors-crawl/cc.th.300.vec.gz`
>>
`gunzip cc.th.300.vec.gz`

### Models ###
* [cc.th.300.vec](https://dl.fbaipublicfiles.com/fasttext/vectors-crawl/cc.th.300.vec.gz) (Word embeddings)
### Options for training the model ###
* Text classification
### Algorithm ###
* Maximum Matching (Word segmentation)
* Cosine similarity formula (Searching for word similarity)
### Word List ###
* [PyThaiNLP/lexicon-thai](https://github.com/PyThaiNLP/lexicon-thai.git)
### Parallel Config
The default value for 'ParallelConfig' is optimized for CPU: 8 cores and RAM: 16 GB.

`Cpu: 4, ErrorBufferSize: 1000, EmbedBufferSize: 100000, ScanBufferSize: 256 * 1024, WorkerPoolSize: 12`

You can use functions in an interface to set default values or manually set them.
```
type Config interface {
	IsEnabled() bool
	Get() config
}

func Default() Config {
	return &config{
		true,
		routineConfig{
			cpu,
			errorBufferSize,
			embedBufferSize,
			scanBufferSize,
			workerPoolSize,
		},
	}
}

func Set(core, errorBufferSize, embedBufferSize, scanBufferSize, workerPoolSize int) Config {
	return &config{
		true,
		routineConfig{
	      core,
			 errorBufferSize,
			 embedBufferSize,
			 scanBufferSize,
			  workerPoolSize,
		},
	}
}
```
### Example 1 ###
```go
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
```
### Example 2 train model text classification ###
```python
from train import Train
from load_data import LoadData


def main():
    df = LoadData.csv("../data/csv/review_shopping.csv", ["text", "label"])
    Train.text_classification(df, "../data/test.pkl").text_classfication_test()


if __name__ == "__main__":
    main()

// output
Classification Report:
               precision    recall  f1-score   support

         neg       0.76      1.00      0.87        13
         pos       1.00      0.69      0.82        13

    accuracy                           0.85        26
   macro avg       0.88      0.85      0.84        26
weighted avg       0.88      0.85      0.84        26
```