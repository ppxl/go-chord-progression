go-chord-progression
---

go-chord-progression is a project that aims to create and play harmonic chords.

This project was further developed during my learning day 2025 and the main point is to build schemes of harmonics. While it would be optimal to build chords, I didn't get this to work with the underlying golang library. At least there are melodies which is fine, too.

## How to run

```bash
go run .
```

Or download the binary release and go:
```shell
chmod +x go-chord-progression
./go-chord-progression
```

Then point your browser to http://localhost:9001/index.html

You can view/listen to [example output](go-chord-progression.mp4). Note the output is generated procedurally and never the same.

## Licenses

This work is licensed under the Creative Commons Attribution-NonCommercial-ShareAlike 4.0 International
Public License.

As an artist, I want to encourage others to make art and music but not avoid commercial grifters to take their share.

The [Reset CSS styles](https://meyerweb.com/eric/tools/css/reset/) origins from Eric Meyer, meyerweb, which was released
into the public domain.

## Prior Art

Considered detecting, managing and working with USB audio devices by means of Golang is a bit a toughie for an art project of one day, the idea of playing music through the browser is absolutely lovely. I got heavily inspired by this Dynamsoft page about [playing music with HTML5](https://www.dynamsoft.com/codepool/implement-simple-music-player-in-go.html).

I did not copy code from Dynamsoft, though.

...
