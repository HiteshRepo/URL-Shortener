package tracker

type Domain struct {
	Name  string
	Count int
}

type DomainTrackerer interface {
	GetTopN() []Domain
	Visit(domainName string)
}
