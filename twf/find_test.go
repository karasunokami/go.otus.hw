package twf

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFind(t *testing.T) {
	text := "Олень - северное животное. В летнее время оленям в тайге жарко, а в горах даже в июле холодно. Олень как бы создан для северных просторов жёсткого ветра длинных морозных ночей. Олень легко бежит вперёд по тайге подминает под себя кусты переплывает быстрые реки. Олень не тонет потому, что каждая его шерстинка это длинная трубочка которую внутри наполняет воздух. Нос у оленя покрыт серебристой шёрсткой. Если бы шерсти на носу не было, олень бы его отморозил."
	result, err := Find(text, 10)
	assert.LessOrEqual(t, len(result), 10)
	assert.Nil(t, err)

	result, err = Find("", 10)
	assert.Len(t, result, 0)
	assert.EqualError(t, err, "text parameter cannot be empty")
}
