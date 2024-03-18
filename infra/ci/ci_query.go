package ci

import (
	"fmt"
	"os"
	"strings"

	infutils "github.com/OKESTRO-AIDevOps/nkia/infra/utils"
)

func QueryV1(target_ci *TargetV1) error {

	st := target_ci.GitPackage.Strategy

	var err error

	dest_loc := ".npia.infra/" + target_ci.GitPackage.Name

	src_loc := ".npia.infra/" + target_ci.Root

	ndir_tb, nfile_tb, dir_tb, file_tb, err := TableBuilderV1(src_loc, target_ci)

	if err != nil {

		return fmt.Errorf("query v1: table builder: %s", err.Error())

	}

	if st == "reset" {

		err = ExecuteResetV1(dest_loc, &ndir_tb, &nfile_tb, &dir_tb, &file_tb)

		// err = ExecuteResetV1_Test(target_ci.GitPackage.Name, dest_loc, &ndir_tb, &nfile_tb, &dir_tb, &file_tb)

		if err != nil {

			return fmt.Errorf("v1 reset: %s", err.Error())

		}

	} else {

		return fmt.Errorf("query v1: no such strategy: %s", st)
	}

	if err != nil {

		return fmt.Errorf("query v1: %s", err.Error())

	}

	return nil

}

func TableBuilderV1(src_loc string, target_ci *TargetV1) (N_DIR_TABLE, N_FILE_TABLE, DIR_TABLE, FILE_TABLE, error) {

	var ndt N_DIR_TABLE
	var nft N_FILE_TABLE
	var dt DIR_TABLE
	var ft FILE_TABLE

	locked_ones := target_ci.GitPackage.Lock

	select_len := len(target_ci.Select)

	for i := 0; i < select_len; i++ {

		ndt_tmp, nft_tmp, dt_tmp, ft_tmp, err := SelectV1(src_loc, locked_ones, &target_ci.Select[i])

		if err != nil {

			return ndt, nft, dt, ft, fmt.Errorf("failed to select v1: %s", err.Error())

		}

		ndt = append(ndt, ndt_tmp...)

		nft = append(nft, nft_tmp...)

		dt = append(dt, dt_tmp...)

		ft = append(ft, ft_tmp...)

	}

	return ndt, nft, dt, ft, nil

}

func SelectV1(src_loc string, locked_ones []string, selector *TargetV1Select) (N_DIR_TABLE, N_FILE_TABLE, DIR_TABLE, FILE_TABLE, error) {

	var ndt N_DIR_TABLE
	var nft N_FILE_TABLE

	var dt DIR_TABLE
	var ft FILE_TABLE

	var dt_s DIR_TABLE
	var ft_s FILE_TABLE

	from_src_loc := src_loc + "/" + selector.From

	w_flag, err := SelectV1_WhatFlagHelper(selector)

	if err != nil {

		return ndt, nft, dt, ft, fmt.Errorf("failed to infer what flag: %s", err.Error())

	}

	a_flag, err := SelectV1_AsFlagHelper(selector)

	if err != nil {

		return ndt, nft, dt, ft, fmt.Errorf("failed to infer as flag: %s", err.Error())
	}

	if w_flag != "*" {

		return ndt, nft, dt, ft, fmt.Errorf("what flag not implemented: %s", w_flag)

	}

	dt_s, ft_s, err = SelectV1_RecursiveHelper(from_src_loc)

	if a_flag != "*" {

		return ndt, nft, dt, ft, fmt.Errorf("as flag not implemented: %s", a_flag)

	}

	dt_s_len := len(dt_s)

	ft_s_len := len(ft_s)

	for i := 0; i < dt_s_len; i++ {

		tmp := strings.Replace(dt_s[i], from_src_loc+"/", "", 1)

		if found := infutils.FindFromSlice[string](locked_ones, tmp); found >= 0 {

			continue

		}

		dt = append(dt, dt_s[i])

	}

	for i := 0; i < ft_s_len; i++ {

		tmp := strings.Replace(ft_s[i], from_src_loc+"/", "", 1)

		if found := infutils.FindFromSlice[string](locked_ones, tmp); found >= 0 {

			continue

		}

		ft = append(ft, ft_s[i])
	}

	dt_len := len(dt)

	ft_len := len(ft)

	for i := 0; i < dt_len; i++ {

		tmp := strings.Replace(dt[i], from_src_loc+"/", "", 1)

		ndt = append(ndt, tmp)

	}

	for i := 0; i < ft_len; i++ {

		tmp := strings.Replace(ft[i], from_src_loc+"/", "", 1)

		nft = append(nft, tmp)

	}

	return ndt, nft, dt, ft, nil

}

func SelectV1_WhatFlagHelper(selector *TargetV1Select) (string, error) {

	var q_flag string

	q_flag = "*"

	return q_flag, nil

}

func SelectV1_AsFlagHelper(selector *TargetV1Select) (string, error) {

	var q_flag string

	q_flag = "*"

	return q_flag, nil

}

func SelectV1_RecursiveHelper(src_loc string) (DIR_TABLE, FILE_TABLE, error) {

	var dt DIR_TABLE
	var ft FILE_TABLE

	src_dir_entry, err := os.ReadDir(src_loc)

	if err != nil {

		return dt, ft, fmt.Errorf("source dir: %s", err.Error())
	}

	src_dir_len := len(src_dir_entry)

	for i := 0; i < src_dir_len; i++ {

		if src_dir_entry[i].Name() == ".git" {
			continue
		}

		src_loc_next := src_loc + "/" + src_dir_entry[i].Name()

		if src_dir_entry[i].IsDir() {

			dt = append(dt, src_loc_next)

			dt_r, ft_r, err_r := SelectV1_RecursiveHelper(src_loc_next)

			if err_r != nil {

				return DIR_TABLE{}, FILE_TABLE{}, fmt.Errorf("failed to recursively build table v1: %s", err_r.Error())

			}

			dt = append(dt, dt_r...)

			ft = append(ft, ft_r...)

		} else {

			ft = append(ft, src_loc_next)

		}

	}

	return dt, ft, nil

}

func ExecuteResetV1(dest_loc string, ndir_tb *N_DIR_TABLE, nfile_tb *N_FILE_TABLE, dir_tb *DIR_TABLE, file_tb *FILE_TABLE) error {

	err := ExecuteResetV1_DeleteWildCard(dest_loc)

	if err != nil {

		return fmt.Errorf("failed to execute reset v1: delete: %s", err.Error())
	}

	err = ExecuteResetV1_ForceCreateAtDest(dest_loc, ndir_tb, dir_tb)

	if err != nil {

		return fmt.Errorf("failed to execute reset v1: create dir: %s", err.Error())
	}

	err = ExecuteResetV1_ForceCopyToDest(dest_loc, nfile_tb, file_tb)

	if err != nil {
		return fmt.Errorf("failed to execute reset v1: copy file: %s", err.Error())
	}

	return nil
}

func ExecuteResetV1_DeleteWildCard(dest_loc string) error {

	dest_dir_entry, err := os.ReadDir(dest_loc)

	if err != nil {

		return fmt.Errorf("dest dir: %s", err.Error())
	}

	dest_dir_len := len(dest_dir_entry)

	for i := 0; i < dest_dir_len; i++ {

		top_name := dest_dir_entry[i].Name()

		full_name := dest_loc + "/" + top_name

		if top_name == ".git" {

			continue
		}

		err := os.RemoveAll(full_name)

		if err != nil {

			return fmt.Errorf("failed to delete wildcard: %s", err.Error())

		}

	}

	return nil
}

func ExecuteResetV1_ForceCreateAtDest(dest_loc string, ndir_tb *N_DIR_TABLE, dir_tb *DIR_TABLE) error {

	ndir_len := len(*ndir_tb)

	dir_len := len(*dir_tb)

	if ndir_len != dir_len {

		return fmt.Errorf("v1 force create at dest: length not matched: %d", ndir_len-dir_len)

	}

	total_dir_len := ndir_len

	for i := 0; i < total_dir_len; i++ {

		dest_dir_name := dest_loc + "/" + (*ndir_tb)[i]

		// src_dir_name := (*dir_tb)[i]

		err := os.MkdirAll(dest_dir_name, 0644)

		if err != nil {

			return fmt.Errorf("v1 force create at dest: %s", err.Error())
		}

	}

	return nil
}

func ExecuteResetV1_ForceCopyToDest(dest_loc string, nfile_tb *N_FILE_TABLE, file_tb *FILE_TABLE) error {

	nfile_len := len(*nfile_tb)

	file_len := len(*file_tb)

	if nfile_len != file_len {

		return fmt.Errorf("v1 force copy to dest: length not matched: %d", nfile_len-file_len)
	}

	total_file_len := nfile_len

	for i := 0; i < total_file_len; i++ {

		dest_file_name := dest_loc + "/" + (*nfile_tb)[i]

		src_file_name := (*file_tb)[i]

		file_b, err := os.ReadFile(src_file_name)

		if err != nil {

			return fmt.Errorf("v1 force copy to dest: read: %s", err.Error())
		}

		err = os.WriteFile(dest_file_name, file_b, 0644)

		if err != nil {

			return fmt.Errorf("v1 force copy to dest: write: %s", err.Error())
		}

	}

	return nil
}

func ExecuteResetV1_Test(fname string, dest_loc string, ndir_tb *N_DIR_TABLE, nfile_tb *N_FILE_TABLE, dir_tb *DIR_TABLE, file_tb *FILE_TABLE) error {

	nd_len := len(*ndir_tb)

	nf_len := len(*nfile_tb)

	d_len := len(*dir_tb)

	f_len := len(*file_tb)

	fptr, err := os.OpenFile(fname, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {

		return fmt.Errorf("failed to open file: %s", err.Error())
	}

	fptr.Write([]byte("NDIR\n"))

	for i := 0; i < nd_len; i++ {

		new_str := (*ndir_tb)[i] + "\n"

		fptr.Write([]byte(new_str))

	}

	fptr.Write([]byte("NFILE\n"))

	for i := 0; i < nf_len; i++ {

		new_str := (*nfile_tb)[i] + "\n"

		fptr.Write([]byte(new_str))

	}

	fptr.Write([]byte("DIR\n"))

	for i := 0; i < d_len; i++ {

		new_str := (*dir_tb)[i] + "\n"

		fptr.Write([]byte(new_str))

	}

	fptr.Write([]byte("FILE\n"))

	for i := 0; i < f_len; i++ {

		new_str := (*file_tb)[i] + "\n"

		fptr.Write([]byte(new_str))

	}

	return nil

}
