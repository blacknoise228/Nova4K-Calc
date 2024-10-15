# Resolution Calculator

A simple GUI application for calculating resolutions for various NovaStar sending cards.

## Overview

This application provides a user-friendly interface for calculating resolutions for MCTRL4K, VX1000, and MCTRL660Pro sending cards. It allows users to input the width, height, and refresh rate of their desired resolution and checks whether it is within the supported limits of the selected sending card.

## Features

* Supports MCTRL4K, VX1000, and MCTRL660Pro sending cards
* Calculates resolutions based on user input (width, height, and refresh rate)
* Checks whether the calculated resolution is within the supported limits of the selected sending card
* Displays error messages for invalid or unsupported resolutions

## Requirements

* Go programming language (version 1.14 or later)
* Fyne GUI library (version 2.0 or later)

## Installation

1. Clone the repository: `git clone https://github.com/your-username/resolution-calculator.git`
2. Install the required dependencies: `go get -u fyne.io/fyne/v2`
3. Build the application: `go build main.go`
4. Run the application: `./main`

## Usage

1. Launch the application
2. Select the desired sending card from the main menu
3. Input the width, height, and refresh rate of your desired resolution
4. Click "Submit" to calculate the resolution
5. Check the result: if the resolution is supported, a success message will be displayed; otherwise, an error message will be shown

## Contributing

Contributions are welcome! If you'd like to report a bug or suggest a feature, please open an issue on the GitHub repository. If you'd like to contribute code, please fork the repository and submit a pull request.
