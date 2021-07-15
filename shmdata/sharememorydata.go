package shmdata

import (
	"fmt"
	"github.com/kevinu2/shm"
	"unsafe"
)

//todo this will be remove after test ok ,this code form  C
//typedef struct SHMInfo
//{
//unsigned long long max_topic_len;
//unsigned long long max_content_len;
//unsigned long long max_shm_size;
//unsigned long long count;
//key_t key[200];
//} SHMI, *PSHMI;

//todo this will read from  key of 999999  SHM

type SHMInfo struct {
	MaxTopicLen   uint32
	MaxContentLen uint32
	MaxSHMSize    uint32
	Count         uint32
	Key           [200]uint32
}

var MTL uint = 64
var MCL uint = 102400

type TagTLV struct {
	Tag      uint32
	Len      uint32
	TopicLen string
	Value    string
}


//todo  run it use root
func GetShareMemoryInfo(defaultKey int) (SHMInfo, error) {
	var shmi SHMInfo
	shmilen := unsafe.Sizeof(shmi)
	fmt.Printf("sizeof:shmi=%d\r\n",shmilen)
	sh, err := shm.GetSHMInfo(defaultKey, int(shmilen))

	if nil != err {
		return shmi, err
	}

	x := (*[200]uintptr)(unsafe.Pointer(&sh.Data))
	fmt.Printf("data:%#v\n",x)
	fmt.Printf("data:%#v\n",sh)
	shmi = *(*SHMInfo )(unsafe.Pointer(&sh.Data))
	fmt.Printf("shmi:%#v",shmi)
	return shmi, err
}

