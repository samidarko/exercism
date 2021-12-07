package tournament

import (
	"fmt"
	"io"
	"sort"
	"strings"
)

type Stat struct {
	name string
	MP   int
	W    int
	D    int
	L    int
	P    int
}

func Tally(reader io.Reader, writer io.Writer) error {
	data, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	statMap := make(map[string]Stat)

	records := strings.Split(strings.Trim(string(data), "\n"), "\n")
	for _, record := range records {
		elements := strings.Split(record, ";")
		homeTeam, awayTeam, result := elements[0], elements[1], elements[2]
		homeStat := statMap[homeTeam]
		homeStat.name = homeTeam
		homeStat.MP += 1

		awayStat := statMap[awayTeam]
		awayStat.name = awayTeam
		awayStat.MP += 1

		switch result {
		case "win":
			homeStat.W += 1
			homeStat.P += 3
			awayStat.L += 1
		case "loss":
			awayStat.W += 1
			awayStat.P += 3
			homeStat.L += 1
		case "draw":
			homeStat.D += 1
			homeStat.P += 1
			awayStat.D += 1
			awayStat.P += 1
		}
		statMap[homeTeam] = homeStat
		statMap[awayTeam] = awayStat

	}

	statList := make([]Stat, 0)

	for _, stat := range statMap {
		statList = append(statList, stat)
	}

	sort.Slice(statList, func(i, j int) bool {
		return statList[i].P > statList[j].P
	})
	_, _ = io.WriteString(writer, "Team                           | MP |  W |  D |  L |  P\n")
	if err != nil {
		return err
	}
	for _, stat := range statList {
		_, _ = fmt.Fprintf(writer, "%-31s|%3d |%3d |%3d |%3d |%3d\n", stat.name, stat.MP, stat.W, stat.D, stat.L, stat.P)
	}
	return nil
}
