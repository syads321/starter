package schema

// Schema for graphql
var Schema = `
schema {
	query: Query
	mutation: Mutation
}

type Page {
	id: ID!
	title: String!
	description: String!
	url: String!
}

type Activity {
	id: ID!
	title: String!
	date: String!
	url: String!
	caption: String
	shortTitle: String
	images: [ActivityImage!]!
}

type ActivityImage {
	alt: String!
	url: String!
}

input ActivityImageInput {
	alt: String!
	url: String!
}

type Mutation {
	createPage(
		title: String!
		description: String!
	): Page!
	updatePage(
		id: ID!
		title: String!
		description: String!
	): Page!
	deletePage(
		id: ID!
	): Page!
	createActivity(
		title: String!
		date: String!
		caption: String!
		shortTitle: String
		images: [ActivityImageInput!]!
	):Activity!
	updateActivity(
		id: ID!
		title: String!
	):Activity!
	deleteActivity(
		id: ID!
	): Activity!
}

type Query {
	pages: [Page!]!
}

`
