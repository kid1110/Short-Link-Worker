package pkg

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/spaolacci/murmur3"
)

func GenerateByMurmurHash(data string) string {
	judge := judgeWebUrl(data)
	var ret string
	if(judge){
		fmt.Println("is url")
		newData := strings.SplitAfter(data, "/")
		res := murmur3.Sum32([]byte(newData[len(newData)-1]))
		newHash :=string(strconv.Itoa(int(res)))
		newData[len(newData)-1] = newHash
		ret = fmt.Sprintf(strings.Join(newData,""))
	}else{
		fmt.Println("not url")
		res := murmur3.Sum32([]byte(data))
		ret = string(strconv.Itoa(int(res)))
	}
	return ret
}

func judgeWebUrl(ul string)bool{
	_, err := url.ParseRequestURI(ul)
	if err != nil{
		return false
	}
	u, err := url.Parse(ul)
	if err != nil || u.Scheme == "" || u.Host == ""{
		return false
	}
	return true
}
