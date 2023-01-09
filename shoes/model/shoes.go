package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
	"github.com/satori/go.uuid"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
	_ = uuid.UUID{}
)

/*
DB Table Details
-------------------------------------


CREATE TABLE `shoes` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `brand` varchar(255) DEFAULT NULL,
  `model` varchar(255) DEFAULT NULL,
  `occasion` varchar(255) DEFAULT NULL,
  `favorite` tinyint(1) DEFAULT NULL,
  `pic` varchar(255) DEFAULT NULL,
  `picture` mediumblob,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=60 DEFAULT CHARSET=utf8mb3

JSON Sample
-------------------------------------
{    "id": 29,    "brand": "RShpwThAcDuyXYlKkhEqyxEnZ",    "model": "rieviFhXURExoHrsELXhAagqX",    "occasion": "bgsCKKOAdkdcBMZCBYsEHwHhu",    "favorite": 38,    "pic": "xRtiCXYATbgHhjlVUFoobIHuX",    "picture": "AUIfPFpeP04wHF5fYDAeMgVjA0pJGgQwMwM7G2FZAyJYXiwlRitURzsPUCAcOCcbRDY4Cz0oNUATJxVGSyYwTT8HPSJdGB5cAWALQzxiHyoHUA4XCjhj",    "created_at": "2257-05-16T04:31:44.715969549-04:00",    "updated_at": "2243-06-27T14:24:20.917684743-04:00"}



*/

// Shoes struct is a row record of the shoes table in the shoes_development database
type Shoes struct {
	//[ 0] id                                             bigint               null: false  primary: true   isArray: false  auto: true   col: bigint          len: -1      default: []
	ID int64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:bigint;" json:"id"`
	//[ 1] brand                                          varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Brand null.String `gorm:"column:brand;type:varchar;size:255;" json:"brand"`
	//[ 2] model                                          varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Model null.String `gorm:"column:model;type:varchar;size:255;" json:"model"`
	//[ 3] occasion                                       varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Occasion null.String `gorm:"column:occasion;type:varchar;size:255;" json:"occasion"`
	//[ 4] favorite                                       tinyint              null: true   primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: []
	Favorite null.Int `gorm:"column:favorite;type:tinyint;" json:"favorite"`
	//[ 5] pic                                            varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	Pic null.String `gorm:"column:pic;type:varchar;size:255;" json:"pic"`
	//[ 6] picture                                        blob                 null: true   primary: false  isArray: false  auto: false  col: blob            len: -1      default: []
	Picture []byte `gorm:"column:picture;type:blob;" json:"picture"`
	//[ 7] created_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;" json:"created_at"`
	//[ 8] updated_at                                     datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;" json:"updated_at"`
}

var shoesTableInfo = &TableInfo{
	Name: "shoes",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "bigint",
			DatabaseTypePretty: "bigint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "bigint",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "int64",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "int64",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "brand",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Brand",
			GoFieldType:        "null.String",
			JSONFieldName:      "brand",
			ProtobufFieldName:  "brand",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "model",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Model",
			GoFieldType:        "null.String",
			JSONFieldName:      "model",
			ProtobufFieldName:  "model",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "occasion",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Occasion",
			GoFieldType:        "null.String",
			JSONFieldName:      "occasion",
			ProtobufFieldName:  "occasion",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "favorite",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "Favorite",
			GoFieldType:        "null.Int",
			JSONFieldName:      "favorite",
			ProtobufFieldName:  "favorite",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "pic",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "Pic",
			GoFieldType:        "null.String",
			JSONFieldName:      "pic",
			ProtobufFieldName:  "pic",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "picture",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "blob",
			DatabaseTypePretty: "blob",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "blob",
			ColumnLength:       -1,
			GoFieldName:        "Picture",
			GoFieldType:        "[]byte",
			JSONFieldName:      "picture",
			ProtobufFieldName:  "picture",
			ProtobufType:       "bytes",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "created_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "CreatedAt",
			GoFieldType:        "time.Time",
			JSONFieldName:      "created_at",
			ProtobufFieldName:  "created_at",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "updated_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "UpdatedAt",
			GoFieldType:        "time.Time",
			JSONFieldName:      "updated_at",
			ProtobufFieldName:  "updated_at",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        9,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *Shoes) TableName() string {
	return "shoes"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *Shoes) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *Shoes) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *Shoes) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *Shoes) TableInfo() *TableInfo {
	return shoesTableInfo
}
