package lib

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type lineGroup struct {
	lines   []string
	game_id string
	world   string
}

func ReadLog(name string) []Game {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	name_match, _ := regexp.Compile(`user_id ([a-z0-9\-]{36}) and display name (\w+),`)
	game_match, _ := regexp.Compile(`game_id ([a-z0-9\-]{36})`)

	var world, display_name, game_id, user_id string
	groups := make([]lineGroup, 0)
	lines := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "display name") {
			user_id = name_match.FindStringSubmatch(line)[1]
			display_name = name_match.FindStringSubmatch(line)[2]
		}

		if strings.Contains(line, "game_id") {
			if len(lines) > 0 {
				groups = append(groups, lineGroup{lines: lines, game_id: game_id, world: world})
			}
			lines = make([]string, 0)
			matches := game_match.FindStringSubmatch(line)
			if len(matches) > 1 {
				game_id = matches[1]
			}
		}

		if strings.Contains(line, "Loading '") {
			world = strings.Trim(strings.Split(strings.TrimPrefix(line, "Loading '"), "/")[1], "'")
		}

		if strings.Contains(line, "Unit ") {
			lines = append(lines, line)
		}
	}

	if len(lines) > 0 {
		groups = append(groups, lineGroup{lines: lines, game_id: game_id, world: world})
	}

	player := Player{Id: uuid.MustParse(user_id), Name: display_name}
	games := make([]Game, 0)

	for _, group := range groups {
		games = append(games,
			Game{Id: uuid.MustParse(group.game_id),
				World:  group.world,
				Player: player,
				Units:  CategorizeLines(group.lines)})
	}
	return games
}

func CategorizeLines(lines []string) []Unit {
	events := make(map[int][]string)

	unit_number_match, _ := regexp.Compile(`Unit (\d+) `)
	player_match, _ := regexp.Compile(`Player (\d)`)
	type_match, _ := regexp.Compile(`Type: (\d+)`)

	for _, line := range lines {
		n, _ := strconv.Atoi(unit_number_match.FindStringSubmatch(line)[1])
		if _, ok := events[n]; !ok {
			events[n] = make([]string, 0)
		}
		events[n] = append(events[n], line)
	}

	// fmt.Println(events)
	units := make([]Unit, 0)

	for n, v := range events {
		unit := Unit{Id: n}

		for _, line := range v {
			if strings.Contains(line, "Type: ") {
				unit.Type = type_match.FindStringSubmatch(line)[1]
			}
			if strings.Contains(line, "Player ") {
				unit.Player = player_match.FindStringSubmatch(line)[1]
			}
			if strings.Contains(line, "spawned") {
				unit.Spawned = true
			}
			if strings.Contains(line, "death") {
				unit.Died = true
			}
		}

		units = append(units, unit)
	}

	sort.Sort(Units(units))
	return units
}
