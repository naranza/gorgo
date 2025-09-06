// Naranza Gorgo - https://naranza.org
// SPDX-License-Identifier: MPL-2.0
// Copyright (c) Andrea Davanzo and contributors

package gorgo

import (
    "os"
    "strconv"
)

// GetEnv returns the value of the environment variable key,
// or fallback if the key is not set.
func GetEnv(key, fallback string) string {
    if val := os.Getenv(key); val != "" {
        return val
    }
    return fallback
}

// GetEnvInt returns the value of the environment variable key
// parsed as an int, or fallback if not set or invalid.
func GetEnvInt(key string, fallback int) int {
  if val := os.Getenv(key); val != "" {
      if i, err := strconv.Atoi(val); err == nil {
          return i
      }
  }
  return fallback
}
