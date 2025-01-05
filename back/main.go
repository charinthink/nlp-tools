package main

import (
	"fmt"

	"github.com/charinthink/thai-sentence/adapter/utilities"
)

func main() {
	result, _ := utilities.ThaiSegmentation("../data/dict/th/th_corpus.txt", "สำหรับการตัดคำภาษาไทยเป็นสิ่งสำคัญในการพัฒนาโปรแกรมที่เกี่ยวข้องกับข้อความภาษาไทย ไม่ว่าจะเป็นการแปลภาษา หรือระบบแชทบอท การลงทุนเวลาเขียนชุดการทดสอบที่ดีจะช่วยให้ระบบทำงานได้อย่างแม่นยำและลดเวลาในการแก้ไขข้อผิดพลาดในอนาคต")
	fmt.Println(result)
}