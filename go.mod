module github.com/my-org-org/self-consumer

go 1.22

require github.com/my-org-org/self-producer v1.0.3
replace github.com/my-org-org/self-producer v1.0.3 => github.com/my-org-for-test/self-producer v1.0.3


