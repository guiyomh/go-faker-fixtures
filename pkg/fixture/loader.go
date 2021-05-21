package fixture

import (
	"io/fs"
	"io/ioutil"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type Loader struct {
	logger logrus.FieldLogger
}

func NewLoader(l logrus.FieldLogger) *Loader {
	return &Loader{
		logger: l,
	}
}

func (l Loader) Load(fixturePath string) (map[string]interface{}, error) {
	files, err := ioutil.ReadDir(fixturePath)
	if err != nil {
		return nil, err
	}
	data := make([]byte, 0)
	for _, fi := range files {
		c, err := l.readFile(fi, fixturePath)
		if err != nil {
			l.logger.Warnf("Couldn't read file \"%s\"", fi.Name())
			l.logger.Debug(err.Error())
			continue
		}
		data = append(data, c...)
	}
	return l.parse(data)
}

func (l Loader) readFile(file fs.FileInfo, path string) ([]byte, error) {
	fullPath, _ := filepath.Abs(path + "/" + file.Name())
	l.logger.Debugf("Load file : %s", fullPath)
	content, err := ioutil.ReadFile(fullPath)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func (l Loader) parse(c []byte) (map[string]interface{}, error) {
	tbls := map[string]interface{}{}
	err := yaml.Unmarshal(c, tbls)
	return tbls, err
}
