# 💻 maclabs

Cross-platform command-line text-to-speech powered by ElevenLabs. Works on macOS, Windows, and Linux.

## Install

Requires Go 1.24+.

```bash
go install github.com/krishnashahane/maclabs/cmd/maclabs@latest
```

Or build from source:

```bash
git clone https://github.com/krishnashahane/maclabs.git
cd maclabs
go build ./cmd/maclabs
```

### Cross-platform builds

```bash
# macOS
GOOS=darwin GOARCH=amd64 go build -o maclabs-darwin ./cmd/maclabs

# Windows
GOOS=windows GOARCH=amd64 go build -o maclabs.exe ./cmd/maclabs

# Linux
GOOS=linux GOARCH=amd64 go build -o maclabs-linux ./cmd/maclabs
```

## Configuration

Set your ElevenLabs API key using one of:

- `ELEVENLABS_API_KEY` environment variable
- `MACLABS_API_KEY` environment variable
- `--api-key` flag
- `--api-key-file` flag (or `ELEVENLABS_API_KEY_FILE` / `MACLABS_API_KEY_FILE`)

Optional voice default: `ELEVENLABS_VOICE_ID` or `MACLABS_VOICE_ID`

## Usage

Speak text (streams audio to speakers):

```bash
maclabs "Hello world"
maclabs speak -v Roger "Hello world"
```

Pipe text:

```bash
echo "piped input" | maclabs
```

Save to file:

```bash
maclabs -o output.mp3 "Save this to a file"
```

List voices:

```bash
maclabs voices
maclabs -v ?
maclabs voices --search english --limit 10
```

Speed and rate control:

```bash
maclabs -v Roger -r 200 "Faster speech"
maclabs speak -v Roger --speed 1.2 "Talk a bit faster"
```

Model selection:

```bash
maclabs speak -v Roger --model-id eleven_multilingual_v2 "Stable v2 baseline"
maclabs speak -v Roger --model-id eleven_flash_v2_5 "Fast and cheap"
```

Prompting tips:

```bash
maclabs prompting
```

## Key flags

| Flag | Description |
|---|---|
| `-v, --voice` | Voice name or ID (`?` to list) |
| `--model-id` | Model ID (default: `eleven_v3`) |
| `-r, --rate` | Words per minute (default 175) |
| `--speed` | Speed multiplier (0.5-2.0) |
| `-o, --output` | Write audio to file |
| `-f, --input-file` | Read text from file (`-` for stdin) |
| `--stability` | Voice stability (0-1) |
| `--similarity` | Voice similarity boost (0-1) |
| `--style` | Style exaggeration (0-1) |
| `--stream/--no-stream` | Toggle streaming (default on) |
| `--play/--no-play` | Toggle speaker playback |
| `--metrics` | Print stats to stderr |
| `--normalize` | Text normalization: auto, on, off |
| `--lang` | 2-letter language code |

## Models

| Engine | `--model-id` | Best for |
|---|---|---|
| v3 (default) | `eleven_v3` | Most expressive |
| v2 (stable) | `eleven_multilingual_v2` | Reliable baseline |
| v2.5 Flash | `eleven_flash_v2_5` | Ultra-low latency, 50% cheaper |
| v2.5 Turbo | `eleven_turbo_v2_5` | Low latency, 50% cheaper |

## Development

```bash
go fmt ./...
go test ./...
go build ./cmd/maclabs
```

## License

MIT License - see [LICENSE](LICENSE) for details.
