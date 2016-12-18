package lib_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetCategory(t *testing.T) {

	params := make(map[string]interface{})

	Convey("", t, func() {
		ctg, err := xmlFredClient.GetCategory(params)
		So(err, ShouldBeNil)

		Convey("", t, func() {
			So(ctg, ShouldNotBeNil)
		})
	})

	Convey("", t, func() {
		ctg, err := jsonFredClient.GetCategory(params)

		So(err, ShouldBeNil)
		Convey("", t, func() {
			So(ctg, ShouldNotBeNil)
		})
	})

}

func TestGetCategoryChildren(t *testing.T) {

	params := make(map[string]interface{})

	Convey("", t, func() {
		ctgs, err := xmlFredClient.GetCategoryChildren(params)
		So(err, ShouldBeNil)

		Convey("", t, func() {
			So(ctgs, ShouldNotBeNil)
		})
	})

	Convey("", t, func() {
		ctgs, err := jsonFredClient.GetCategoryChildren(params)

		So(err, ShouldBeNil)
		Convey("", t, func() {
			So(ctgs, ShouldNotBeNil)
		})
	})

}

func TestGetRelatedCategory(t *testing.T) {

	params := make(map[string]interface{})

	Convey("", t, func() {
		ctg, err := xmlFredClient.GetRelatedCategory(params)
		So(err, ShouldBeNil)

		Convey("", t, func() {
			So(ctg, ShouldNotBeNil)
		})
	})

	Convey("", t, func() {
		ctg, err := jsonFredClient.GetRelatedCategory(params)

		So(err, ShouldBeNil)
		Convey("", t, func() {
			So(ctg, ShouldNotBeNil)
		})
	})

}

func TestGetCategorySeries(t *testing.T) {

	params := make(map[string]interface{})

	Convey("", t, func() {
		srs, err := xmlFredClient.GetCategorySeries(params)
		So(err, ShouldBeNil)

		Convey("", t, func() {
			So(srs, ShouldNotBeNil)
		})
	})

	Convey("", t, func() {
		srs, err := jsonFredClient.GetCategorySeries(params)

		So(err, ShouldBeNil)
		Convey("", t, func() {
			So(srs, ShouldNotBeNil)
		})
	})

}

func TestGetCategoryTags(t *testing.T) {

	params := make(map[string]interface{})

	Convey("", t, func() {
		tags, err := xmlFredClient.GetCategoryTags(params)
		So(err, ShouldBeNil)

		Convey("", t, func() {
			So(tags, ShouldNotBeNil)
		})
	})

	Convey("", t, func() {
		tags, err := jsonFredClient.GetCategoryTags(params)

		So(err, ShouldBeNil)
		Convey("", t, func() {
			So(tags, ShouldNotBeNil)
		})
	})

}

func TestGetCategoryRelatedTags(t *testing.T) {

	params := make(map[string]interface{})

	Convey("", t, func() {
		tags, err := xmlFredClient.GetCategoryRelatedTags(params)
		So(err, ShouldBeNil)

		Convey("", t, func() {
			So(tags, ShouldNotBeNil)
		})
	})

	Convey("", t, func() {
		tags, err := jsonFredClient.GetCategoryRelatedTags(params)

		So(err, ShouldBeNil)
		Convey("", t, func() {
			So(tags, ShouldNotBeNil)
		})
	})

}