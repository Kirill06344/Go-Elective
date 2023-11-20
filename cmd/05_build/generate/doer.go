package generate

//go:generate mockgen -source=doer.go -destination=mock_doer.go -package=generate . Doer

type Doer interface {
	DoSomething(int, string) error
}
