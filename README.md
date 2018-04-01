# untrack-url [![Build Status](https://travis-ci.org/vbauerster/untrack-url.svg?branch=master)](https://travis-ci.org/vbauerster/untrack-url)

**Why?**

If you follow http://ali.ski/gkMqy and commit a purchase, somebody will earn some money on you.

If you untrack the link with help of `untrack-url`, like:

```
$ untrack-url http://ali.ski/gkMqy
```

all tracking url params will be wiped out and **nobody** will earn money on you.

## Installation
`untrack-url` requires Go 1.7 or later.
```
$ go get -u github.com/vbauerster/untrack-url
```

## Usage
```
Usage: untrack-url [OPTIONS] URL

OPTIONS:
  -p    print only: don't open URL in browser
  -v    print version number

Known trackers:

        ad.admitad.com
        epnclick.ru
        lenkmio.com
        s.click.aliexpress.com
        shopeasy.by
```

## License

[BSD 3-Clause](https://opensource.org/licenses/BSD-3-Clause)
