# Muzip - Music ZIP archive indexer

Muzip is a handy little command-line tool to review the contents of a zip archive containing music.

**Supported formats:**
- FLAC
- MP3


## Installation

Muzip can be installed with `go` through the following command

```sh
$ go get -u github.com/tobbbles/muzip
```

Alternatively, you can grab a pre-compiled binary from the [releases](https://github.com/tobbbles/muzip/releases) tab.


## Usage

The `muzip` takes a file through the `-file` command line flag. It'll then parse the supported zip file and print out a table of recognised tracks from the file's metadata.

Example

```
$ muzip -file Calibre\ -\ Even\ If\ \(MP3\).zip
            ARCHIVE           | ARTIST  |        # TITLE
------------------------------+---------+-------------------------
  Calibre - Even If (MP3).zip | Calibre | 1 - All You Got
+                             +         +------------------------+
                              |         | 2 - Even If
+                             +         +------------------------+
                              |         | 3 - Rose
+                             +         +------------------------+
                              |         | 4 - Broken
+                             +         +------------------------+
                              |         | 5 - Thirst Dub
+                             +         +------------------------+
                              |         | 6 - Me Myself & I
+                             +         +------------------------+
                              |         | 7 - Steptoe
+                             +         +------------------------+
                              |         | 8 - Open Your Eyes
+                             +         +------------------------+
                              |         | 9 - Acid Hands
+                             +         +------------------------+
                              |         | 10 - Section Dub
+                             +         +------------------------+
                              |         | 11 - Gone Away
+                             +         +------------------------+
                              |         | 12 - No Reply
+                             +         +------------------------+
                              |         | 13 - Manchester Nights
------------------------------+---------+-------------------------
```