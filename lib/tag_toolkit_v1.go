package lib

type Tag struct {
	Name        string `json:"name" xml:"name"`
	GroupID     string `json:"group_id" xml:"group_id"`
	Notes       string `json:"notes" xml:"notes"`
	Created     string `json:"created" xml:"created"`
	Popularity  int    `json:"popularity" xml:"popularity"`
	SeriesCount int    `json:"series_count" xml:"series_count"`
}

/********************************
 ** GetTags
 **
 ** Get all tags, search for tags,
 ** or get tags by name.
 ********************************/
func (f *FredClient) GetTags(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, tags)

	if err != nil {
		return nil, err
	}

	return fc, nil

}

/********************************
 ** GetRelatedTags
 **
 ** Get the related tags for one
 ** or more tags.
 ********************************/
func (f *FredClient) GetRelatedTags(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, relatedTags)

	if err != nil {
		return nil, err
	}

	return fc, nil

}

/********************************
 ** GetTagSeries
 **
 ** Get the series matching tags.
 ********************************/
func (f *FredClient) GetTagSeries(params map[string]interface{}) (*FredType, error) {

	fc, err := f.operate(params, tagsSeries)

	if err != nil {
		return nil, err
	}

	return fc, nil

}
