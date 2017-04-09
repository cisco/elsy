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

package system

import (
	"errors"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/blang/semver"
	"github.com/cheggaaa/pb"
	"github.com/cisco/elsy/helpers"
	"github.com/codegangsta/cli"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	releaseUrl        = "https://api.github.com/repos/cisco/elsy/releases/latest"
	numberOfPlatforms = 2
)

type releaseInfo struct {
	Name    string
	Size    int
	URL     string
	TagName string
}

func CmdUpgrade(c *cli.Context) error {
	currentVersion := strings.Trim(helpers.Version(), "\x00")

	logrus.Infof("Current version: %s", currentVersion)
	logrus.Info("Checking for newer version...")

	if strings.Contains(currentVersion, "snapshot") {
		return errors.New("Upgrade not available on snapshot versions")
	}

	cv, err := semver.Make(currentVersion[1:])
	if err != nil {
		return fmt.Errorf("error parsing current verison: %v", err)
	}

	info, err := getReleaseInfo()
	if err != nil {
		return err
	}

	availableVersion := info.TagName[1:]

	av, err := semver.Make(availableVersion)
	if err != nil {
		return fmt.Errorf("error parsing available verison: %v", err)
	}

	if av.LTE(cv) {
		logrus.Info("No new version available")
		return nil
	}

	logrus.Infof("Upgrading to %s", info.TagName)

	bar := pb.New(info.Size).SetUnits(pb.U_BYTES)
	bar.Start()

	resp, err := http.Get(info.URL)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return fmt.Errorf("error getting installation directory: %v", err)
	}

	outfile := fmt.Sprintf("%s/%s", dir, info.Name)
	writer, err := os.Create(outfile)
	if err != nil {
		return err
	}

	defer writer.Close()

	multiWriter := io.MultiWriter(writer, bar)

	bytesWritten, err := io.Copy(multiWriter, resp.Body)
	if err != nil {
		return err
	}

	if bytesWritten != int64(info.Size) {
		return fmt.Errorf("incorrect byte count %v != %v", info.Size, bytesWritten)
	}

	bar.Finish()

	lcPath := fmt.Sprintf("%s/%s", dir, filepath.Base(os.Args[0]))
	lcWithVersionPath := fmt.Sprintf("%s/%s", dir, info.Name)
	lcLinkName := fmt.Sprintf("%s/lc", dir)

	err = os.Remove(lcPath)
	if err != nil {
		return err
	}

	err = os.Chmod(lcWithVersionPath, 0755)
	if err != nil {
		return err
	}

	err = os.Symlink(lcWithVersionPath, lcLinkName)
	if err != nil {
		return fmt.Errorf("Error symlinking lc: %v", err)
	}

	logrus.Info("Done!")
	return nil
}

func getReleaseInfo() (*releaseInfo, error) {
	resp, err := http.Get(releaseUrl)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	json, err := simplejson.NewFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	tagName, err := json.Get("tag_name").String()
	if err != nil {
		return nil, err
	}

	info := &releaseInfo{TagName: tagName}

	assets := json.Get("assets")

	currentOs := runtime.GOOS

	for i := 0; i < numberOfPlatforms; i++ {
		asset := assets.GetIndex(i)

		info.Name, err = asset.Get("name").String()
		if err != nil {
			return nil, err
		}

		info.URL, err = asset.Get("browser_download_url").String()
		if err != nil {
			return nil, err
		}

		if strings.Contains(info.Name, currentOs) {
			info.Size, err = asset.Get("size").Int()
			return info, nil
		}
	}

	return nil, errors.New("URL not found for OS")
}
