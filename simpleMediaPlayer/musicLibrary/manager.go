package musicLibrary

import "errors"

type MusicEntry struct {
	Id     string //唯一ID
	Name   string //音乐名称
	Artist string //作者
	Source string //音乐位置
	Type   string //音乐类型 mp3 wav
}

type MusicManager struct {
	musics []MusicEntry
}

//NewMusicManager 初始化音乐管理库
func NewMusicManager() *MusicManager {
	return &MusicManager{make([]MusicEntry, 0)}
}

//Len 获取音乐库中音乐的数量
func (m *MusicManager) Len() int {
	return len(m.musics)
}

//Get 根据 索引 查找指定的音乐信息
func (m *MusicManager) Get(index int) (music *MusicEntry, err error) {
	if index < 0 || index >= len(m.musics) {
		return nil, errors.New("index out of range")
	}
	return &m.musics[index], nil
}

//Find 根据 名称 搜索指定的音乐信息
func (m *MusicManager) Find(name string) *MusicEntry {
	if len(m.musics) < 0 {
		return nil
	}
	for _, v := range m.musics {
		if v.Name == name {
			return &v
		}
	}
	return nil
}

//Add 向管理器中添加音乐
func (m *MusicManager) Add(music *MusicEntry) {
	m.musics = append(m.musics, *music)
}

//Remove 删除指定索引的元素
func (m *MusicManager) Remove(index int) *MusicEntry {
	if index < 0 || index >= len(m.musics) {
		return nil
	}

	removeMusic := &m.musics[index]
	//从数组切片中删除元素
	switch {
	case index > 0 && index < len(m.musics)-1: //删除中间元素
		m.musics = append(m.musics[:index-1], m.musics[index+1:]...)
	case index == 0: //删除唯一一个元素
		m.musics = m.musics[index+1:]
	default:
		m.musics = m.musics[:index-1]
	}
	return removeMusic
}
