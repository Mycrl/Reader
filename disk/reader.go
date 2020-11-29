package disk

import (
    "os"
    "bytes"
    "io/ioutil"
    "golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func Reader(p string, o int64, l int64) ([]byte, error) {
    file, open_e := os.Open(p)
    if open_e != nil {
        return nil, open_e
    }
defer file.Close()
    buffer := make([]byte, l)
    s, read_e := file.ReadAt(buffer, o)
    if read_e != nil {
        return nil, read_e
    }
    packet, to_e := GbkToUtf8(buffer[:s])
    if to_e != nil {
        return nil, to_e
    }
    return packet, nil
}

func GbkToUtf8(s []byte) ([]byte, error) {
	if !is_gbk(s) {
		return s, nil
	}
    b_read := bytes.NewReader(s)
    decoder := simplifiedchinese.GBK.NewDecoder()
	reader := transform.NewReader(b_read, decoder)
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func is_gbk(s []byte) bool {
	var i int = 0
	length := len(s)
for i < length {
    if s[i] <= 0xff {
        i++
        continue
    }
    if s[i] < 0x81 {
        return false
    }
    if s[i] > 0xfe {
        return false
    }
    if s[i+1] < 0x40 {
        return false
    }
    if s[i+1] > 0xfe {
        return false
    }
    if s[i+1] == 0xf7 {
        return false
    }
    i += 2
    continue
}
	return true
}