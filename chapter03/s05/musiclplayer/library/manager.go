package library

import "errors"

//音乐类
type MusicEntry struct {
	Id     string
	Name   string
	Artist string
	Source string
	Type   string
}

type MusicManager struct {
	musics []MusicEntry
}

func NewMusicManager() *MusicManager {
	return &MusicManager{make([]MusicEntry, 0)}
}
func (m *MusicManager) Len() int {
	return len(m.musics)
}
func (m *MusicManager) Get(index int) (music *MusicEntry, err error) {
	if index < 0 || index > len(m.musics) {
		return nil, errors.New("index < 0 or index > len")
	}
	return &m.musics[index], nil
}
func (m *MusicManager) Find(name string) (music *MusicEntry) {
	if m.Len() == 0 {
		return nil
	}
	for _, m := range m.musics {
		if name == m.Name {
			return &m
		}
	}
	return nil
}
func (m *MusicManager) Add(music *MusicEntry) {
	m.musics = append(m.musics, *music)
}
func (m *MusicManager) Remove(index int) *MusicEntry {
	if index < 0 || index > len(m.musics) {
		return nil
	}
	removeMusic := &m.musics[index]
	//从切片中删除元素
	if index < len(m.musics)-1 { //中间元素
		m.musics = append(m.musics[:index-1], m.musics[index+1:]...)
	} else if index == 0 { //删除仅有的元素
		m.musics = make([]MusicEntry, 0)
	} else {
		m.musics = m.musics[:index-1]
	}
	return removeMusic
}

func (m *MusicManager) RemoveByName(name string) *MusicEntry {
	if name == "" || len(name) == 0 {
		return nil
	}
	index := -1
	for i, m := range m.musics {
		if m.Name == name {
			index = i
			break
		}
	}
	if index != -1 {
		return m.Remove(index)
	}
	return nil
}
