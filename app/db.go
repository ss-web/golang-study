package app

import "log"

// Connect - коментарний к методу
func Connect(val int) {
	log.Fatal(val)
	// db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	// if err != nil {
	// 	panic("Соединение с бд не установлено")
	// }
	// defer db.Close()
}
