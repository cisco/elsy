package system

import (
  "io"
  "io/ioutil"
  "errors"
  "fmt"
  "os"
  "net/http"
  "crypto/tls"
  "crypto/md5"
  "runtime"

  "github.com/codegangsta/cli"
  "github.com/Sirupsen/logrus"
  "github.com/kardianos/osext"
)

const tmpDirPrefix = "lcupgrade"
const binaryURL = "https://artifactory1.eng.lancope.local/generic-dev-infrastructure/lc/lc-%s-%s-%s"

// CmdUpgrade will upgrade the current lc binary
func CmdUpgrade(c *cli.Context) error {
  version := c.String("version")
  if len(version) == 0 {
    return errors.New("upgrade command requires a version argument, none found")
  }

  platform := runtime.GOOS
  arch := runtime.GOARCH
  url := fmt.Sprintf(binaryURL, platform, arch, version)
  logrus.Debugf("using url: %s", url)

  // find location of lc currently running
  lcPath, err := getLcLocation()
  if err != nil {
    logrus.Errorf("could not find location of current lc")
    return err
  }

  // hash current binary for comparison with new binary
  oldMd5, err := computeMd5(lcPath)
  if err != nil {
    logrus.Debugf("could not compute md5 for old lc binary")
  }

  //download new binary to staging location
  newTmpDir, newLcTmp, err := downloadNew(url)
  if err != nil {
    return err
  }
  defer os.Remove(newTmpDir)

  // rename current binary in preparation for replacing
  tmpDir, oldLcTmp, err := mvLc(lcPath)
  if err != nil {
    return err
  }
  defer os.Remove(tmpDir)

  //swap in new lc
  if err := swap(newLcTmp, lcPath); err != nil {
    logrus.Debugf("failed swaping new lc from %q to %q, err: %q", newLcTmp, lcPath, err)
    return fmt.Errorf("failed replacing your lc, your old binary is located at %q", oldLcTmp)
  }

  if newMd5, err := computeMd5(lcPath); err != nil {
    logrus.Debugf("could not compute md5 for new lc binary, not comparing them")
  } else {
    if oldMd5 != newMd5 {
      logrus.Infof("lc install finished, new lc binary installed")
    } else {
      logrus.Infof("lc install finished, lc binary was already the latest")
    }
  }
  return nil
}

// move src file into a temp location
// returns:
//  * temporary directory that should be deleted after the upgrade finishes
//  * filePath location of temporary location
func mvLc(src string) (string, string, error) {
  tmpDir, err := ioutil.TempDir("", tmpDirPrefix)
  if err != nil {
    logrus.Debugf("failed creating temp dir ", err)
    return "", "", err
  }
  tmpLocation := fmt.Sprintf("%s/%s", tmpDir, "lc.old")
  logrus.Debugf("moving binary '%s' to '%s'", src, tmpLocation)
  if err := swap(src, tmpLocation); err != nil {
      logrus.Debugf("failed moving binary ", err)
      return "", "", err
  }
  return tmpDir, tmpLocation, nil
}

// swap will rename the src file to the dst file
func swap(src string, dst string) error {
  if err := os.Rename(src, dst); err != nil {
    logrus.Debugf("failed swapping '%s' to '%s'", src, dst, err)
    return err
  }
  return nil
}

func computeMd5(filePath string) (string, error) {
  file, err := os.Open(filePath)
  if err != nil {
    logrus.Debugf("could not open file at '%'", filePath, err)
    return "", err
  }
  defer file.Close()

  md5 := md5.New()
  if _, err := io.Copy(md5, file); err != nil {
    logrus.Debugf("could copy file at '%'", filePath, err)
    return "", err
  }

  var result []byte
  return string(md5.Sum(result)), nil
}

func getLcLocation() (string, error){
  // NOTE: if os.Args[0] is a symlink, this code will update the actual binary, not the link
  lcPath, err := osext.Executable()
  if err != nil {
    logrus.Debugf("lc not found", err)
    return "", err
  }
  return lcPath, nil
}

// Will download the new binary from the given url to a temp location
// returns (dir holding binary, full path of binary)
func downloadNew(url string) (string, string, error) {
  tmpDir, err := ioutil.TempDir("", tmpDirPrefix)
  if err != nil {
    logrus.Debugf("failed creating temp dir ", err)
    return "", "", err
  }
  tmpLocation := fmt.Sprintf("%s/%s", tmpDir, "lc.new")
  logrus.Debugf("downloading new binary to '%s'", tmpLocation)
  if err := installNew(url, tmpLocation); err != nil {
    logrus.Debugf("failed downloading binary, err: %q", err)
    return "", "", err
  }
  return tmpDir, tmpLocation, nil
}

func installNew(url string, target string) error {
  tr := &http.Transport{
      TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
  }
  client := &http.Client{Transport: tr}

  resp, err := client.Get(url)
  if err != nil {
    logrus.Debugf("failed downloading binary", err)
  	return err
  }
  if resp.StatusCode != 200 {
    return fmt.Errorf("failed downloading binary, invalid http response: %d", resp.StatusCode)
  }
  defer resp.Body.Close()

  out, err := os.Create(target)
  if err != nil {
    logrus.Debugf("failed creating new lc file", err)
    return err
  }
  defer out.Close()

  if n, err := io.Copy(out, resp.Body); err != nil {
    logrus.Debugf("failed copying new lc file", err)
    return err
  } else {
    logrus.Debugf("successfully coppied %d bytes", n)
  }

  if err := os.Chmod(target, os.FileMode(0755)); err != nil {
    logrus.Debugf("failed making lc executable", err)
    return err
  }
  return nil
}
