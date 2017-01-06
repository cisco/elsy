/*
 *  Copyright 2016 Cisco Systems, Inc.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package main

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

var configFileOptions map[string]interface{}

func LoadConfigFile(file string) error {
	// ensure we always start with a fresh map
	configFileOptions = nil
	if _, err := os.Stat(file); err != nil {
		return nil
	}
	if s, err := ioutil.ReadFile(file); err != nil {
		return err
	} else if err := yaml.Unmarshal(s, &configFileOptions); err != nil {
		return err
	}
	return nil
}

func getConfigFileValue(key string) interface{} {
	if v, ok := configFileOptions[key]; ok {
		return v
	} else {
		return nil
	}
}

func IsKeyInConfig(key string) bool {
	_, ok := configFileOptions[key]
	return ok
}

func getConfigFileValueWithDefault(key string, defaultValue interface{}) interface{} {
	if v := getConfigFileValue(key); v != nil {
		return v
	} else {
		return defaultValue
	}
}

// GetConfigFileSlice will attempt to convert the value found at the given key
// into a slice, if it fails, or if the key is not found it will return an empty slice.
func GetConfigFileSlice(key string) []string {
	if v, ok := configFileOptions[key].([]interface{}); ok {
		retVal := []string{}
		for _, e := range v {
			if strV, ok := e.(string); ok {
				retVal = append(retVal, strV)
			}
		}
		return retVal
	} else {
		return []string{}
	}
}

func GetConfigFileString(key string) string {
	v, _ := getConfigFileValue(key).(string)
	return v
}

func GetConfigFileStringWithDefault(key string, defaultValue interface{}) string {
	v, _ := getConfigFileValueWithDefault(key, defaultValue).(string)
	return v
}
