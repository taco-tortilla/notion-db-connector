package request

type Block interface {
	GetType() string
}

type CustomerFeedbackRequest struct {
	Parent     Parent     `json:"parent"`
	Properties Properties `json:"properties"`
	Children   []Children `json:"children"`
}

type Parent struct {
	DatabaseID string `json:"database_id"`
}

type Properties struct {
	CompanyName  CompanyName  `json:"company_name"`
	ContractType ContractType `json:"contract_type"`
	Plan         Plan         `json:"plan"`
	Read         Read         `json:"read"`
	CreatedBy    CreatedBy    `json:"created_by"`
}

type CompanyName struct {
	Title []Title `json:"title"`
}

type Title struct {
	Text Text `json:"text"`
}

type ContractType struct {
	Select Select `json:"select"`
}

type Plan struct {
	Select Select `json:"select"`
}

type Select struct {
	Name string `json:"name"`
}

type Read struct {
	Checkbox bool `json:"checkbox"`
}

type CreatedBy struct {
	RichText []RichText `json:"rich_text"`
}

type RichText struct {
	Type string `json:"type"`
	Text Text   `json:"text"`
}

type Text struct {
	Content string `json:"content"`
}

type Children struct {
	Object    string     `json:"object"`
	Type      string     `json:"type"`
	Heading2  *Heading2  `json:"heading_2,omitempty"`
	Paragraph *Paragraph `json:"paragraph,omitempty"`
}

type Heading2Block struct {
	Object   string   `json:"object"`
	Type     string   `json:"type"`
	Heading2 Heading2 `json:"heading_2"`
}

type ParagraphBlock struct {
	Object    string    `json:"object"`
	Type      string    `json:"type"`
	Paragraph Paragraph `json:"paragraph"`
}

type Heading2 struct {
	RichText []RichText `json:"rich_text"`
}

type Paragraph struct {
	RichText []RichText `json:"rich_text"`
}
