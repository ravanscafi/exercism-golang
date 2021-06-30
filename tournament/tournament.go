package tournament

import (
	"encoding/csv"
	"fmt"
	"io"
	"sort"
)

type TeamMap map[string]TeamStats

type TeamStats struct {
	name                string
	wins, draws, losses int
}

func (t TeamStats) matches() int {
	return t.wins + t.draws + t.losses
}

func (t TeamStats) points() int {
	return t.wins*3 + t.draws
}

func Tally(r io.Reader, w io.Writer) error {
	m := make(TeamMap)

	if err := read(r, m); err != nil {
		return err
	}

	l := sortTeams(m)

	return write(w, l)
}

func read(r io.Reader, m TeamMap) error {
	cr := csv.NewReader(r)
	cr.Comma = ';'
	cr.Comment = '#'

	for {
		rec, err := cr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if err := handleRecord(m, rec); err != nil {
			return err
		}
	}

	return nil
}

func handleRecord(m TeamMap, record []string) error {
	if len(record) != 3 {
		return fmt.Errorf("wrong record data: %v", record)
	}

	tOne, tTwo, status := record[0], record[1], record[2]

	if tOne == tTwo {
		return fmt.Errorf("a team can't play with itself: %s", tOne)
	}

	f, s := m[tOne], m[tTwo]
	f.name, s.name = tOne, tTwo

	switch status {
	case "win":
		f.wins++
		s.losses++
		break
	case "loss":
		f.losses++
		s.wins++
		break
	case "draw":
		f.draws++
		s.draws++
		break
	default:
		return fmt.Errorf("wrong status: %s", status)
	}

	m[tOne], m[tTwo] = f, s

	return nil
}

func sortTeams(m TeamMap) []TeamStats {
	l := make([]TeamStats, len(m))
	i := 0

	for _, t := range m {
		l[i] = t
		i++
	}

	sort.SliceStable(l, func(i, j int) bool {
		if l[i].points() == l[j].points() {
			return l[i].name < l[j].name
		}

		return l[i].points() > l[j].points()
	})

	return l
}

func write(w io.Writer, l []TeamStats) error {
	if _, err := fmt.Fprintf(w, "%-30s | MP |  W |  D |  L |  P\n", "Team"); err != nil {
		return err
	}

	for _, t := range l {
		if _, err := fmt.Fprintf(
			w,
			"%-30s |  %d |  %d |  %d |  %d |  %d\n",
			t.name,
			t.matches(),
			t.wins,
			t.draws,
			t.losses,
			t.points(),
		); err != nil {
			return err
		}
	}

	return nil
}
