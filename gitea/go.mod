module code.gitea.io/sdk/gitea

go 1.12

require (
	code.gitea.io/gitea v1.10.0-dev.0.20190711052757-a0820e09fbf7
	github.com/gorilla/sessions v1.1.3 // indirect
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51 // indirect
	github.com/shurcooL/sanitized_anchor_name v0.0.0-20170918181015-86672fcb3f95 // indirect
	k8s.io/api v0.0.0-20190531132109-d3f5f50bdd94
	pharmer.dev/cloud v0.3.0
	pharmer.dev/pharmer v0.5.0
	sigs.k8s.io/cluster-api v0.1.7
)

replace (
	github.com/renstrom/fuzzysearch => github.com/lithammer/fuzzysearch v1.0.1-0.20160331204855-2d205ac6ec17
	k8s.io/api => k8s.io/api v0.0.0-20190313235455-40a48860b5ab
)
