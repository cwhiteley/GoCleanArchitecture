apk --no-cache add pcre make curl git
curl https://glide.sh/get | sh
cd /go/src/recipes && make build && make run
