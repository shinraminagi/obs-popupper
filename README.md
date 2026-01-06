# OBS PopUpper

OBS PopUpper is a simple tool designed to display images or play audio on an OBS overlay in real-time. By dragging and dropping files onto a web interface, the content immediately appears or plays on a connected browser source.

## Features

- **Drag & Drop Interface**: Easily send content by dragging files into the controller page.
- **Image Support**: Displays images with automatic fade-in/out and scaling.
- **Audio Support**: Plays audio files immediately upon drop.
- **Real-time Updates**: Uses WebSockets to instantly push content to the overlay.
- **OBS Integration**: Designed to be used as a Browser Source in OBS Studio.
- **Single Binary**: A standalone executable that is easy to run without complex dependencies.

## Installation

### Prerequisites

- Go 1.18 or higher (for building from source)

### Build from Source

#### 1. Clone the repository:

```bash
git clone https://github.com/shinraminagi/obs-popupper.git
cd obs-popupper
```

#### 2. Build the application:

```bash
go build -o obs-popupper
```

## Usage

### 1. **Start the Server**

Run the executable. By default, it listens on port `25252`.

```bash
./obs-popupper
# Or specify a custom port
./obs-popupper 8080
```

### 2. **Setup OBS Browser Source**

- Add a new "Browser" source in OBS Studio.
- Set the URL to: `http://localhost:25252/popup`
- Set the Width and Height to match your canvas (e.g., 1920x1080).
- Enable "Control audio via OBS" if you want to manage audio levels within OBS.

### 3. **Send Content**

- Open a web browser and navigate to: `http://localhost:25252/`
- Drag and drop an image (`.jpg`, `.png`, etc.) or audio (`.mp3`, `.wav`, etc.) file into the drop zone.
- The content will appear or play on the OBS Browser Source automatically.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
