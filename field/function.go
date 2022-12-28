package field

import (
	"gorm.io/gorm/clause"
)

// Func sql functions
var Func = new(function)

type function struct{}

func (f *function) UnixTimestamp(date ...string) Uint64 {
	if len(date) > 0 {
		return Uint64{expr{e: clause.Expr{SQL: "UNIX_TIMESTAMP(?)", Vars: []interface{}{date[0]}}}}
	}
	return Uint64{expr{e: clause.Expr{SQL: "UNIX_TIMESTAMP()"}}}
}


func (f *function) GetProductTypeFullName(id Int64) String {
	return String{expr{e: clause.Expr{SQL: "getProductTypeFullName(?)", Vars: []interface{}{id.RawExpr()}}}}
}


func (f *function) VersionOrder(version String) Int {
	//INET_ATON(CONCAT(`t_pm_firmware_version`.`version`, '.0'))
	return Int{expr{e: clause.Expr{SQL: "INET_ATON(CONCAT(?, '.0'))", Vars: []interface{}{version.RawExpr()}}}}
}


func (f *function) UnVersionOrder(version Float64) Int {
	//SUBSTRING_INDEX(INET_NTOA(max(INET_ATON(CONCAT(`t_pm_firmware_version`.`version`, '.0')))), '.', 3)
	return Int{expr{e: clause.Expr{SQL: "SUBSTRING_INDEX(INET_NTOA(?), '.', 3)", Vars: []interface{}{version.RawExpr()}}}}
}
