package main

import (
	// "fmt"
	// "io"
	"bufio"
	"fmt"
	"os"
)

func main() {
	// f, err := os.Open("23_FileHAndling/ex.txt")
	// if err!=nil{
	// 	panic(err)
	// // }
	// fileInfo ,err:=f.Stat()
	// if err!=nil{
	// 	panic(err)
	// }
	// // fmt.Println("fileName",fileInfo.Name())
	// fmt.Println("fileName",fileInfo.Mode())
	// fmt.Println("fileName",fileInfo.Name())
	// fmt.Println("fileName",fileInfo.ModTime())

	// f, err := os.Open("23_FileHAndling/ex.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()

	// buf := make([]byte, 10)

	// d, err := f.Read(buf)
	// if err != nil && err != io.EOF {
	// 	panic(err)
	// }

	// fmt.Println("bytes read:", d)
	// fmt.Println("data:", string(buf))
	// // for i:=0 ; i<len(buf) ; i++{
	// // 	fmt.Println("data:", string(buf[i]))

	// // }

	// read folders
	// dir , er:=os.Open(".")
	// if er!=nil{
	// 	panic(er)
	// }
	// defer dir.Close()
	// filInfos , er :=dir.ReadDir(-1)
	// for _,fi :=range filInfos{
	// 	fmt.Println(fi.Name())
	// }
	// f , err:=os.Create("ex2.txt")
	// if err!=nil{
	// 	panic(err)
	// }
	// defer f.Close()
	// // f.WriteString("hi go")
	// bytes:= []byte("jhel golang")
	// f.Write(bytes)



	// read and write to another file

	sourecFile , err := os.Open("23_FileHAndling/ex.txt")
	if err!=nil{
		panic(err)
	}
	defer sourecFile.Close()
	destFile ,err := os.Create("ex4.txt")

    if err!=nil{
		panic(err)
	}
	defer destFile.Close()
	reader := bufio.NewReader(sourecFile)
	writer := bufio.NewWriter(destFile)

	for{
	 b ,err	:=reader.ReadByte()
	 if err!=nil{
		if err.Error()!="EOF"{
		panic(err)
	}
	break
}

	er :=writer.WriteByte(b)
	if er!=nil{
		panic(err)
	}


	}
	writer.Flush()
	fmt.Println("wriiten succesfully")


}