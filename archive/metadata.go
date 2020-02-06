package archive

type Metadata struct {
	File   string
	Title  string
	Artist string
	Track  int
	Sum    string
	Type   string

	Size           uint64
	CompressedSize uint64
}

type Archive struct {
	Attr *Attributes

	TrackCount int
	Contents   []*Metadata
}
