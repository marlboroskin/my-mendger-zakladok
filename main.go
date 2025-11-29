package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Приложение для закладок
// 1 - посмотреть
// 2 - добавить
// 3 - удалить
// 4 - выйти

type bookmarkMap map[string]string

const dataFile = "bookmarks.json"

func main() {
	bookmarks := loadBookmarks()
	fmt.Println("=== Закладки ===")

	for {
		showMenu()
		choice := getChoice()

		switch choice {
		case 1:
			list(bookmarks)
		case 2:
			add(bookmarks)
			save(bookmarks)
		case 3:
			del(bookmarks)
			save(bookmarks)
		case 4:
			fmt.Println("Пока!")
			return
		default:
			fmt.Println("Только 1-4")
		}

		fmt.Println()
		fmt.Print("Нажмите Enter...")
		fmt.Scanln()
	}
}

func showMenu() {
	fmt.Println("\n--- Меню ---")
	fmt.Println("1. Посмотреть закладки")
	fmt.Println("2. Добавить закладку")
	fmt.Println("3. Удалить закладку")
	fmt.Println("4. Выход")
	fmt.Print("Введите номер закладки: ")
}

func getChoice() int {
	var input string
	fmt.Scanln(&input)
	input = strings.TrimSpace(input)

	num, err := strconv.Atoi(input)
	if err != nil {
		return -1
	}
	return num
}

func list(bm bookmarkMap) {
	fmt.Println("\n📌 Закладки:")
	if len(bm) == 0 {
		fmt.Println("Нету.")
		return
	}
	for k, v := range bm {
		fmt.Printf("  %s → %s\n", k, v)
	}
}

func add(bm bookmarkMap) {
	var name, url string

	fmt.Print("Название: ")
	fmt.Scanln(&name)
	name = strings.TrimSpace(name)
	if name == "" {
		fmt.Println("Пустое имя.")
		return
	}

	fmt.Print("Адрес: ")
	fmt.Scanln(&url)
	url = strings.TrimSpace(url)
	if url == "" {
		fmt.Println("Пустой адрес.")
		return
	}

	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}

	bm[name] = url
	fmt.Println("Добавлено.")
}

func del(bm bookmarkMap) {
	var name string
	fmt.Print("Удалить: ")
	fmt.Scanln(&name)
	name = strings.TrimSpace(name)

	if name == "" {
		fmt.Println("Имя пустое.")
		return
	}

	if _, ok := bm[name]; !ok {
		fmt.Println("Нет такой.")
		return
	}

	delete(bm, name)
	fmt.Println("Удалено.")
}

func loadBookmarks() bookmarkMap {
	data, err := ioutil.ReadFile(dataFile)
	if err != nil {
		fmt.Println("Файл не найден. Создаю новый.")
		return make(bookmarkMap)
	}

	var bm bookmarkMap
	err = json.Unmarshal(data, &bm)
	if err != nil {
		fmt.Println("Ошибка чтения. Новый список.")
		return make(bookmarkMap)
	}
	return bm
}

func save(bookmarks bookmarkMap) {
	data, err := json.MarshalIndent(bookmarks, "", "  ")
	if err != nil {
		return
	}
	ioutil.WriteFile(dataFile, data, 0644)
}
