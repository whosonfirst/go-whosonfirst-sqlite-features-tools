# go-whosonfirst-sqlite-features-tools

Internal tool for ensuring that Who's On First SQLite databases can be indexed.

## How it works

A list of repos is fetch from the `whosonfirst-data` organization (using the `go-whosonfirst-github` package). Each repo is cloned in to memory (using the `go-whosonfirst-index-git` package and indexed (using the `go-whosonfirst-sqlite-features-index` package).

Any repo that fails indexing is reported (on `STDOUT`).

This package has not been optimized in any way for performance yet. Depending on the size of any given repo it might take a while to complete. It's the kind of thing you'd imagine running in a nightly or weekly cron job, or equivalent.

## Tools

### inspect

For example:

```
$> go run -mod vendor cmd/inspect/main.go -prefix whosonfirst-data-admin-

2019/12/19 11:59:52 [INFO] Inspect whosonfirst-data-admin-ad at 2019-12-19 11:59:52.930561 -0800 PST m=+15.532730128
2019/12/19 12:00:00 [INFO] Time to inspect whosonfirst-data-admin-ad, 7.650164118s
2019/12/19 12:00:00 [INFO] Inspect whosonfirst-data-admin-ae at 2019-12-19 12:00:00.580876 -0800 PST m=+23.182927640
2019/12/19 12:00:14 [INFO] Time to inspect whosonfirst-data-admin-ae, 13.483440707s
2019/12/19 12:00:14 [INFO] Inspect whosonfirst-data-admin-af at 2019-12-19 12:00:14.064548 -0800 PST m=+36.666391855
...
```

Here's an example of a failed repo:

```
TBW
```

## See also

* https://github.com/whosonfirst/go-whosonfirst-github
* https://github.com/whosonfirst/go-whosonfirst-index-git
* https://github.com/whosonfirst/go-whosonfirst-index-sqlite
* https://github.com/whosonfirst/go-whosonfirst-sqlite
* https://github.com/whosonfirst/go-whosonfirst-sqlite-features
* https://github.com/whosonfirst/go-whosonfirst-sqlite-features-index
