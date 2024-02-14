package main

import (
	"fmt"
	"testing"
	"time"

	fasttext "github.com/myagmartseren/fasttext_golang"
)

func TestPredict(t *testing.T) {
	// Initialize FastText model
	ft := fasttext.NewFasttext()

	// Defer closing the model to ensure it's closed after its use
	defer ft.Close()

	// Load the model
	if err := ft.LoadModel("./model.bin"); err != nil {
		fmt.Println("error loading model:", err)
		return
	}

	// Define input text for prediction
	text := `Нийслэлийн онцгой комисс (НОК) хуралдаж мэргэжлийн байгууллагуудын саналыг үндэслэн шар усны үерийн эрсдэлтэй байршлуудад гамшгийн шар түвшин тогтоосон билээ. Сэлбэ голын халиа гүүрний доорх дулааны шугамд тулсан байна.

	Тодруулбал нийслэлд нийт 22 хорооны 33 байршилд халиа тошин үүссэн. Тиймээс холбогдох байгууллагын 284 ажилтан, алба хаагч, 48 техник хэрэгсэлтэй 8,460 метр суваг шуудуу татаж 401 тонн мөсийг гаргаж тээвэрлэн, далан босгох ажлыг хийжээ. 
	
	Үерийн хамгаалалтын барилга байгууламжийн дотор байгаа дулааны шугам нь тухайн хэсэгт голын мөсийг хайлуулж, халиа үүсгэн түвшнийг нэмэгдүүлсэн. 
	
	Иймд дулааны шугамын ойролцоо үүссэн халиа тошингийн мөсийг гар аргаар суваг татан, түвшнийг бууруулах ажлыг хүний оролцоотой хийсэн аж. Энэ ажилд хүн хүчний оролцоо дутмаг байгааг мэргэжлийн байгууллагынхан онцоллоо. 
	
	Сэлбэ гол дагуух халиа, тошин, голын ёроолын хэвгийг гаргах, лаг хагшаас зөөж зайлуулах, хог хаягдлыг цэвэрлэх тээвэрлэх ажлыг богино хугацаанд гүйцэтгэхэд машин механизм, ачаа тээвэрлэлтийн үйлчилгээ үзүүлэх боломжтой ахуйн нэгж байгууллагатай хамтран ажиллахаар нээлттэй сонгон шалгаруулалтыг зарлав`

	// Perform prediction and measure time
	start := time.Now()
	result, err := ft.Predict(text)
	if err != nil {
		fmt.Println("error during prediction:", err)
		return
	}
	elapsed := time.Since(start)

	// Output results including milliseconds
	t.Logf("Prediction result: %v\n", result)
	t.Logf("Time taken: %v\n", elapsed)
}
