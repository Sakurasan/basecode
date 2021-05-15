package basecode

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

var (
	DNACODE     = []string{"A", "G", "T", "C"}
	encodeStd   = []string{"A", "G", "C", "U"}
	StdEncoding = NewEncoding(encodeStd)
)

type Encoding struct {
	codeBook []string
	codeMap  map[string]int
	codeLen  int
	sync.RWMutex
}

/* NewEncoding
codebook length greater than or equal to 4
*/
func NewEncoding(codebook []string) *Encoding {
	if len(codebook) < 4 {
		panic(fmt.Sprintf("Invalid length [%d] of [codebook],The length of the codebook is less than 4", len(codebook)))
	}
	e := new(Encoding)
	e.codeBook = codebook
	e.codeMap = make(map[string]int)
	e.codeLen = len(e.codeBook)
	e.initcodemap()

	return e
}

func (e *Encoding) initcodemap() {
	e.codeMap = make(map[string]int)
	for idx, v := range e.codeBook {
		e.codeMap[v] = idx
	}
}

func (e *Encoding) Encode(src string) string {
	var desc string
	hexstr := hex.EncodeToString([]byte(src))
	for i := 0; i < len(hexstr); i++ {
		_num, _ := strconv.ParseInt(string(hexstr[i]), 16, 10)
		answer := int(_num) / e.codeLen
		remainder := int(_num) % e.codeLen
		desc += e.codeBook[answer] + e.codeBook[remainder]
	}
	return desc
}

func (e *Encoding) Decode(src string) string {
	var rawcode []byte
	arr := strings.Split(src, "")
	e.RLock()
	defer e.RUnlock()
	for i := 0; i < len(arr); i += 4 {
		hdeanswer := e.codeMap[arr[i]]
		hderemainder := e.codeMap[arr[i+1]]
		ldeanswer := e.codeMap[arr[i+2]]
		lderemainder := e.codeMap[arr[i+3]]

		rawcode = append(rawcode,
			byte(hdeanswer*e.codeLen+hderemainder)<<4|
				byte(ldeanswer*e.codeLen+lderemainder))
	}
	return string(rawcode)
}
