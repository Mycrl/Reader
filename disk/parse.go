package disk

import (
	"Reader/storage"
	"bufio"
	"io"
	"os"
	"regexp"
	"strings"
)

const REGEXP_TEMPLATE = `(^[\x{7B2C}]{0,1})
    ([0-9]{1,5}|[\x{4E00}\x{4E8C}\x{4E09}\x{56DB}\x{4E94}\x{516D}\x{4E03}\x{516B}\x{4E5D}\x{5341}\x{767E}\x{5343}]{1,10})
    ([\x{7AE0}]{0,1}|[\x{8282}]{0,1})\s`

func Parse(p string) (*storage.Book, error) {
	file, open_e := os.Open(p)
	if open_e != nil {
		return nil, open_e
	}
defer file.Close()
	var Chapters []storage.Chapter
	var offset int32 = 0
	empty := regexp.MustCompile("\r|\n")
	must := regexp.MustCompile(REGEXP_TEMPLATE)
	reader := bufio.NewReader(file)
for {
    Offset := offset
    buf, read_e := reader.ReadBytes('\n')
    offset += int32(len(buf))
    if read_e == io.EOF {
        break
    }
    if read_e != nil {
        return nil, read_e
    }
    line, transform_e := GbkToUtf8(buf)
    if transform_e != nil {
        return nil, transform_e
    }
    if !must.Match(line) {
        continue
    }
    p := empty.ReplaceAll(line, []byte(""))
    Name := strings.SplitAfterN(string(p), " ", 2)[1]
    c := storage.Chapter{Name, Offset}
    Chapters = append(Chapters, c)
}
    stat, stat_e := file.Stat()
    if stat_e != nil {
        return nil, stat_e
    }
    Size := stat.Size()
    book := storage.Book{Chapters, Size}
    return &book, nil
}
