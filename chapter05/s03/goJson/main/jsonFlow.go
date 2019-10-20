package main

import (
	"encoding/json"
	"log"
	"os"
)

//json 的流式读写
//Go内建的encoding/json 包还提供Decoder和Encoder两个类型，用于支持JSON数据的
//流式读写，并提供NewDecoder()和NewEncoder()两个函数来便于具体实现：

func JsonFlow() {
	dec := json.NewDecoder(os.Stdin)
	enc := json.NewEncoder(os.Stdout)
	for {
		var v map[string]interface{}
		if err := dec.Decode(&v); err != nil {
			log.Println(err)
			return
		}
		for k := range v {
			if k != "Title" {
				//v[k] = nil, false
			}
		}
		if err := enc.Encode(&v); err != nil {
			log.Println(err)
		}
	}
}
