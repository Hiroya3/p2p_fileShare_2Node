package p2p

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"p2p_fileShare_2Node/Node/errorStatus"
	"p2p_fileShare_2Node/Node/service"
	"strings"
	"time"
)

//Run Server
func Run(address, port string) {
	//listenの開始
	listener, err := net.Listen("tcp", address+":"+port)
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()
	log.Println("webサーバーを開始します")

	//1回でlistenerが閉じてしまわないようにfor文かつgoroutineで回す
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("listenerのacceptでエラーが発生しました。\nerr:%s\n", err)
		}
		defer conn.Close()
		go func() {
			// リクエストを読み込む
			messageSlice, err := readRequestMessage(conn)
			if err != nil {
				//エラーコードに応じたハンドリング処理
				switch err {
				case errorStatus.ErrCode300:
					fmt.Println(err)
					break
				default:
					fmt.Println(err)
				}
			}

			if len(messageSlice) > 0 {
				//自分のノードの検索
				searchiedFiles := service.SearchLocalFiles(messageSlice)
				fmt.Printf("これ%s\n", searchiedFiles)
				//connに書き込み検索元に戻す
				writeSearchedFiles(conn, searchiedFiles)
			}

			//ここにくる処理はエラーもmessageSliceもnilのもの＝ヘッダーが想定外のパケットであるため無視
		}()
	}
}

//リクエストを読み込む
func readRequestMessage(conn net.Conn) ([]string, error) {

	conn.SetReadDeadline(time.Now().Add(100 * time.Second))

	res, err := readFromConn(conn)
	if err != nil {
		log.Printf("検索ワードのコネクション読み込みでエラーが発生しました。\nエラー:%s", err)
	}

	//elementsがリクエストの要素
	//elements[0] : 001
	//elements[1] : method
	//elements[2] : body
	//elements[3] : sha256
	elements := strings.Split(string(res[:]), ":")

	//ヘッダーが異なっているものは捨てる
	if elements[0] != "001" {
		return nil, nil
	}

	//改竄がないかハッシュ値の比較
	if !compareHash(elements[0]+":"+elements[1]+":"+elements[2]+":", elements[3]) {
		//TODO:connectionに改竄されたことを書き込む
		err := errorStatus.ReturnErrorCode300()
		return nil, err
	}

	//検索ワードの読み取り
	//,が含まれない場合はelements[2]全体を1つのsliceの要素として返す
	messageSlice := strings.Split(elements[2], ",")

	return messageSlice, nil
}

//ローカルファイルの検索結果をconnに書き込みます
//検索ファイルの場合は001:searchedFiles:[ファイル名]:チェックサム（sha256）
func writeSearchedFiles(connection net.Conn, searchedFilesSli []string) {
	defer connection.Close()
	requestStr := createRequestStr("001", "searchedFiles", searchedFilesSli)
	_, err := connection.Write([]byte(requestStr))
	if err != nil {
		log.Printf("connectionへの書き込みに失敗しました。\nerr:%s", err)
	}
}

func compareHash(requestBodyStr, requestHash string) bool {
	//hash値の計算
	sum := sha256.Sum256([]byte(requestBodyStr))

	return hex.EncodeToString(sum[:]) == requestHash
}

//SearchFile forward to target Node
func SearchFile(address, port string, searchingWords []string) {
	connection, err := net.Dial("tcp", address+":"+port)

	if err != nil {
		log.Printf("query error!!! error:%s", err)
	}

	sendSearchWords(connection, searchingWords)
	slice, err := readRequestMessage(connection)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("検索されたファイル")
	fmt.Println(slice)
}

//connetionに検索ワードを書き込む
//検索の場合は001:searchWords:[キーワード]:チェックサム（sha256）
func sendSearchWords(connection net.Conn, searchingWords []string) {
	if len(searchingWords) == 0 {
		log.Println("キーワードを指定してください")
		return
	}

	requestStr := createRequestStr("001", "searchWords", searchingWords)

	_, err := connection.Write([]byte(requestStr))
	if err != nil {
		log.Printf("connectionへの書き込みに失敗しました。\nerr:%s", err)
	}
}

//リクエストの作成
func createRequestStr(headerNumStr, methodStr string, bodySlice []string) string {
	var messageBui strings.Builder

	messageBui.WriteString(headerNumStr + ":" + methodStr + ":")
	messageBui.WriteString(strings.Join(bodySlice, ",") + ":")

	requestBodyStr := messageBui.String()

	//requestBodyStrのハッシュ値の計算
	sum := sha256.Sum256([]byte(requestBodyStr))
	return requestBodyStr + hex.EncodeToString(sum[:])
}

//connectionからの読み込みの実装
func readFromConn(conn net.Conn) ([]byte, error) {
	res := make([]byte, 4*1024)
	n, err := conn.Read(res)
	if err != nil {
		log.Printf("コネクションからの読み込みに失敗しました。\nエラー：%s", err)
		return nil, err
	}
	return res[:n], nil
}
