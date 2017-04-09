# radali [![Build Status](https://travis-ci.org/vbauerster/radali.svg?branch=master)](https://travis-ci.org/vbauerster/radali)

I like online shopping, but I don't like others make money on me, while I'm online shopping.

Have you ever thought why advertisers give you short url like [this](https://goo.gl/yZqJ3p) one,
and when you click it expands to a mile length url, bloated with ugly referral query params?
Well, short answer because they make money like this, without your acknowledgment.
But you can stop this, with the help of this cmd tool.

**radali** will inspect where short url points to, and remove any referral query params,
before opening target url in default web browser.

At the moment following online markets are supported:

* [aliexpress](https://ru.aliexpress.com)
* [banggood](http://www.banggood.com)
* [cashback.epn.bz](https://cashback.epn.bz)
* [coolicool](http://www.coolicool.com)
* [gearbest](http://www.gearbest.com)
* [letyshops](https://letyshops.ru)
* [tinydeal](http://www.tinydeal.com)

## Installation
`radali` requires Go 1.7 or later.
```
$ go get -u github.com/vbauerster/radali
```
Or download [binary](https://github.com/vbauerster/radali/releases/latest).

## Usage
```
Usage: radali [OPTIONS] URL

OPTIONS:
  -d    debug: print debug info
  -p    print only: don't open URL in browser
  -v    print version number
```

## License

[BSD 3-Clause](https://opensource.org/licenses/BSD-3-Clause)
