<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [playstore-ratings](#playstore-ratings)
- [how to build](#how-to-build)
  - [for your current platform](#for-your-current-platform)
  - [for another platform](#for-another-platform)
  - [for all platforms](#for-all-platforms)
- [how to run](#how-to-run)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# playstore-ratings
this application pulls the playstore ratings for all applications on a landing page

# how to build

## for your current platform
```
make
```

## for another platform
```
make linux
make windows
make darwin
```

## for all platforms
```
make all
```

# how to run
by executing the binary it will pull the ratings of all the applications for a landing page and populate your copy/paste buffer with the results.  You can then 'paste' those results into Excel.
double click on the binary for default behavior, which pulls playstore ratings for verizon apps

or run it from the command line and pass the landing page url that you'd like to pull ratings for
