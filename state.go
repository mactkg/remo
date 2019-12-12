package remo

type Action struct {
    name string
    from *State
    to *State
}

type State interface {
    name string
    actions []*Action
}
