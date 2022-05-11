package config

/* jinzhu is for open database from db product i.e. mysql */
import{
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
}

var (
	db * gorm.DB
)

func Connect(){
	d, err := gorm.Open("mysql", "jumpsus:Password/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		panic(err)
	}
	db = d
}

func GetDB(){
	return db
}