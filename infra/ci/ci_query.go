package ci

import "fmt"

func QueryV1(target_ci *TargetV1) error {

	st := target_ci.GitPackage.Strategy

	var err error

	dest_loc := ".npia.infra/" + target_ci.GitPackage.Name

	src_loc := ".npia.infra/" + target_ci.Root

	dir_tb, file_tb, err := TableBuilderV1(dest_loc, src_loc, target_ci)

	if err != nil {

		return fmt.Errorf("query v1: table builder: %s", err.Error())

	}

	if st == "reset" {

		err = StrategyResetV1(dest_loc, src_loc, &dir_tb, &file_tb)

	} else {

		return fmt.Errorf("query v1: no such strategy: %s", st)
	}

	if err != nil {

		return fmt.Errorf("query v1: %s", err.Error())

	}

	return nil

}

func TableBuilderV1(dest_loc string, src_loc string, target_ci *TargetV1) (DIR_TABLE, FILE_TABLE, error) {

	var dt DIR_TABLE
	var ft FILE_TABLE

	return dt, ft, nil

}

func StrategyResetV1(dest_loc string, src_loc string, dir_tb *DIR_TABLE, file_tb *FILE_TABLE) error {

	return nil
}
