package dao

import "fmt"

const (
	_selectJob = "SELECT (name,sex,age) FROM job WHERE name=?"
)

type Resp struct {
	Name string
	Sex  string
	Age  int
}

func (d *MainDao) GetJob() *Resp {

	resp := &Resp{}

	rows, err := d.DB.Query(_selectJob, "zbc")
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return resp
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&resp.Name, &resp.Sex, &resp.Age); err != nil {
			fmt.Printf("err:%v\n", err)
			return resp
		}
	}

	return resp

}
