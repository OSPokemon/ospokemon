package engine

type Control byte

const (
	CTRLdead    Control = 0x01
	CTRLimmune  Control = 0x02
	CTRLstasis  Control = 0x04
	CTRLstun    Control = 0x08
	CTRLroot    Control = 0x10
	CTRLcloak   Control = 0x20
	CTRLsilence Control = 0x40
	// pseudo control states
	CTRLPprotected Control = CTRLimmune | CTRLstasis | CTRLdead
	CTRLPnocast    Control = CTRLstasis | CTRLstun | CTRLdead | CTRLsilence
	CTRLPstuck     Control = CTRLstasis | CTRLroot | CTRLdead
)
