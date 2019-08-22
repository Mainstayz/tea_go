package main

import (
	"../encrypt"
	"log"
)

func main() {
	log.Println("加密前数据")
	data := []byte{0x15, 0x11, 0x11, 0x11, 0x12, 0x11, 0x11, 0x11}
	log.Printf("%x",data)

	outData := []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	key := []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11}

	log.Println("TeaEncryptECB ------ 加密")
	encrypt.TeaEncryptECB(data, key, outData)
	log.Printf("%x", outData)

	outData1 := []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	log.Println("TeaEncryptECB ------ 解密")
	encrypt.TeaDecryptECB(outData, key, outData1)
	log.Printf("%x", outData1)


	outData2 := []byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		}
	var outLen int
	log.Println("OiSymmetryEncrypt2 ------ 加密")
	encrypt.OiSymmetryEncrypt2(data, len(data), key, outData2, &outLen)
	log.Printf("%x", outData2[:outLen])
	outData3 := []byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}
	//var outLen1 int
	outLen1 := 2 * outLen
	//outData4 := outData2[:outLen]
	encrypt.OiSymmetryDecrypt2(outData2, outLen, key, outData3, &outLen1)
	log.Println("OiSymmetryDecrypt2 ------ 解密")
	log.Printf("%x", outData3[:outLen1])
}