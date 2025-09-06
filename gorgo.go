// Naranza Gorgo - https://naranza.org
// SPDX-License-Identifier: MPL-2.0
// Copyright (c) Andrea Davanzo and contributors

package gorgo

import (
  "os"
  "strconv"
)

func GetEnv(key, fallback string) string {
  val := os.Getenv(key)
	if val != "" {
    return val
  }
  return fallback
}

func GetEnvInt(key string, fallback int) int {
  val := os.Getenv(key)
	if val != "" {
		i, err := strconv.Atoi(val)
    if err == nil {
      return i
    }
  }
  return fallback
}
