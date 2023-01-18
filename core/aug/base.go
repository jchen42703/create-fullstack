package aug

type TemplateAugmentor interface {
	Id() string
	Augment() error
}
