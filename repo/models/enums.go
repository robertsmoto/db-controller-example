package models

/*
transaction types -- time series??
asn
transfer
refunc
order
*/

type SchemaVersion int

const (
	V2001_001 SchemaVersion = iota
	// add more schema versions here
)

type MemberType int

const (
	article MemberType = iota + 100
	_
	_
	author
	_
	_
	component
	_
	_
	customer
	_
	_
	detail
	_
	_
	document
	_
	_
	image
	_
	_
	location
	_
	_
	order
	_
	_
	page
	_
	_
	part
	_
	_
	person
	_
	_
	product
)

type CollectionType int

const (
	allMembers CollectionType = iota + 200 // needed for account maintenance
	_
	_
	articles
	_
	_
	articleCategories
	articlesByArticleCategory
	_
	_
	articleKeywords
	articlesByArticleKeyword
	_
	_
	articleTags
	articlesByArticleTag
	_
	_
	authors
	articlesByAuthor
	docsByAuthor
	pagesByAuthor
	_
	_
	components
	_
	_
	details
	_
	_
	docs
	_
	_
	docCategories
	docsByDocCategory
	_
	_
	docKeywords
	docByDocKeyword
	_
	_
	docTags
	docsByDocTag
	_
	_
	images
	_
	_
	pages
	_
	_
	pageCategories
	pagesByPageCategory
	_
	_
	pageKeywords
	pagesByPageKeyword
	_
	_
	pageTags
	pagesByPageTag
	_
	_
	parts
	_
	_
	products
	_
	_
	productBrands
	productsByProductBrand
	_
	_
	productCategories
	productsByProductCategory
	_
	_
	productKeywords
	productsByProductKeyword
	_
	_
	productTags
	productsByProductTag
	_
	_
)
