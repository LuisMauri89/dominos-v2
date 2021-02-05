package entity

type Entity interface {
	Prepare()
	Validate() error
}
