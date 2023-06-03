# ElevenLabs Go Library

> ⚠️ This library is not complete yet and I am mostly building for myself at the moment ⚠️

This is an unofficial Go library for [ElevenLabs](https://elevenlabs.io). This library is intended to have a similar API to the official Python package but it may take some time as I am only building out the bits I need for a personal project I am working on now.

# Installation

```bash
go get github.com/aosasona/eleven@latest
```

# Usage

To begin using the package in your project, you need to setup an instance of the `Eleven` struct that allows you to make requests using the library like this:

```go
import "github.com/aosasona/eleven"

func main() {
    e := eleven.New() // you can pass in your API key here

    // you can also set the API key like this
    // if no API key is set, it will use the `ELEVEN_API_KEY` environment variable i.e os.GetEnv("ELEVEN_API_KEY")
    e.SetAPIKey("...")

}
```

# Methods

These are the methods currently available in the library and how to use them.

## Generate

## Voice

### List Voices

```go
func main() {
    ...
    voices, err := e.ListVoices() // => ([]VoiceResponse, error)
    if err != nil {
        // handle the error here
    }
    ...
}
```
