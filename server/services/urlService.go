package services

import "math/rand"

func GenerateShortString(size int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
  shortCode := make([]byte, size)

  for i := range shortCode{
    shortCode[i] = letters[rand.Intn(len(letters))]
  }
  return string(shortCode)
}
