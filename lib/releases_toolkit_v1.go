package lib

type Releases struct {
	Start      string    `json:"realtime_start"`
	End        string    `json:"realtime_end"`
	OrderBy    string    `json:"order_by"`
	SortOrder  string    `json:"sort_order"`
	Count      int       `json:"count"`
	Offset     int       `json:"offset"`
	Limit      int       `json:"limit"`
	ReleaseCol []Release `json:"releases"`
}

type Release struct {
	ID           int    `json:"id"`
	Start        string `json:"realtime_start"`
	End          string `json:"realtime_end"`
	Name         string `json:"name"`
	PressRelease bool   `json:"press_release"`
	Link         string `json:"link"`
}

type ReleaseDates struct {
	Start           string        `json:"realtime_start"`
	End             string        `json:"realtime_end"`
	OrderBy         string        `json:"order_by"`
	SortOrder       string        `json:"sort_order"`
	Count           int           `json:"count"`
	Offset          int           `json:"offset"`
	Limit           int           `json:"limit"`
	ReleaseDatesCol []ReleaseDate `json:"release_dates"`
}

type ReleaseDate struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Date string `json:"date"`
}

/********************************
 ** GetReleases
 **
 ** Get all releases of economic
 ** data.
 ********************************/
func (f *FredClient) GetReleases(params map[string]interface{}) (*Releases, error) {

	rls, err := f.operate(params, releases)

	if err != nil {
		return nil, err
	}

	return (rls.(*Releases)), nil

}

/********************************
 ** GetReleasesDates
 **
 ** Get release dates for all
 ** releases of economic data.
 ********************************/
func (f *FredClient) GetReleasesDates(params map[string]interface{}) (*ReleaseDates, error) {

	rlDts, err := f.operate(params, releasesDates)

	if err != nil {
		return nil, err
	}

	return (rlDts.(*ReleaseDates)), nil

}

/********************************
 ** GetRelease
 **
 ** Get a release of economic data.
 ********************************/
func (f *FredClient) GetRelease(params map[string]interface{}) (*Releases, error) {

	rls, err := f.operate(params, release)

	if err != nil {
		return nil, err
	}

	return (rls.(*Releases)), nil

}

/********************************
 ** GetReleaseDates
 **
 ** Get release dates for a release
 ** of economic data.
 ********************************/
func (f *FredClient) GetReleaseDates(params map[string]interface{}) (*ReleaseDates, error) {

	rlDts, err := f.operate(params, releaseDates)

	if err != nil {
		return nil, err
	}

	return (rlDts.(*ReleaseDates)), nil

}

/********************************
 ** GetReleaseSeries
 **
 ** Get the series on a release of
 ** economic data.
 ********************************/
func (f *FredClient) GetReleaseSeries(params map[string]interface{}) (*Seriess, error) {

	rlSrs, err := f.operate(params, releaseSeries)

	if err != nil {
		return nil, err
	}

	return (rlSrs.(*Seriess)), nil

}

/********************************
 ** GetReleaseSources
 **
 ** Get the sources for a Release
 ** of economic data.
 ********************************/
func (f *FredClient) GetReleaseSources(params map[string]interface{}) (*Sources, error) {

	rlSrcs, err := f.operate(params, releaseSources)

	if err != nil {
		return nil, err
	}

	return (rlSrcs.(*Sources)), nil

}

/********************************
 ** GetReleaseTags
 **
 ** Get the tags for a release.
 ********************************/
func (f *FredClient) GetReleaseTags(params map[string]interface{}) (*Tags, error) {

	rlTags, err := f.operate(params, releaseTags)

	if err != nil {
		return nil, err
	}

	return (rlTags.(*Tags)), nil

}

/********************************
 ** GetReleaseRelatedTags
 **
 ** Get the related tags for a
 ** release.
 ********************************/
func (f *FredClient) GetReleaseRelatedTags(params map[string]interface{}) (*Tags, error) {

	rlTags, err := f.operate(params, releaseRelatedTags)

	if err != nil {
		return nil, err
	}

	return (rlTags.(*Tags)), nil

}

/********************************
 ** GetReleaseTables
 **
 ** Get the related tags for a
 ** category.
 ********************************/
/*
func (f *FredClient) GetReleaseTables(params map[string]interface{}) (*Tags, error) {
	if err := f.validateAPIKEY(); err != nil {

		return nil, err

	}

	tags := &Tags{}

	resp, err := f.callAPI(params, "RELEASE_TABLES")

	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(tags)

	if err != nil {
		return nil, errors.New(errorLibraryFail)
	}

	return tags, nil

}
*/