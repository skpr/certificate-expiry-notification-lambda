package slack

// BlockType is used to declare the type of a block.
type BlockType string

const (
	// BlockTypeContext used to declare a block as context.
	BlockTypeContext BlockType = "context"
	// BlockTypeSection used to declare a block as section.
	BlockTypeSection BlockType = "section"
	// BlockTypeDivider used to declare a block as divider.
	BlockTypeDivider BlockType = "divider"
)

// BlockElementType is used to declare the type of an element.
type BlockElementType string

const (
	// BlockElementTypeMarkdown used to declare a block as markdown.
	BlockElementTypeMarkdown BlockElementType = "mrkdwn"
	// BlockElementTypeImage used to declare a block as image.
	BlockElementTypeImage BlockElementType = "image"
)

// BlockTextType is used to declare the type of text.
type BlockTextType string

const (
	// BlockTextTypeMarkdown used to declare a block as markdown.
	BlockTextTypeMarkdown BlockTextType = "mrkdwn"
)

// Message to be sent to Slack.
type Message struct {
	Blocks []interface{} `json:"blocks"`
}

// BlockContext used to declare a block as context.
type BlockContext struct {
	Type     BlockType             `json:"type"`
	Elements []BlockContextElement `json:"elements"`
}

// BlockContextElement used to declare an element in a context block.
type BlockContextElement struct {
	Type BlockElementType `json:"type"`
	Text string           `json:"text"`
}

// BlockDivider used to declare a block as divider.
type BlockDivider struct {
	Type BlockType `json:"type"`
}

// BlockSection used to declare a block as section.
type BlockSection struct {
	Type      BlockType              `json:"type"`
	Text      BlockSectionText       `json:"text"`
	Accessory *BlockSectionAccessory `json:"accessory,omitempty"`
}

// BlockSectionAccessory used to declare an accessory in a section block.
type BlockSectionAccessory struct {
	Type     BlockElementType `json:"type"`
	ImageURL string           `json:"image_url"`
	AltText  string           `json:"alt_text"`
}

// BlockSectionText used to declare text in a section block.
type BlockSectionText struct {
	Type BlockTextType `json:"type"`
	Text string        `json:"text"`
}
