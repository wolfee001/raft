package internal

type raftMember struct{}

func (m *raftMember) Write(string) error {
	return nil
}

func NewRaftMember(members []CommunicationInterface, onWrite func(string)) *raftMember {
	return &raftMember{}
}
