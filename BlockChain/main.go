package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
    "fmt"
    "strings"
	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
)

// 难度系数 hash前面几个0
const difficulty = 3

// 区块结构体
type Block struct{
    Index int           //这个区块在整个链中的索引
    Timestamp string    //区块生成时的时间戳
    Content string      //需要记录的内容
    Hash string         // 区块通过 SHA256 算法生成的哈希值
    PrevHash string     // 前一个区块的sha256哈希值
    Difficulty int
    Nonce      string
}

// 定义一个数组存储区块数据
var Blockchain []Block

// 接受post过来的参数
type Message struct {
	Content string
}

var mutex = &sync.Mutex{}

//把结构体中的信息拼一起 然后算出hash值
func calculateHash(block Block) string{
              // 数字转换为10进制字符串
    record := strconv.Itoa(block.Index) + block.Timestamp + block.Content + block.PrevHash + block.Nonce
    
    h := sha256.New()
    h.Write([]byte(record))
    hashed := h.Sum(nil)
    return hex.EncodeToString(hashed)
}

// 生成新区块：上一块的索引加一 上一区块的hash赋值给当前区块的prevHash
func generateBlock(oldBlock Block,Content string) Block{
    var newBlock Block
    t := time.Now()
    newBlock.Index = oldBlock.Index+1
    newBlock.Timestamp = t.String()
    newBlock.Content = Content
    newBlock.PrevHash = oldBlock.Hash
    newBlock.Difficulty = difficulty
    for i := 0; ; i++ {
        hex := fmt.Sprintf("%x",i)
        fmt.Println(hex)
        newBlock.Nonce = hex
        if !isHashValid(calculateHash(newBlock),newBlock.Difficulty){
            fmt.Println(calculateHash(newBlock),"继续算！")
            continue
        }else{
            fmt.Println(calculateHash(newBlock),"中！")
            newBlock.Hash = calculateHash(newBlock)
            break
        }
    }
    return newBlock
}

/**
 * 验证区块
    老索引+1 = 新索引
    新的PrevHash = 老的Hash
    验证新区块的hash是否正确
 */
func isBlockValid(newBlock,oldBlock Block) bool{
    if oldBlock.Index+1 != newBlock.Index{
        return false
    }
    if oldBlock.Hash != newBlock.PrevHash{
        return false
    }
    if calculateHash(newBlock) != newBlock.Hash{
        return false
    }
    return true
}


//设置http需要的参数，并开启服务
func run() error{
    mux := makeMuxRouter()
    httpAddr := "8000"
    s := &http.Server{
        Addr : ":"+httpAddr,
        Handler: mux,
        ReadTimeout : 10 * time.Second,
        WriteTimeout : 10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }
    if err := s.ListenAndServe(); err!= nil{
        return err
    }
    return nil
}
//生成NewRouter对象
func makeMuxRouter() http.Handler {
    muxRouter := mux.NewRouter()
    muxRouter.HandleFunc("/", handleGetBlockchain).Methods("GET")
    muxRouter.HandleFunc("/", handleWriteBlock).Methods("POST")
    return muxRouter
}

// 处理get请求
func handleGetBlockchain(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(Blockchain, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

// 把区块的信息组装成json格式返回
func handleWriteBlock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var m Message

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&m); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()

	mutex.Lock()
	newBlock := generateBlock(Blockchain[len(Blockchain)-1], m.Content)
	mutex.Unlock()

	if isBlockValid(newBlock, Blockchain[len(Blockchain)-1]) {
		Blockchain = append(Blockchain, newBlock)
		spew.Dump(Blockchain)
	}

	respondWithJSON(w, r, http.StatusCreated, newBlock)

}

//
func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	response, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)
}

// 判断是否是符合要求的哈希值
func isHashValid(hash string,difficulty int) bool{
    prefix := strings.Repeat("0",difficulty)
    return strings.HasPrefix(hash,prefix)
}


func main() {


	go func() {
		t := time.Now()
		genesisBlock := Block{}
		genesisBlock = Block{0, t.String(), "start", calculateHash(genesisBlock), "",difficulty,""}
		spew.Dump(genesisBlock)

		mutex.Lock()
		Blockchain = append(Blockchain, genesisBlock)
		mutex.Unlock()
	}()
	log.Fatal(run())

}



