# go-opensubtitles #
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/99a6484956d0492db384e518e684e82a)](https://www.codacy.com/manual/TheForgotten69/go-opensubtitles?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=TheForgotten69/go-opensubtitles&amp;utm_campaign=Badge_Grade)

go-opensubtitles is a Go library for accessing the [OpenSubtitles API](https://www.opensubtitles.com/docs/api/html/index.htm).

## Install

To get a specific version from the list of [versions](https://github.com/TheForgotten69/go-opensubtitles/releases):

```sh
go get github.com/TheForgotten69/go-opensubtitles@vX.Y.Z
```

Or for the latest version:

```sh
go get github.com/TheForgotten69/go-opensubtitles
```

### Authentication ###

The go-github library does not directly handle authentication. Instead, when
creating a new client, pass an `http.Client` that can handle authentication for
you. The easiest and recommended way to do this is using the [oauth2][]
library, but you can always use any other library that provides an
`http.Client`. If you have an OAuth2 access token (for example, a [personal
API token][]), you can use it with the oauth2 library using:


## TODO
 

## Ressources
- [Open Subtitles API](https://www.opensubtitles.com/docs/api/html/index.htm)
- https://forum.opensubtitles.org/viewtopic.php?f=8&t=17146