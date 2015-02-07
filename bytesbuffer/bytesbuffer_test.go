package bytesbuffer

import (
	//"io/ioutil"
	//"log"
	//"os"
	"testing"
)

func Test_Hex_to_str(t *testing.T) {
	ret, _ := Hex_to_str(`\x31\x32\x33\x34\x35\x36\x37\x38\x39\x30\x61\x62\x64\xE4\xB8\xAD\xE6\x96\x87`)
	if ret != "1234567890abd中文" {
		t.Error("result not right: ", ret)
	}
}

func Test_Str_to_hex(t *testing.T) {
	ret := Str_to_hex("1234567890abd中文", "\\x")
	if ret != `\x31\x32\x33\x34\x35\x36\x37\x38\x39\x30\x61\x62\x64\xE4\xB8\xAD\xE6\x96\x87` {
		t.Error("result not right: ", ret)
	}
}

func Test_WriteByteString(t *testing.T) {
	buf1 := NewBufferString("")
	buf1.WriteByteString(`\x31\x32\x33\x34`)
	buf1.Buffer.WriteString("abcd中文ef")

	ret = buf1.Buffer.String()
	if ret != `1234abcd中文ef` {
		t.Error("result not right: ", ret)
	}

	ret := buf1.ByteString("\\x")
	if ret != `\x31\x32\x33\x34\x61\x62\x63\x64\xE4\xB8\xAD\xE6\x96\x87\x65\x66` {
		t.Error("result not right: ", ret)
	}
}

//func main() {

//
//log.Println(ret)

//

//fd1, err := os.Open("in.txt")
//if err != nil {
//	log.Fatalln(err)
//}
//defer fd1.Close()
//inbytes, err := ioutil.ReadAll(fd1)

//buf2 := NewBuffer(inbytes)

//fd2, err := os.Create("123.txt")
//if err != nil {
//	log.Fatalln(err)
//}
//defer fd2.Close()
////fd2.WriteString(buf2.ByteString(" "))
//buf2.WriteToByteString(fd2, "")
//return
//}
