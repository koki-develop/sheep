package model

import "strings"

type sheep struct {
	Asleep []string
	Awake  string
}

var simpleSheep = sheep{
	Asleep: []string{
		strings.Join([]string{
			"           __  _",
			"       .-:'  `; `-._",
			"      (_,           )",
			"    ,'-\"(            )>",
			"   (__,-'            )",
			"      (             )",
			"       `-'._.--._.-'",
			"",
		}, "\n"),
		strings.Join([]string{
			"           __  _",
			"       .-:'  `; `-._",
			"   z  (_,           )",
			"    ,'-\"(            )>",
			"   (__,-'            )",
			"      (             )",
			"       `-'._.--._.-'",
			"",
		}, "\n"),
		strings.Join([]string{
			"           __  _",
			"       .-:'  `; `-._",
			"  zz  (_,           )",
			"    ,'-\"(            )>",
			"   (__,-'            )",
			"      (             )",
			"       `-'._.--._.-'",
			"",
		}, "\n"),
		strings.Join([]string{
			"           __  _",
			"       .-:'  `; `-._",
			" zzz  (_,           )",
			"    ,'-\"(            )>",
			"   (__,-'            )",
			"      (             )",
			"       `-'._.--._.-'",
			"",
		}, "\n"),
	},
	Awake: strings.Join([]string{
		"           __  _",
		"       .-:'  `; `-._",
		"      (_,           )",
		"    ,'o\"(            )>",
		"   (__,-'            )",
		"      (             )",
		"       `-'._.--._.-'",
		"",
	}, "\n"),
}
