//376938_ABDULLAH_BIYIK
//KTÜ-BİLGİSAYAR-MÜHENDİSLİĞİ-4.SINIF-1.ÖĞRETİM
//ÖDEV-1-SAYISAL-İŞARET-İŞLEME-(HİSTOGRAM PMF ORTALAMA STANDART SAPMA)
package main

import (
	"fmt"
	"image"
	"math"
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

func imageLengthXy(tankImg *image.Paletted) uint32 { //tank fotoğrafının boyutu(X*Y) döndürüldü.
	var lengthTankImage uint32 = uint32(tankImg.Bounds().Max.X * tankImg.Bounds().Max.Y)
	return lengthTankImage
}

func yapHistogram(tankImg *image.Paletted, binDegeri uint8) *[]uint16 { //tank image'ini parametre olarak Paletted türünden alır ve geriye int dizisi pointer'ı döndürür.

	var histoArray = make([]uint16, 255)         //Histogram için 255'lik bir dizi.
	var lengthTankImage = imageLengthXy(tankImg) //"Tank.bmp" Fotograf uzunluğu.

	for i := 0; i < int(lengthTankImage); i++ { //"Tank.bmp" X*Y (Yükseklik*Genişlik) toplam pixel sayısı kadarlık döngü.

		histoArray[uint8(tankImg.Pix[i])]++ //"Tank.bmp" Pixeller içinde gezerken çekilen değer başka bir dizinin indexi olmalı ve o indexin değeri "1" arttırılmalıdır.

	}

	if binDegeri != 1 {
		for j := 0; j < 255; j += int(binDegeri) {
			for k := 1; k < int(binDegeri); k++ {
				if j+k < 255 {
					histoArray[j] += histoArray[j+k]
					histoArray[j+k] = 0
				}
			}
		}
	}

	return &histoArray //histoArray dizisi pointer'ı döndürüldü.

}

func yapPMF(histoPtr *[]uint16, binDegeri uint8, tankImg *image.Paletted) *[]float32 {

	var pmf_array = make([]float32, 255)
	var lengthTankImage = imageLengthXy(tankImg)
	for i := 0; i < int(len(*histoPtr)); i += int(binDegeri) { // PMF alt alan 1 olacak şekilde toplam piksel sayısına oranla indexleniyor...
		pmf_array[i] = float32((*histoPtr)[i]) / float32((lengthTankImage))
		/*if pmf_array[i] > 1 {
			fmt.Println("\n histo->", (*histoPtr)[i], "\n length->", lengthTankImage, "\n pmfarray->", pmf_array[i])// 1den küçük değermi kontrolü için yazılımştı...
		}*/
	}

	//(*histoPtr)[i]

	return &pmf_array

}

func hesaplaOrtalama(histoPtr *[]uint16, binDegeri uint8) float32 {

	var ortalama float32 = 0
	var havuz int = 0
	var adet uint16 = 0
	for i := 0; i < int(len(*histoPtr)); i += int(binDegeri) { //ortalama formülü gerçekleştirildi.
		havuz += int((*histoPtr)[i])
		adet++
	}
	ortalama = float32(havuz) / float32(adet)

	return ortalama
}

func hesaplaStandartSapma(histoPtr *[]uint16, binDegeri uint8, histoOrtalamaDegeri float32) float32 {

	var s_sapma float32 = 0
	var toplam_havuz float32 = 0
	var fark float32 = 0
	var adet uint16
	for i := 0; i < int(len(*histoPtr)); i += int(binDegeri) { //standart sapma için formül gerçekleştiriliyor.
		fark = float32((*histoPtr)[i]) - histoOrtalamaDegeri
		fark = fark * fark
		toplam_havuz = toplam_havuz + fark
		fark = 0
		adet++
	}
	//s_sapma = toplam_havuz / float32((len(*histoPtr) - 1))
	s_sapma = toplam_havuz / float32(adet)

	s_sapma = float32(math.Sqrt(float64(s_sapma)))

	return s_sapma
}

func histoYazdir(histoPtr *[]uint16, binDegeri uint8) {
	fmt.Print("\n [--Histogram--] \n")
	for i := 0; i < 255; i += int(binDegeri) { //bin degerine gore indexi geziyoruz.
		if binDegeri != 1 {
			var k uint8 = 0
			for k = 1; k < binDegeri; k++ {
				fmt.Print(".") //bin degerine göre gösterimde boş indexleri "nokta" ile temsil ettik.
			}
		} else {
			fmt.Print(",") //bin değeri 1 ise "virgül" temsil kullandık, çünkü boş index yok.
		}
		fmt.Print((*histoPtr)[i])

	}
	fmt.Print("\n [--Histogram--]")
}

func main() {

	bitMapCall, _ := bitMapYukle("tank.bmp")
	var binningDegeri uint8
	fmt.Println("->BIN DEGERI GIRISI:(1,2,3,4) 0->EXIT")
	fmt.Scanln(&binningDegeri)

	for {
		if binningDegeri != 0 {

			histoCall := yapHistogram(bitMapCall, binningDegeri)                                            //histoCall histogramArray ptr tutuyor.
			histoOrtalamaDegeri := hesaplaOrtalama(histoCall, binningDegeri)                                //histocall ptr ve binnig değeri parametre alan ortalama fonk. çağırılıyor.
			histoStandartSapmaDegeri := hesaplaStandartSapma(histoCall, binningDegeri, histoOrtalamaDegeri) //histocall ptr ve binnig parametre alan StndrtSapma fonksiyonu çağırılıyor.

			histoYazdir(histoCall, binningDegeri) //histogram indexleri terminale basılıyor.

			fmt.Print("\n\n[___OrtalamaDegeri__]->[", histoOrtalamaDegeri, "]\n[___StandartSapma___]->[", histoStandartSapmaDegeri, "]\n") //ortalama ve ss terminale basılyıor.

			ProMassFunc := yapPMF(histoCall, binningDegeri, bitMapCall) // PMF kısmı
			fmt.Print("\n [---PMF---] \n", "[", ProMassFunc, "]")       // PMF doğrudan ekrana basıldı.

			fmt.Println("\n\nBIN DEGERI GIRISI:(1,2,3,4) 0->EXIT")
			fmt.Scanln(&binningDegeri)
		} else {
			break
		}

	}

}
