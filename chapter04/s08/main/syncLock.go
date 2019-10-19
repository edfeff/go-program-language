package main

import "sync"

//同步锁
//Go语言包中的sync包提供了两种锁类型：sync.Mutex和sync.RWMutex
//Mutex 互斥锁
//RWMutex 读写锁

//，任何一个Lock()或RLock()均需要保证对应有Unlock()或RUnlock() 调用与之对应

//。锁的典型使用模式如下：
var LOCK sync.Mutex

func foo() {
	LOCK.Lock()
	defer LOCK.Unlock()
	//...
}
