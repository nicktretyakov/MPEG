# MPEG

A cross-language demonstration of real-time video capture and processing using **Go** (with GoCV/OpenCV) and **Python** (with OpenCV). This project showcases:

- Capturing and displaying live webcam feed in Go
- Handling multiple camera streams concurrently in Python
- Real-time frame display with OpenCV bindings

---

## Table of Contents

1. [Overview](#overview)
2. [Prerequisites](#prerequisites)
3. [Setup Instructions](#setup-instructions)
4. [Project Structure](#project-structure)
5. [Usage](#usage)
   - [Go Webcam Capture](#go-webcam-capture)
   - [Python Multi-Camera Capture](#python-multi-camera-capture)
6. [Code Walkthrough](#code-walkthrough)
7. [Customization](#customization)
8. [Troubleshooting](#troubleshooting)
9. [Contributing](#contributing)
10. [License](#license)

---

## Overview

`MPEG` contains two sample programs:

- **`main.go`**: A Go application using [GoCV](https://gocv.io/) to open and display a single webcam stream.
- **`multiple.py`**: A Python script leveraging [OpenCV](https://opencv.org/) to capture and display frames from two cameras in parallel.

These examples illustrate basic workflows for video input, real-time display, and clean resource management in desktop environments.

## Prerequisites

### Go Application

- **Go** 1.22 or newer
- **OpenCV** 4.x installed on your system
- GoCV v0.35.0 bindings
  ```bash
  go install gocv.io/x/gocv@v0.35.0
  ```

### Python Script

- **Python** 3.7 or newer
- **OpenCV-Python** package
  ```bash
  pip install opencv-python
  ```

## Setup Instructions

1. **Clone the repository**
   ```bash
   git clone https://github.com/nicktretyakov/MPEG.git
   cd MPEG
   ```

2. **Install GoCV dependencies** (Linux example using Ubuntu):
   ```bash
   sudo apt-get update
   sudo apt-get install -y libopencv-dev \
     libgtk-3-dev libjpeg-dev libpng-dev libtiff-dev \
     libavcodec-dev libavformat-dev libavutil-dev libswscale-dev
   ```

3. **Install Python dependencies**:
   ```bash
   pip install opencv-python
   ```

## Project Structure

```
MPEG/
├── main.go         # Go program for single-camera capture
├── multiple.py     # Python script for dual-camera capture
├── go.mod          # Go module definition (GoCV dependency)
└── go.sum          # Checksums for Go modules
```

## Usage

### Go Webcam Capture

Build and run the Go application:

```bash
# Compile
go build -o mpeg-go main.go

# Run
./mpeg-go
```

- A window named **`Webcam`** will open and display the default camera feed.
- Press any key in the window to exit.

### Python Multi-Camera Capture

Execute the Python script:

```bash
python3 multiple.py
```

- Two windows, **`Camera 0`** and **`Camera 1`**, will display the respective webcam feeds.
- Press **`q`** in any window to quit and close all windows.

## Code Walkthrough

### `main.go`

- **Imports** `gocv.io/x/gocv` for OpenCV bindings.
- Opens default video device (`0`) with `gocv.OpenVideoCapture`.
- Creates a GUI window (`gocv.NewWindow`).
- In a loop, reads frames into a `gocv.Mat`, displays with `window.IMShow`, and listens for key events via `window.WaitKey(1)`.
- Cleans up resources with `defer` calls to `Close()`.

### `multiple.py`

- Imports `cv2` from OpenCV-Python.
- Initializes two `cv2.VideoCapture` objects for device indices `0` and `1`.
- In a loop, reads frames from both cameras, shows each in separate windows, and listens for the **`q`** key to exit.
- Releases both capture objects and destroys all windows on exit.

## Customization

- **Device Index**: Change `0`/`1` to your camera IDs.
- **Frame Processing**: Insert image processing (e.g., filtering, edge detection) before display.
- **Resolution**: Use `webcam.Set(gocv.VideoCaptureFrameWidth, width)` and corresponding height calls in Go, or `cap.set(cv2.CAP_PROP_FRAME_WIDTH, width)` in Python.

## Troubleshooting

- **Cannot open camera**: Ensure no other application uses the camera and correct device index.
- **Black window**: Verify OpenCV installation and camera permissions.
- **GoCV errors**: Check that GoCV version matches your OpenCV installation.

## Contributing

Contributions are welcome! Feel free to:

- Report issues or request enhancements.
- Add examples for more advanced video processing.
- Improve OS-specific installation instructions.

1. Fork the repo.
2. Create a feature branch: `git checkout -b feature/xyz`.
3. Commit your changes and push to your fork.
4. Open a pull request.

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.

