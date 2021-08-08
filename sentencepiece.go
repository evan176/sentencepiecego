package sentencepiecego

// #cgo LDFLAGS: -lsentencepiecego -lstdc++
// #include <stdlib.h>
// typedef void* SentencePieceProcessorGo;
// SentencePieceProcessorGo loadSentencePieceProcessor(char *path);
// int Encode(SentencePieceProcessorGo sp, char *text, int *tokenIDs, int maxTokens);
// void Free(SentencePieceProcessorGo sp);
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

type SentencePieceProcessor struct {
	index          C.SentencePieceProcessorGo
	beginMaxTokens int
}

func Load(path string) (*SentencePieceProcessor, error) {
	var sp SentencePieceProcessor
	cPath := C.CString(path)
	sp.index = C.loadSentencePieceProcessor(cPath)
	if sp.index == nil {
		return &sp, fmt.Errorf("Can not load model: %s\n", path)
	}
	sp.beginMaxTokens = 128
	C.free(unsafe.Pointer(cPath))
	return &sp, nil
}

func (sp *SentencePieceProcessor) Encode(text string) ([]int, error) {
	var size int
	var ids []int
	size, ids = sp.encode(text, sp.beginMaxTokens)
	if size < 0 {
		return ids, errors.New("Encode error!")
	}
	// If size exceed beginMaxTokens, we extend ids size again!
	if size > sp.beginMaxTokens {
		size, ids = sp.encode(text, size)
		if size < 0 {
			return ids, errors.New("Encode error!")
		}
	}
	return ids, nil
}

func (sp *SentencePieceProcessor) encode(text string, maxTokens int) (int, []int) {
	ids := make([]int, 0, maxTokens)
	CIDs := make([]C.int, maxTokens, maxTokens)
	cText := C.CString(text)
	size := int(C.Encode(sp.index, cText, &CIDs[0], C.int(maxTokens)))
	C.free(unsafe.Pointer(cText))
	// Size < 0 means error
	if size < 0 {
		return size, ids
	}
	// Copy value from CIDs to ids if -1 < size < maxTokens
	if size <= maxTokens {
		for i := 0; i < size; i++ {
			ids = append(ids, int(CIDs[i]))
		}
	}
	return size, ids
}

func (sp *SentencePieceProcessor) Free() {
	C.Free(sp.index)
}
