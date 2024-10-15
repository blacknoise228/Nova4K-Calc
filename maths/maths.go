package maths

import (
	"fmt"
	"strconv"
)

func MCTRL4kResolution(width, height, hz string) (bool, error) {
	w, err := strconv.Atoi(width)
	if err != nil {
		return false, err
	}
	h, err := strconv.Atoi(height)
	if err != nil {
		return false, err
	}
	if w == 0 || h == 0 {
		return false, fmt.Errorf("Invalid resolution")
	}
	if w > MCTRL4KMaxWidthAndHeight || h > MCTRL4KMaxWidthAndHeight {
		return false, fmt.Errorf("MCTRL4K resolution too high")
	}
	resolution := w * h
	switch hz {
	case "60Hz":
		if resolution > Max60MCTRL4K {
			return false, fmt.Errorf("MCTRL4K resolution oversized")
		}
		return true, nil

	case "50Hz":
		if resolution > Max50MCTRL4K {
			return false, fmt.Errorf("MCTRL4K resolution oversized")
		}
		return true, nil
	}
	return false, fmt.Errorf("Error: Not implemented")
}
func VX1000Resolution(width, height, hz string) (bool, error) {
	w, err := strconv.Atoi(width)
	if err != nil {
		return false, err
	}
	h, err := strconv.Atoi(height)
	if err != nil {
		return false, err
	}
	if w == 0 || h == 0 {
		return false, fmt.Errorf("Invalid resolution")
	}
	if w > VX1000MaxWidthAndHeight || h > VX1000MaxWidthAndHeight {
		return false, fmt.Errorf("VX1000 resolution too high")
	}
	resolution := w * h
	switch hz {
	case "60Hz":
		if resolution > Max60VX1000 {
			return false, fmt.Errorf("VX1000 resolution oversized")
		}
		return true, nil
	case "50Hz":
		if resolution > Max50VX1000 {
			return false, fmt.Errorf("VX1000 resolution oversized")
		}
		return true, nil
	case "25Hz":
		if resolution > Max25VX1000 {
			return false, fmt.Errorf("VX1000 resolution oversized")
		}
		return true, nil
	}
	return false, fmt.Errorf("Error: Not implemented")
}
func MCTRL660ProResolution(width, height, hz string) (bool, error) {
	w, err := strconv.Atoi(width)
	if err != nil {
		return false, err
	}
	h, err := strconv.Atoi(height)
	if err != nil {
		return false, err
	}
	if w == 0 || h == 0 {
		return false, fmt.Errorf("Invalid resolution")
	}
	if w > MCTRL660ProMaxWidthAndHeight || h > MCTRL660ProMaxWidthAndHeight {
		return false, fmt.Errorf("MCTRL660Pro resolution too high")
	}
	resolution := w * h
	switch hz {
	case "60Hz":
		if resolution > Max60MCTRL660Pro {
			return false, fmt.Errorf("MCTRL660Pro resolution oversized")
		}
		return true, nil
	case "50Hz":
		if resolution > Max50MCTRL660Pro {
			return false, fmt.Errorf("MCTRL660Pro resolution oversized")
		}
		return true, nil
	}
	return false, fmt.Errorf("Error: Not implemented")
}
