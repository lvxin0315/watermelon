package data_model

type LineData struct {
	Grid    *Grid     `json:"grid"`
	Legend  *Legend   `json:"legend"`
	Series  []*Series `json:"series"`
	Title   *Title    `json:"title"`
	Tooltip *Tooltip  `json:"tooltip"`
	XAxis   *XAxis    `json:"xAxis"`
	YAxis   *YAxis    `json:"yAxis"`
}

type Series struct {
	Data  []int64 `json:"data"`
	Name  string  `json:"name"`
	Stack string  `json:"stack"`
	Type  string  `json:"type"`
}

type Grid struct {
	Bottom       string `json:"bottom"`
	ContainLabel bool   `json:"containLabel"`
	Left         string `json:"left"`
	Right        string `json:"right"`
}

type Legend struct {
	Data []string `json:"data"`
}

type Title struct {
	Text string `json:"text"`
}

type Tooltip struct {
	Trigger string `json:"trigger"`
}

type XAxis struct {
	BoundaryGap bool     `json:"boundaryGap"`
	Data        []string `json:"data"`
	Type        string   `json:"type"`
}

type YAxis struct {
	Type string `json:"type"`
}
