package main

import "testing"

func TestFetchRegularStories(t *testing.T) {
	urls := []string {
		"https://someurl.com/photo.jpg",
		"https://someurl.com/photo_preview.jpg",
		"https://someurl.com/video250.mp4",
		"https://someurl.com/video100.mp4",
		"https://someurl.com/video10.mp4",
	}
	expected := []string {
		"https://someurl.com/photo.jpg",
		"https://someurl.com/video250.mp4",
	}

	result := findOriginalStories(&urls)

	for i, value := range result {
		if value != expected[i] {
			t.Error("Expected ", expected[i], ", but got ", value)
		}
	}
}

func TestFetchPhotoStories(t *testing.T) {
	urls := []string {
		"https://someurl.com/photo.jpg",
	}
	expected := []string {
		"https://someurl.com/photo.jpg",
	}

	result := findOriginalStories(&urls)

	for i, value := range result {
		if value != expected[i] {
			t.Error("Expected ", expected[i], ", but got ", value)
		}
	}
}

func TestFetchVideo3Stories(t *testing.T) {
	urls := []string {
		"https://someurl.com/photo_preview.jpg",
		"https://someurl.com/video250.mp4",
		"https://someurl.com/video100.mp4",
		"https://someurl.com/video10.mp4",
	}
	expected := []string {
		"https://someurl.com/video250.mp4",
	}

	result := findOriginalStories(&urls)

	for i, value := range result {
		if value != expected[i] {
			t.Error("Expected ", expected[i], ", but got ", value)
		}
	}
}

func TestFetchVideo2Stories(t *testing.T) {
	urls := []string {
		"https://someurl.com/photo_preview.jpg",
		"https://someurl.com/video250.mp4",
		"https://someurl.com/video100.mp4",
	}
	expected := []string {
		"https://someurl.com/video250.mp4",
	}

	result := findOriginalStories(&urls)

	for i, value := range result {
		if value != expected[i] {
			t.Error("Expected ", expected[i], ", but got ", value)
		}
	}
}

func TestFetchVideo1Stories(t *testing.T) {
	urls := []string {
		"https://someurl.com/photo_preview.jpg",
		"https://someurl.com/video250.mp4",
	}
	expected := []string {
		"https://someurl.com/video250.mp4",
	}

	result := findOriginalStories(&urls)

	for i, value := range result {
		if value != expected[i] {
			t.Error("Expected ", expected[i], ", but got ", value)
		}
	}
}

func TestFetchPhotosStories(t *testing.T) {
	urls := []string {
		"https://someurl.com/photo1.jpg",
		"https://someurl.com/photo2.jpg",
		"https://someurl.com/photo3.jpg",
	}
	expected := []string {
		"https://someurl.com/photo1.jpg",
		"https://someurl.com/photo3.jpg",
		"https://someurl.com/photo2.jpg",
	}

	result := findOriginalStories(&urls)

	for i, value := range result {
		if value != expected[i] {
			t.Error("Expected ", expected[i], ", but got ", value)
		}
	}
}

func TestFetchStories1(t *testing.T) {
	urls := []string {
		"https://someurl.com/photo1.jpg",
		"https://someurl.com/photo_preview1.jpg",
		"https://someurl.com/video250.mp4",
		"https://someurl.com/video100.mp4",
		"https://someurl.com/video10.mp4",
		"https://someurl.com/photo2.jpg",
		"https://someurl.com/photo3.jpg",
		"https://someurl.com/photo_preview2.jpg",
		"https://someurl.com/video250.mp4",
	}
	expected := [5]string {
		"https://someurl.com/photo1.jpg",
		"https://someurl.com/video250.mp4",
		"https://someurl.com/photo2.jpg",
		"https://someurl.com/photo3.jpg",
		"https://someurl.com/video250.mp4",
	}

	result := findOriginalStories(&urls)

	for i, value := range result {
		if value != expected[i] {
			t.Error("Expected ", expected[i], ", but got ", value)
		}
	}
}

func TestFetchStories2(t *testing.T) {
	urls := []string {
		"https://someurl.com/photo_preview1.jpg",
		"https://someurl.com/video250.mp4",
		"https://someurl.com/video100.mp4",
		"https://someurl.com/video10.mp4",
		"https://someurl.com/photo1.jpg",
		"https://someurl.com/photo2.jpg",
		"https://someurl.com/photo_preview2.jpg",
		"https://someurl.com/video250.mp4",
		"https://someurl.com/photo_preview3.jpg",
		"https://someurl.com/video250.mp4",
		"https://someurl.com/video100.mp4",
		"https://someurl.com/photo3.jpg",
	}
	expected := [6]string {
		"https://someurl.com/video250.mp4",
		"https://someurl.com/photo1.jpg",
		"https://someurl.com/photo2.jpg",
		"https://someurl.com/video250.mp4",
		"https://someurl.com/video250.mp4",
		"https://someurl.com/photo3.jpg",
	}

	result := findOriginalStories(&urls)

	for i, value := range result {
		if value != expected[i] {
			t.Error("Expected ", expected[i], ", but got ", value)
		}
	}
}
