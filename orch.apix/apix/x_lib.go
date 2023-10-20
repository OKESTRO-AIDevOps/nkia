package apix

func LinearInstructionBuildFromRawMap(std_cmd_in map[string]string) string {

	var lininst string

	count := 1

	cmd_len := len(std_cmd_in)

	for _, v := range std_cmd_in {

		if count == 1 {
			lininst = v + ":"
			continue
		}

		lininst += v + ","

		if count == cmd_len {
			lininst += v
			break
		}

		count += 1

	}

	return lininst
}
