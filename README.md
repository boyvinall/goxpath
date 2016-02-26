# goxpath

[![Build Status](https://travis-ci.org/boyvinall/goxpath.svg?branch=master)](https://travis-ci.org/boyvinall/goxpath)  [![Issue Count](https://codeclimate.com/github/boyvinall/goxpath/badges/issue_count.svg)](https://codeclimate.com/github/boyvinall/goxpath)

Stupid thing to make xpath queries easier from bash. Try a quick google for `bash xpath` and you might be disappointed. #wahwahwah

NB - due to dependence on gokogiri, you need to install libxml2-devel to build this. (Sorry, I'm still a bit of a golang n00b. :blush:)

It'll read from stdin by default

## Install
    yum install libxml2-devel     # see above, soz
    go get github.com/boyvinall/goxpath

## Examples

    ./goxpath -xmlns foo:http://foo -xmlns bar:http://bar -xpath //foo:stuff/bleh -file file.xml

    ./gopath -h    # NB: see ENV names e.g. $XP_XMLNS

    curl -s http://example.com/doc.xml | goxpath -xpath //example

    export XP_XMLNS=foo:http://foo,bar:http://bar # define in the ENV, separate with comma
    goxpath -xpath //foo:stuff file.xml
