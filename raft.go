package raft

import "github.com/wolfee001/raft/internal"

type CommunicationInterface = internal.CommunicationInterface

type RaftMember interface {
	Write(string) error
}

func NewRaftMember(members []CommunicationInterface, onWrite func(string)) RaftMember {
	return internal.NewRaftMember(members, onWrite)
}
