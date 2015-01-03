package kala

import "time"

var (
	Truti            time.Duration = time.Nanosecond * 0.031
	Renu             time.Duration = 60 * Truti
	Lava             time.Duration = 60 * Renu
	Leekshaka        time.Duration = 60 * Lava
	Lipta            time.Duration = 60 * Leekshaka
	Vipala           time.Duration = 60 * Leekshaka
	Pala             time.Duration = 60 * Lipta
	Vighati          time.Duration = 60 * Lipta
	Vinadi           time.Duration = 60 * Lipta
	Ghati            time.Duration = 60 * Vighati
	Nadi             time.Duration = 60 * Vighati
	Danda            time.Duration = 60 * Vighati
	Muhurta          time.Duration = 2 * Ghati
	NaksatraAhoratam               = 30 * Muhurta

	// Alternate System
	TrutiAlt            time.Duration = time.Nanosecond * 35.5
	Tatpara             time.Duration = TrutiAlt * 100
	Nimesha             time.Duration = Tatpara * 30
	Kastha              time.Duration = Nimesha * 30
	Kala                time.Duration = Kastha * 30
	MuhurtaAlt          time.Duration = Kala * 30
	NaksatraAhoratamAlt time.Duration = MuhurtaAlt * 30

	//Small units of time used in Vedas
	Paramanu             time.Duration = time.Nanosecond * 26.3
	Anu                  time.Duration = 2 * Paramanu
	Trasarenu            time.Duration = 3 * Anu
	TrutiVeda            time.Duration = 3 * Trasarenu
	Vedha                time.Duration = 100 * TrutiVeda
	LavaVeda             time.Duration = 3 * Vedha
	NimeshaVeda          time.Duration = 3 * LavaVeda
	Ksana                time.Duration = 3 * NimeshaVeda
	KashtaVeda           time.Duration = 5 * Ksana
	Laghu                time.Duration = 15 * KashtaVeda
	DandaVeda            time.Duration = 15 * Laghu
	MuhurtaVeda          time.Duration = 2 * DandaVeda
	NaksatraAhoratamVeda time.Duration = 30 * MuhurtaVeda
	Masa                 time.Duration = 30 * NaksatraAhoratamVeda
	Ritu                 time.Duration = 2 * Masa
	Ayana                time.Duration = 3 * Ritu
	Samvatsara           time.Duration = 2 * Ayana
)
