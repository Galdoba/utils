package xdg

import (
	"fmt"
	"os"
	"path/filepath"
)

// Типы каталогов согласно XDG
type DirType int

const (
	Config DirType = iota
	Data
	Cache
	State
	Runtime
)

// Структура для управления путями
type Pathfinder struct {
	appName string
	user    string
	profile string
}

// Инициализация менеджера путей
func New(options ...PathfinderOption) *Pathfinder {
	p := Pathfinder{}
	for _, modify := range options {
		modify(&p)
	}
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	p.user = filepath.Base(home)
	return &p
}

type PathfinderOption func(*Pathfinder)

func AppName(appName string) PathfinderOption {
	return func(p *Pathfinder) {
		p.appName = appName
	}
}

func Profile(profile string) PathfinderOption {
	return func(p *Pathfinder) {
		p.profile = profile
	}
}

// Получение пути по типу
func (p *Pathfinder) FindPath(dirType DirType, layers ...string) (string, error) {
	path := ""
	err := fmt.Errorf("path not called")
	switch dirType {
	case Config:
		path, err = p.configDir()
	case Data:
		path, err = p.dataDir()
	case Cache:
		path, err = p.cacheDir()
	case State:
		path, err = p.stateDir()
	case Runtime:
		path, err = p.runtimeDir()
	default:
		return "", os.ErrInvalid
	}
	if err != nil {
		return "", err
	}
	switch p.profile {
	case "":
		//return "", fmt.Errorf("app name not provided to pathfinder")
	default:
		path = filepath.Join(path, p.profile)
	}
	for _, l := range layers {
		path = filepath.Join(path, l)
	}

	return path, nil
}

// Создание каталога по типу
// func (p *Pathfinder) EnsureDir(dirType DirType, perm os.FileMode) (string, error) {
// 	path, err := p.FindPath(dirType)
// 	if err != nil {
// 		return "", err
// 	}

// 	if err := os.MkdirAll(path, perm); err != nil {
// 		return "", err
// 	}
// 	return path, nil
// }

// func (p *Pathfinder) ConfigFile(filename string, layers ...string) (string, error) {
// 	dir, err := p.EnsureDir(Config, 0700)
// 	fullPath := []string{dir}
// 	fullPath = append(fullPath, layers...)
// 	fullPath = append(fullPath, filename)
// 	return filepath.Join(fullPath...), err
// }

// func (p *Pathfinder) AssetFile(filename string, layers ...string) (string, error) {
// 	dir, err := p.EnsureDir(Data, 0700)
// 	fullPath := []string{dir}
// 	fullPath = append(fullPath, layers...)
// 	fullPath = append(fullPath, filename)
// 	return filepath.Join(fullPath...), err
// }

// func (p *Pathfinder) LogFile(filename string, layers ...string) (string, error) {
// 	dir, err := p.EnsureDir(State, 0700)
// 	fullPath := []string{dir}
// 	fullPath = append(fullPath, layers...)
// 	fullPath = append(fullPath, filename)
// 	return filepath.Join(fullPath...), err
// }

// func (p *Pathfinder) CacheFile(filename string, layers ...string) (string, error) {
// 	dir, err := p.EnsureDir(Cache, 0700)
// 	fullPath := []string{dir}
// 	fullPath = append(fullPath, layers...)
// 	fullPath = append(fullPath, filename)
// 	return filepath.Join(fullPath...), err
// }

// Реализация базовых путей
func (p *Pathfinder) configDir() (string, error) {
	if path := os.Getenv("XDG_CONFIG_HOME"); path != "" {
		return filepath.Join(path, p.appName), nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".config", p.appName), nil
}

func (p *Pathfinder) dataDir() (string, error) {
	if path := os.Getenv("XDG_DATA_HOME"); path != "" {
		return filepath.Join(path, p.appName), nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".local", "share", p.appName), nil
}

func (p *Pathfinder) cacheDir() (string, error) {
	if path := os.Getenv("XDG_CACHE_HOME"); path != "" {
		return filepath.Join(path, p.appName), nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".cache", p.appName), nil
}

func (p *Pathfinder) stateDir() (string, error) {
	if path := os.Getenv("XDG_STATE_HOME"); path != "" {
		return filepath.Join(path, p.appName), nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".local", "state", p.appName), nil
}

func (p *Pathfinder) runtimeDir() (string, error) {
	if path := os.Getenv("XDG_RUNTIME_DIR"); path != "" {
		return filepath.Join(path, p.appName), nil
	}
	return "", os.ErrNotExist
}

// Системные пути (для root-приложений)
func SystemDataDir(appName string) string {
	return filepath.Join("/usr/share", appName)
}

func SystemConfigDir(appName string) string {
	return filepath.Join("/etc", appName)
}

func SystemLogDir(appName string) string {
	return filepath.Join("/var/log", appName)
}
