package charts

// Cfg structure concentrates chart configuration.
type Cfg struct {
	ChartType         uint
	TemplatePath      string
	JSLibraryPath     string
	ScreenCapturePath string
}

// ChartDataset type defined for easy refactoring if to move away from null interface.
type ChartDataset [][]interface{}

// Chart structure concentrates data for rendering charts.
// Dataset number of series should be limited to three.
type Chart struct {
	Cfg
	Dataset ChartDataset
}

// NewChart
func NewChart(cfg Cfg, data ChartDataset) (*Chart, error) {
	// TODO: add validations got configuration and dataset

	return &Chart{
		Cfg:     cfg,
		Dataset: data,
	}, nil
}

func (c *Chart) Render() error {

}
