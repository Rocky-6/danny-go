package instrument

import (
	"gitlab.com/gomidi/midi/v2"
	"gitlab.com/gomidi/midi/v2/gm"
	"gitlab.com/gomidi/midi/v2/smf"
)

type Bass struct {
	s *smf.SMF
}

func (b Bass) mkSMF() *smf.SMF {
	var (
		clock = smf.MetricTicks(96)
		tr    smf.Track
	)

	tr.Add(0, smf.MetaMeter(4, 4))
	tr.Add(0, smf.MetaTempo(182))
	tr.Add(0, smf.MetaInstrument(gm.Instr_ElectricBassPick.String()))
	tr.Add(0, midi.ProgramChange(0, gm.Instr_ElectricBassPick.Value()))
	tr.Add(0, midi.NoteOn(0, midi.G(3), 140))
	tr.Add(clock.Ticks8th(), midi.NoteOff(0, midi.G(3)))
	tr.Add(0, midi.NoteOn(0, midi.G(3), 140))
	tr.Add(clock.Ticks8th(), midi.NoteOff(0, midi.G(3)))
	tr.Add(0, midi.NoteOn(0, midi.D(4), 140))
	tr.Add(clock.Ticks8th(), midi.NoteOff(0, midi.D(4)))
	tr.Add(0, midi.NoteOn(0, midi.G(3), 140))
	tr.Add(clock.Ticks8th(), midi.NoteOff(0, midi.G(3)))
	tr.Add(0, midi.NoteOn(0, midi.G(3), 140))
	tr.Add(clock.Ticks8th(), midi.NoteOff(0, midi.G(3)))
	tr.Add(0, midi.NoteOn(0, midi.D(4), 140))
	tr.Add(clock.Ticks8th(), midi.NoteOff(0, midi.D(4)))
	tr.Add(0, midi.NoteOn(0, midi.G(3), 140))
	tr.Add(clock.Ticks8th(), midi.NoteOff(0, midi.G(3)))
	tr.Add(0, midi.NoteOn(0, midi.G(3), 140))
	tr.Add(clock.Ticks8th(), midi.NoteOff(0, midi.G(3)))

	tr.Add(0, midi.NoteOn(0, midi.E(3), 140))
	tr.Add(clock.Ticks8th(), midi.NoteOff(0, midi.E(3)))
	tr.Add(0, midi.NoteOn(0, midi.E(3), 140))
	tr.Add(clock.Ticks8th(), midi.NoteOff(0, midi.E(3)))
	tr.Add(0, midi.NoteOn(0, midi.B(3), 140))
	tr.Add(clock.Ticks8th(), midi.NoteOff(0, midi.B(3)))
	tr.Add(0, midi.NoteOn(0, midi.E(3), 140))
	tr.Add(clock.Ticks8th(), midi.NoteOff(0, midi.E(3)))
	tr.Add(0, midi.NoteOn(0, midi.E(3), 140))
	tr.Add(clock.Ticks8th(), midi.NoteOff(0, midi.E(3)))
	tr.Add(0, midi.NoteOn(0, midi.B(3), 140))
	tr.Add(clock.Ticks8th(), midi.NoteOff(0, midi.B(3)))
	tr.Add(0, midi.NoteOn(0, midi.E(3), 140))
	tr.Add(clock.Ticks8th(), midi.NoteOff(0, midi.E(3)))
	tr.Add(0, midi.NoteOn(0, midi.E(3), 140))
	tr.Add(clock.Ticks8th(), midi.NoteOff(0, midi.E(3)))

	tr.Close(0)

	b.s = smf.New()
	b.s.TimeFormat = clock
	b.s.Add(tr)
	return b.s
}