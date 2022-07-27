package instrument

import (
	"gitlab.com/gomidi/midi/v2/smf"
)

type Inst interface {
	mkSMF() *smf.SMF
}

type Drum struct {
	s *smf.SMF
}

type Guitar struct {
	s *smf.SMF
}

func GetSMF(i Inst) error {
	err := i.mkSMF().WriteFile("bass.mid")
	if err != nil {
		return err
	}
	return nil
}
