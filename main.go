package main

import (
	"fmt"
	"image"
	"os"

	"golang.org/x/image/bmp"
)

func bitMapYukle(YOL string) (*image.Paletted, error) {

	dosya_bmp, err := os.Open(YOL) //Dosya yolu ile "tank.bmp" alındı.
	if err != nil {
		return nil, err
	}
	defer dosya_bmp.Close()

	img, err := bmp.Decode(dosya_bmp) //Image formatı için decode ettik.

	if err != nil {
		return nil, err
	}

	return img.(*image.Paletted), nil //image'i pointer img->palatted türünden döndürdük "Örnek: {135,135,135,255} desenine sahiptir."

}

func yapHistogram(tankImg *image.Paletted, binDegeri uint8) *[]uint8 { //tank image'ini parametre olarak Paletted türünden alır ve geriye int dizisi pointer'ı döndürür.

	var histoArray = make([]uint8, uint8(255/binDegeri))                  //Histogram için 255'lik bir dizi.
	var lengthTankImage = tankImg.Bounds().Max.X * tankImg.Bounds().Max.Y //"Tank.bmp" Fotograf uzunluğu.

	for i := 0; i < lengthTankImage; i++ { //"Tank.bmp" X*Y (Yükseklik*Genişlik) toplam pixel sayısı kadarlık döngü.
		histoArray[uint8(tankImg.Pix[i]/binDegeri)]++ //"Tank.bmp" Pixeller içinde gezerken çekilen değer başka bir dizinin indexi olmalı ve o indexin değeri "1" arttırılmalıdır.
	}
	/*
		for j := 0; j < len(histoArray); j++ { //histogram degerlerini döndürdük.
			//fmt.Println(histoArray[j])
		}
	*/

	return &histoArray //histoArray dizisi pointer'ı döndürüldü.

}

func main() {
	bitMapCall, _ := bitMapYukle("tank.bmp")
	var binningDegeri uint8 = 4 //BIN BIN BIN -> DEGER DEGISTIRILEBILIR <- BIN BIN BIN
	histoCall := yapHistogram(bitMapCall, binningDegeri)

	fmt.Println(histoCall)

}
