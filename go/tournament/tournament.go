package tournament

import (
	"fmt"
	"io"
	"sort"
	"strings"
)

// Stat type stores team results
type Stat struct {
	Name string
	MP   int
	W    int
	D    int
	L    int
	P    int
}

// Tally the results of a small football competition
func Tally(reader io.Reader, writer io.Writer) error {
	records, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	statMap := make(map[string]Stat)

	for _, record := range strings.Split(string(records), "\n") {

		// exclude empty lines and comments
		if record == "" || strings.HasPrefix(record, "#") {
			continue
		}

		// Extract teams and match result for record
		elements := strings.Split(record, ";")
		if len(elements) != 3 {
			return fmt.Errorf("bad record: %s", record)
		}
		homeTeam, awayTeam, result := elements[0], elements[1], elements[2]

		// Calculate Stat
		homeStat := statMap[homeTeam]
		homeStat.Name = homeTeam
		homeStat.MP++

		awayStat := statMap[awayTeam]
		awayStat.Name = awayTeam
		awayStat.MP++

		switch result {
		case "win":
			homeStat.W++
			homeStat.P += 3
			awayStat.L++
		case "loss":
			awayStat.W++
			awayStat.P += 3
			homeStat.L++
		case "draw":
			homeStat.D++
			homeStat.P++
			awayStat.D++
			awayStat.P++
		default:
			return fmt.Errorf("unknown match result: %s", result)
		}
		statMap[homeTeam] = homeStat
		statMap[awayTeam] = awayStat

	}

	// create a list of Stat for map of Stat
	statList := make([]Stat, 0)
	for _, stat := range statMap {
		statList = append(statList, stat)
	}

	// sort by Points then by Name
	sort.Slice(statList, func(i, j int) bool {
		if statList[i].P == statList[j].P {
			return statList[i].Name < statList[j].Name
		}
		return statList[i].P > statList[j].P
	})

	// write the Stat list to output
	_, _ = io.WriteString(writer, "Team                           | MP |  W |  D |  L |  P\n")
	if err != nil {
		return err
	}
	for _, stat := range statList {
		_, _ = fmt.Fprintf(writer, "%-31s|%3d |%3d |%3d |%3d |%3d\n", stat.Name, stat.MP, stat.W, stat.D, stat.L, stat.P)
	}
	return nil
}
