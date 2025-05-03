package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

// 保存するデータの構造体
type User struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    CreatedAt time.Time `json:"created_at"`
}

// ディレクトリを作成する関数
func createDirectory(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return fmt.Errorf("ディレクトリ作成エラー: %v", err)
		}
		fmt.Printf("ディレクトリ %s を作成しました\n", path)
	} else if err != nil {
		fmt.Printf("ディレクトリ %s は既に存在します\n", path)
	}
	return nil
}

// データをJSONファイルに保存する関数
func saveToJSON(filename string, data interface{}) error {
    // データをJSON形式に変換（インデント付き）
    jsonData, err := json.MarshalIndent(data, "", "   ")
    if err != nil {
        return fmt.Errorf("JSON変換エラー: %v", err)
    }

    // ファイルに書き込み
    err = os.WriteFile(filename, jsonData, 0644)
    if err != nil {
        return fmt.Errorf("ファイル書き込みエラー: %v", err)
    }

    return nil
}

// JSONファイルからデータを読み込む関数
func loadFromJSON(filename string, data interface{}) error {
    // ファイルを読み込み
    jsonData, err := os.ReadFile(filename)
    if err != nil {
        return fmt.Errorf("ファイル読み込みエラー: %v", err)
    }

    // JSONをデータ構造体に変換
    err = json.Unmarshal(jsonData, data)
    if err != nil {
        return fmt.Errorf("JSON解析エラー: %v", err)
    }

    return nil
}

func main() {
	// ディレクトリを作成
	dataDir := "data"
	err := createDirectory(dataDir)
	if err != nil {
		log.Fatalf("ディレクトリ作成エラー: %v", err)
	}

    // サンプルデータの作成
    users := []User{
        {
            ID:        1,
            Name:      "山田太郎",
            Email:     "yamada@example.com",
            CreatedAt: time.Now(),
        },
        {
            ID:        2,
            Name:      "鈴木花子",
            Email:     "suzuki@example.com",
            CreatedAt: time.Now(),
        },
    }

    // データをJSONファイルに保存
    userFile := filepath.Join(dataDir, "users.json")
    err = saveToJSON(userFile, users)
    if err != nil {
        log.Fatalf("保存エラー: %v", err)
    }
    fmt.Printf("データを %s に保存しました\n", userFile)

    // JSONファイルからデータを読み込み
    var loadedUsers []User
    err = loadFromJSON(userFile, &loadedUsers)
    if err != nil {
        log.Fatalf("読み込みエラー: %v", err)
    }

    // 読み込んだデータを表示
    fmt.Println("\n読み込んだデータ:")
    for _, user := range loadedUsers {
        fmt.Printf("ID: %d, 名前: %s, メール: %s, 作成日時: %s\n",
            user.ID, user.Name, user.Email, user.CreatedAt.Format("2006-01-02 15:04:05"))
    }
}