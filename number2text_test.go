package number2text

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInt2Chinese(t *testing.T) {

	words, err := Int2Chinese(123)
	require.NoError(t, err)
	assert.Equal(t, "一百二十三", words)

	words, err = Int2Chinese(1023)
	require.NoError(t, err)
	assert.Equal(t, "一千零二十三", words)

	words, err = Int2Chinese(10203)
	require.NoError(t, err)
	assert.Equal(t, "一万零二百零三", words)

	words, err = Int2Chinese(-123)
	require.NoError(t, err)
	assert.Equal(t, "负一百二十三", words)

	words, err = Int2Chinese(0)
	require.NoError(t, err)
	assert.Equal(t, "零", words)

}

func TestUInt2Chinese(t *testing.T) {

	words, err := uInt2Chinese(123)
	require.NoError(t, err)
	assert.Equal(t, "一百二十三", words)

	words, err = uInt2Chinese(1023)
	require.NoError(t, err)
	assert.Equal(t, "一千零二十三", words)

	words, err = uInt2Chinese(10203)
	require.NoError(t, err)
	assert.Equal(t, "一万零二百零三", words)

	words, err = uInt2Chinese(0)
	require.NoError(t, err)
	assert.Equal(t, "零", words)

	words, err = uInt2Chinese(1234567890987654321)
	require.EqualError(t, err, "internal_error.number_too_big")
	assert.Equal(t, "", words)

}

func TestFloat2Chinese(t *testing.T) {

	words, err := Float2Chinese(123.01, 2)
	require.NoError(t, err)
	assert.Equal(t, "一百二十三点零一", words)

	words, err = Float2Chinese(123.12, 1)
	require.NoError(t, err)
	assert.Equal(t, "一百二十三点一", words)

	words, err = Float2Chinese(123.01, 3)
	require.NoError(t, err)
	assert.Equal(t, "一百二十三点零一", words)

	words, err = Float2Chinese(-123.01, 2)
	require.NoError(t, err)
	assert.Equal(t, "负一百二十三点零一", words)

	words, err = Float2Chinese(123.00, 2)
	require.NoError(t, err)
	assert.Equal(t, "一百二十三", words)

	words, err = Float2Chinese(-0.8123, 2)
	require.NoError(t, err)
	assert.Equal(t, "负零点八一", words)

	words, err = Float2Chinese(-0.9999999999999998, 2)
	require.NoError(t, err)
	assert.Equal(t, "负一", words)

	words, err = Float2Chinese(1.9999999999999998, 2)
	require.NoError(t, err)
	assert.Equal(t, "二", words)

	words, err = Float2Chinese(1.0099999999999998, 3)
	require.NoError(t, err)
	assert.Equal(t, "一点零一", words)

}
