package domain

import "time"

type FranjaHoraria struct {
	dia    time.Weekday
	inicio string
	fin    string
}
