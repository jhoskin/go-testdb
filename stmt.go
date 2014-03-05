package testdb

import (
	"database/sql/driver"
)

type Stmt struct {
	ExecFunc  func(args []driver.Value) (driver.Result, error)
	QueryFunc func(args []driver.Value) (result driver.Rows, err error)
	rows      driver.Rows
	result    driver.Result
	err       error
}

func (s *Stmt) Close() error {
	return nil
}

func (s *Stmt) NumInput() int {
	// This prevents the sql package from validating the number of inputs
	return -1
}

func (s *Stmt) Exec(args []driver.Value) (driver.Result, error) {
	if s.ExecFunc != nil {
		return s.ExecFunc(args)
	}

	return s.result, s.err
}

func (s *Stmt) Query(args []driver.Value) (driver.Rows, error) {
	if s.QueryFunc != nil {
		return s.QueryFunc(args)
	}

	return s.rows, s.err
}
